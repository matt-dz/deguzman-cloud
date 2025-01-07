package api

import (
	"encoding/json"
	"errors"
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
	"os"
	"path/filepath"
	"qr-generator/internal/logger"
)

var log = logger.GetLogger()
var SRV_DIR = os.Getenv("SRV_DIR")

const urlHashLength = 8
const fileExtension = ".png"

func HandleGenerateCode(w http.ResponseWriter, r *http.Request) {
	log.Info("Handling GenerateCode request")

	if SRV_DIR == "" {
		log.Error("TMP_DIR environment variable not set")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	/* Parse Request */
	log.Info("Parsing request")
	var payload GenerateCodePayload
	err := DecodeJSONBody(w, r, &payload)
	if err != nil {
		var mr *MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		log.Error("Error decoding request body", "error", err.Error())
		return
	}

	if payload.Url == "" {
		log.Error("Missing required fields", "field", "url")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	/* Hash URL */
	log.Info("Hashing URL")
	hash, err := HashURL(payload.Url, urlHashLength)
	if err != nil {
		log.Error("Error hashing URL", "error", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	filename := hash + fileExtension

	/* Generate QR Code */
	log.Info("Generating QR Code")
	path := filepath.Join(SRV_DIR, filename)
	if _, err := os.Stat(path); err == nil {
		log.Info("QR code already exists. Skipping generation.", "path", path)
	} else {
		err = qrcode.WriteFile(payload.Url, qrcode.Medium, 256, path)
		if err != nil {
			log.Error("Error generating QR code", "error", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	log.Info("QR Code generated", "path", path)

	/* Send Response */
	log.Info("Encoding response")
	err = json.NewEncoder(w).Encode(GenerateCodeResponse{Id: hash})
	if err != nil {
		log.Error("Error encoding response", "error", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Info("Response sent")
}

func HandleGetCode(w http.ResponseWriter, r *http.Request) {
	log.Info("Handling GetCode request")

	if SRV_DIR == "" {
		log.Error("TMP_DIR environment variable not set")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	/* Parse Request */
	log.Info("Parsing request")
	id := r.PathValue("id")
	if id == "" {
		log.Error("Missing required fields", "field", "id")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	/* Check if file exists */
	log.Info("Checking if file exists")
	filename := id + fileExtension
	path := filepath.Join(SRV_DIR, filename)
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Error("QR code does not exist", "path", path)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			log.Error("Error checking if file exists", "error", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	/* Send file */
	log.Info("Sending file")
	http.ServeFile(w, r, path)
}
