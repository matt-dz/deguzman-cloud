package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"regexp"
)

// TODO: Make redirect pattern configurable with base url
var (
	redirectPattern = regexp.MustCompile(`^(?:https:\/\/)?[\w-]*\.?deguzman\.cloud(?::\d{1,5})?$`) // Any redirect must be a subdomain of deguzman.cloud
)

func decodeJson(dst interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

func sanitizeRedirect(redirect string) string {
	if os.Getenv("ENV") == "PROD" {
		// Ensure redirect is a relative path or subdomain of deguzman.cloud
		if !((len(redirect) > 0 && redirect[0] == '/') || redirectPattern.MatchString(redirect)) {
			return os.Getenv("BASE_URL")
		}
		return redirect
	}

	if redirect == "" {
		return os.Getenv("SIGNUP_REDIRECT")
	}
	return redirect
}
