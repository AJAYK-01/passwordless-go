package passwordless

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// The SDK client to interact with the passwordless api.
//
// ApiSecret: The Private API Key from Application dashboard.
//
// BaseUrl: Use "v4.passwordless.dev" or equivalent if changed.
//
// Docs: https://docs.passwordless.dev/guide/api.html#backend-api-reference
type Client struct {
	ApiSecret string
	BaseUrl   string
}

// Creates a registration token for a user used for registration.
func (c *Client) CreateRegisterToken(request RegisterRequest) (RegisterTokenResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return RegisterTokenResponse{}, err
	}

	req, err := http.NewRequest("POST", "https://"+c.BaseUrl+"/register/token", bytes.NewBuffer(jsonData))
	if err != nil {
		return RegisterTokenResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ApiSecret", c.ApiSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return RegisterTokenResponse{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var apiErr APIErrorResponse
		json.Unmarshal(body, &apiErr)
		return RegisterTokenResponse{}, apiErr
	}

	var response RegisterTokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return RegisterTokenResponse{}, err
	}

	return response, nil
}

// Verifies a token and returns the user if the token is valid.
func (c *Client) VerifySignin(token string) (VerifySigninResponse, error) {
	data := map[string]interface{}{
		"token": token,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return VerifySigninResponse{}, err
	}

	req, err := http.NewRequest("POST", "https://"+c.BaseUrl+"/signin/verify", bytes.NewBuffer(jsonData))
	if err != nil {
		return VerifySigninResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ApiSecret", c.ApiSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return VerifySigninResponse{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var apiErr APIErrorResponse
		json.Unmarshal(body, &apiErr)
		return VerifySigninResponse{}, apiErr
	}

	var response VerifySigninResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return VerifySigninResponse{}, err
	}

	return response, nil
}

// Adds aliases for a user. Removes any existing aliases not included in this request.
func (c *Client) SetAliases(request AliasRequest) error {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://"+c.BaseUrl+"/alias", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ApiSecret", c.ApiSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		var apiErr APIErrorResponse
		json.Unmarshal(body, &apiErr)
		return apiErr
	}

	return nil
}

// Lists all credentials for a user.
func (c *Client) ListCredentials(userId string) (ListCredentialsResponse, error) {
	req, err := http.NewRequest("GET", "https://"+c.BaseUrl+"/credentials/list?userId="+userId, nil)
	if err != nil {
		return ListCredentialsResponse{}, err
	}
	req.Header.Set("ApiSecret", c.ApiSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ListCredentialsResponse{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var apiErr APIErrorResponse
		json.Unmarshal(body, &apiErr)
		return ListCredentialsResponse{}, apiErr
	}

	var response ListCredentialsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ListCredentialsResponse{}, err
	}

	return response, nil
}

// Deletes a credential.
func (c *Client) DeleteCredential(request DeleteCredentialRequest) error {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://"+c.BaseUrl+"/credentials/delete", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ApiSecret", c.ApiSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		var apiErr APIErrorResponse
		json.Unmarshal(body, &apiErr)
		return apiErr
	}

	return nil
}
