package main

import (
	"context"
	"crypto/sha256"
	"dss/internal/api"
	"dss/internal/database"
	"dss/internal/logger"
	"net/http"
	"os"
)

var logging = logger.GetLogger()
var db = database.GetDb()
var ctx = context.Background()

func HashDB() {

	logging.Info("Hashing unhashed files in DB")

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		logging.Error("BASE_URL not set")
		return
	}

	dssToken := os.Getenv("DSS_TOKEN")
	if dssToken == "" {
		logging.Error("DSS_TOKEN not set")
		return
	}

	logging.Info("Querying DB for unhashed files")
	query := `
	SELECT key FROM files
	`
	rows, err := db.Query(ctx, query)
	if err != nil {
		logging.Error("Failed to query db", "error", err.Error())
		return
	}

	logging.Info("Hashing files")
	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			logging.Error("Failed to scan row", "error", err.Error())
			continue
		}

		logging.Info("Downloading file", "key", key)
		req, err := http.NewRequest("GET", baseUrl+"/"+key, nil)
		if err != nil {
			logging.Error("Failed to create request", "error", err.Error())
			continue
		}
		req.Header.Set("Authorization", "Bearer "+dssToken)

		httpClient := &http.Client{}
		resp, err := httpClient.Do(req)
		if err != nil {
			logging.Error("Failed to download file", "error", err.Error())
			continue
		}
		defer resp.Body.Close()

		logging.Info("Hashing file", "key", key)
		hash, err := api.HashContents(resp.Body, sha256.New())
		if err != nil {
			logging.Error("Failed to hash file", "error", err.Error())
			continue
		}

		logging.Info("Updating hash in DB")
		updateQuery := `
		UPDATE files SET hash = $1 WHERE key = $2
		`
		if _, err := db.Exec(ctx, updateQuery, hash, key); err != nil {
			logging.Error("Failed to update hash in db", "error", err.Error())
		}
		logging.Info("Hashed file", "key", key)
	}

}

func main() {
	HashDB()
}
