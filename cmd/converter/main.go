package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/xamma/gcp-sa2bearer/internal/token"
	"github.com/xamma/gcp-sa2bearer/internal/config"
)

var (
	keyFilePath string
	scopes      = "https://www.googleapis.com/auth/cloud-platform"
)

func main() {
	flag.StringVar(&keyFilePath, "keyfile", "", "Path to the service account JSON file")
	flag.Parse()

	keyFile, err := os.ReadFile(keyFilePath)
	if err != nil {
		log.Fatalf("Error reading service account key file: %v", err)
	}

	var serviceAccount config.ServiceAccountKey
	if err := json.Unmarshal(keyFile, &serviceAccount); err != nil {
		log.Fatalf("Error parsing service account JSON: %v", err)
	}

	jwtToken, err := token.CreateSignedJWT(serviceAccount.PrivateKey, serviceAccount.PrivateKeyID, serviceAccount.ClientEmail, serviceAccount.TokenURI, scopes)
	if err != nil {
		log.Fatalf("Error generating JWT token: %v", err)
	}

	accessToken, err := token.ExchangeJWTForAccessToken(jwtToken, serviceAccount.TokenURI)
	if err != nil {
		log.Fatalf("Error exchanging JWT for access token: %v", err)
	}

	fmt.Printf("Access Token: %s\n", accessToken)
}
