package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/janearc/sux/config"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Transport struct {
	url           string
	httpClient    *http.Client
	config        *config.Config
	authenticated bool
}

// NewOpenAITransport creates a new OpenAI transport object.
func NewOpenAITransport(config *config.Config) *Transport {
	return &Transport{
		url:        config.OpenAI.Url,
		httpClient: &http.Client{},
		config:     config,
	}
}

// AuthGood sets the authenticated flag to true.
func (t *Transport) AuthGood() {
	if t.authenticated == true {
		logrus.Warn("AuthGood called but already authenticated")
		return
	}

	t.authenticated = true
}

// AuthBad sets the authenticated flag to false.
func (t *Transport) AuthBad() {
	// github issue #1, this needs to push out to the reauth flow

	if t.authenticated == false {
		logrus.Warn("AuthBad called but already unauthenticated")
		return
	}

	t.authenticated = false
}

// Authenticated returns the authenticated flag.
func (t *Transport) Authenticated() bool {
	return t.authenticated
}

// API Key returns the OpenAPI key
func (t *Transport) APIKey() string {
	return t.config.OpenAI.APIKey
}

// OpenAIRequest sends a simple to OpenAI's API and logs the full response using logrus.
func (t *Transport) OpenAIRequest(prompt string) (string, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "system", "content": "You are an Application Programming Interface providing information to automated services."},
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		logrus.WithError(err).Error("Failed to marshal request body")
		return "", err
	}

	// TODO: would prefer to do this with the http client and recycle, but let's get this working first.

	req, err := http.NewRequest("POST", t.config.OpenAI.Url, bytes.NewBuffer(requestBody))
	if err != nil {
		logrus.WithError(err).Error("Failed to create request")
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+t.APIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		logrus.WithError(err).Error("Failed to send request")
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Error("Failed to read response body")
		return "", err
	}

	logrus.WithField("response_body", string(body)).Info("Full OpenAI Response")

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		logrus.WithError(err).Error("Failed to unmarshal response")
		return "", err
	}

	if choices, ok := response["choices"].([]interface{}); ok && len(choices) > 0 {
		if message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{}); ok {
			return message["content"].(string), nil
		}
	}

	logrus.Error("Unexpected or invalid response from OpenAI backend")
	return "", fmt.Errorf("Unexpected or invalid response from OpenAI backend")
}
