package db

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
)

func ServiceAccount(secretFile string) *http.Client {
	b, err := os.ReadFile(secretFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	var s struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}
	json.Unmarshal(b, &s)
	config := &jwt.Config{
		Email:      s.Email,
		PrivateKey: []byte(s.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}
