package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

func ExchangeCodeForTokens(domain, clientID, redirectURI, code, codeVerifier string) (*TokenResponse, error) {
	tokenURL := domain + "/oauth2/token"

	data := url.Values{}

	data.Set("grant_type", "authorization_code")
	data.Set("client_id", clientID)
	data.Set("redirect_uri", redirectURI)
	data.Set("code", code)
	data.Set("code_verifier", codeVerifier)

	req, err := http.NewRequest(http.MethodPost, tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf(
			"token exchange failed: status %d, response: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var tokens TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &tokens, nil
}
