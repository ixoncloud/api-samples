package ixon

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AuthService service

func (s *AuthService) Login(username string, otpToken string, password string) error {
	// Create http basic format
	authString := fmt.Sprintf("%s:%s:%s", username, otpToken, password)

	// Base64 encode authentication string
	encodedAuthString := base64.StdEncoding.EncodeToString([]byte(authString))

	// Create authorization header
	authorization := fmt.Sprintf("Basic %s", encodedAuthString)

	headers := make(map[string]string)
	headers[AuthorizationHeader] = authorization

	// Create payload body
	body := make(map[string]int)
	body["expiresIn"] = 604800

	// Create query params
	queryParams := make(map[string]string)
	queryParams["fields"] = "secretId"

	// Make login request
	res, err := s.client.MakeRequest(
		http.MethodPost,
		s.client.apiEndpoints["AccessTokenList"],
		headers,
		body,
		queryParams,
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		if res.StatusCode == http.StatusUnauthorized {
			var errorRes ApiErrorResponse

			err = json.NewDecoder(res.Body).Decode(&errorRes)

			verifyLoginResponse(&errorRes)
		}

		return errors.New("could not login to IXplatform")

	}

	var loggedinResponse AccessTokenResponse

	err = json.NewDecoder(res.Body).Decode(&loggedinResponse)

	if err != nil {
		return err
	}

	s.client.accessToken = loggedinResponse.Data.SecretId
	s.client.authenticated = true

	return nil
}

func verifyLoginResponse(loginResp *ApiErrorResponse) {
	message := loginResp.Data[0].Message
	if message != "" {
		switch message {
		case "One-Time Password required":
			{
				log.Error("Your account has 2FA enabled. Please enter your second factor code with the '--otp' flag")
			}
		case "One-Time Password failed":
			{
				log.Error("The entered second factor code was incorrect")
			}
		default:
			{
				log.Error("Incorrect Credentials")
			}
		}
	} else {
		log.Error("Incorrect Credentials")
	}
}
