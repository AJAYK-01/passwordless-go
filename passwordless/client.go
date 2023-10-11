package passwordless

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) RegisterToken(request RegisterTokenRequest) (RegisterTokenResponse, error) {
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
		return RegisterTokenResponse{}, errors.New("failed to register, error: " + string(body))
	}

	var response RegisterTokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return RegisterTokenResponse{}, err
	}

	return response, nil
}

func (c *Client) VerifySignin(request VerifySigninRequest) (VerifySigninResponse, error) {
	jsonData, err := json.Marshal(request)
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
		return VerifySigninResponse{}, errors.New("failed to verify, error: " + string(body))
	}

	var response VerifySigninResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return VerifySigninResponse{}, err
	}

	return response, nil
}

func (c *Client) AddAliases(request AliasRequest) error {
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
		return errors.New("failed to add aliases, error: " + string(body))
	}

	return nil
}

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
		return ListCredentialsResponse{}, errors.New("failed to list credentials, error: " + string(body))
	}

	var response ListCredentialsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ListCredentialsResponse{}, err
	}

	return response, nil
}

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
		return errors.New("failed to delete credential, error: " + string(body))
	}

	return nil
}
