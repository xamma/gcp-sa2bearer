package token

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	authURL      = "https://www.googleapis.com/oauth2/v4/token"
	expiresInSec = 3600
)

func CreateSignedJWT(privateKey, privateKeyID, email, tokenURI string, scopes string) (string, error) {
	issued := time.Now().Unix()
	expires := issued + expiresInSec

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":    email,
		"sub":    email,
		"aud":    tokenURI,
		"iat":    issued,
		"exp":    expires,
		"scope":  scopes,
		"kid":    privateKeyID,
		"alg":    "RS256",
		"typ":    "JWT",
	})

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func ExchangeJWTForAccessToken(jwtToken, tokenURI string) (string, error) {
	params := url.Values{}
	params.Set("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	params.Set("assertion", jwtToken)

	resp, err := http.PostForm(tokenURI, params)
	if err != nil {
		return "", fmt.Errorf("failed to request access token: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to exchange JWT for access token: %s", string(body))
	}

	var respData struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.Unmarshal(body, &respData); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	return respData.AccessToken, nil
}
