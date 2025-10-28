package oktaauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
	//"pam-okta-helper/util"
	"pam-okta-helper/oktaerror"
)
var logger *log.Logger = log.Default()

func SetLogger(l *log.Logger) {
	if l != nil {
		logger = l
	}
}

const (
  //oktaDomain = "https://trial-1508978.okta.com"
  oktaDomain = "https://dev-86659676.okta.com"
)


// AuthnResponse captures the key fields from /authn
//move this to its own file
type AuthnResponse struct {
	Status       string `json:"status"`
	SessionToken string `json:"sessionToken,omitempty"`
	Embedded     struct {
		Factors []struct {
			ID         string `json:"id"`
			FactorType string `json:"factorType"`
			Provider   string `json:"provider"`
			Profile    struct {
				CredentialID string `json:"credentialId"`
				PhoneNumber  string `json:"phoneNumber,omitempty"`
			} `json:"profile"`
		} `json:"factors"`
	} `json:"_embedded"`
	// add other fields if you need them
}
//for error struct
var oe oktaerror.OktaError
// Login calls /api/v1/authn and returns the parsed response
func Login(username, password string) (*AuthnResponse, error) {
	logger.Printf("POST /api/v1/authn for user=%q", username)
	payload := map[string]string{
		"username": username,
		"password": password,
	}

	// put these as subatomic functs to import
	body, err := json.Marshal(payload)
    if err != nil {
      return nil, fmt.Errorf("failed to marshal login payload: %w", err)
    }

	client := &http.Client{Timeout: 10 * time.Second}
	// make subatomic funct
	req, err := http.NewRequest("POST", oktaDomain+"/api/v1/authn", bytes.NewReader(body))
	if err != nil {
	  return nil, fmt.Errorf("failed to create authn request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	logger.Printf("Authn HTTP %d → %s", resp.StatusCode, string(bodyBytes))
	if resp.StatusCode != http.StatusOK {

	json.Unmarshal(bodyBytes, &oe)
    if err := json.Unmarshal(bodyBytes, &oe); err == nil && oe.ErrorSummary != "" {
        return nil, fmt.Errorf("authentication failed: %s", oe.ErrorSummary)
    }
    // fallback if it wasn’t the expected shape
    return nil, fmt.Errorf("authentication failed: HTTP %d: %s",
        resp.StatusCode, string(bodyBytes))
	}
	var authnResp AuthnResponse
	if err := json.Unmarshal(bodyBytes, &authnResp); err != nil {
    	return nil, fmt.Errorf("failed to parse authn response: %w", err)
	}


	if authnResp.Status == "" {
		return nil, errors.New("no status in authn response")
	}
	logger.Printf("Okta status=%q sessionToken=%q factors=%v",
    authnResp.Status, authnResp.SessionToken, authnResp.Embedded.Factors)
	return &authnResp, nil
}

// VerifyFactor calls /api/v1/authn/factors/{id}/verify with a passCode
func VerifyFactor(factorID, passCode string) (*AuthnResponse, error) {
	payload := map[string]string{
		"passCode": passCode,
	}
	body, _ := json.Marshal(payload)

	client := &http.Client{Timeout: 10 * time.Second}
	endpoint := fmt.Sprintf("%s/api/v1/authn/factors/%s/verify", oktaDomain, factorID)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var factorResp AuthnResponse
	if err := json.NewDecoder(resp.Body).Decode(&factorResp); err != nil {
		return nil, err
	}
	return &factorResp, nil
}
