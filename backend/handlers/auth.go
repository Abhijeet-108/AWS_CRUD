package handlers

import (
	"log"
	"net/http"
	"net/url"

	"employee-management-backend/auth"
	"employee-management-backend/config"

	"github.com/gin-gonic/gin"
)

func LoginHandler(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		state, err := auth.GenerateState()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to generate state",
			})
			return
		}

		codeVerifier, err := auth.GenerateCodeVerifier()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to generate code verifier",
			})
			return
		}

		codeChallenge := auth.GenerateCodeChallenge(codeVerifier)

		c.SetCookie(
			"oauth_state",
			state,
			300,
			"/",
			"",
			false,
			true,
		)

		c.SetCookie(
			"pkce_verifier",
			codeVerifier,
			300,
			"/",
			"",
			false,
			true,
		)

		params := url.Values{}

		params.Add("response_type", "code")
		params.Add("client_id", cfg.CognitoClientID)
		params.Add("redirect_uri", cfg.CognitoRedirectURI)
		params.Add("scope", "openid email")
		params.Add("state", state)
		params.Add("code_challenge_method", "S256")
		params.Add("code_challenge", codeChallenge)

		authorizationURL := cfg.CognitoDomain + "/oauth2/authorize?" + params.Encode()
		// fmt.Println("Authorization URL:")

		// fmt.Println("Authorization URL:")
		// fmt.Println(authorizationURL)

		c.Redirect(http.StatusFound, authorizationURL)

	}
}

func CallbackHandler(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		state := c.Query("state")

		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "authorization code is missing",
			})
			return
		}
		if state == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "state is missing",
			})
			return
		}

		expectedState, err := c.Cookie("oauth_state")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to retrieve state from cookie",
			})
			return
		}

		if state != expectedState {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid state",
			})
			return
		}

		codeVerifier, err := c.Cookie("pkce_verifier")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "pkce_verifier cookie is missing",
			})
			return
		}

		tokens, err := auth.ExchangeCodeForTokens(
			cfg.CognitoDomain,
			cfg.CognitoClientID,
			cfg.CognitoRedirectURI,
			code,
			codeVerifier,
		)

		if err != nil {
			log.Printf("TOKEN EXCHANGE ERROR: %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to exchange authorization code for tokens",
			})
			return
		}

		c.SetCookie(
			"ACCESS_TOKEN",
			tokens.AccessToken,
			tokens.ExpiresIn,
			"/",
			"",
			false,
			true,
		)

		c.SetCookie(
			"REFRESH_TOKEN",
			tokens.RefreshToken,
			30*24*60*60,
			"/",
			"",
			false,
			true,
		)

		// remove temporary Oauth cookies
		c.SetCookie(
			"oauth_state",
			"",
			-1,
			"/",
			"",
			false,
			true,
		)

		c.SetCookie(
			"pkce_verifier",
			"",
			-1,
			"/",
			"",
			false,
			true,
		)

		c.Redirect(http.StatusFound, "http://localhost:5173")
	}
}

func LogoutHandler(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie(
			"ACCESS_TOKEN",
			"",
			-1,
			"/",
			"",
			false,
			true,
		)

		c.SetCookie(
			"REFRESH_TOKEN",
			"",
			-1,
			"/",
			"",
			false,
			true,
		)

		logoutURL :=
			cfg.CognitoDomain +
				"/logout" +
				"?client_id=" +
				url.QueryEscape(cfg.CognitoClientID) +
				"&logout_uri=" +
				url.QueryEscape("http://localhost:5173")

		c.Redirect(
			http.StatusFound,
			logoutURL,
		)

	}
}

func LogoutCompleteHandler(c *gin.Context) {
	c.Redirect(
		http.StatusFound,
		"/api/auth/login",
	)
}
