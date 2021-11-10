package client

import (
	"fmt"
	"github.com/imroc/req"
)

type AppDClient struct {
	BaseUrl string
	Secret  string
}

func (c *AppDClient) createUrl(applicationId int) string {
	return fmt.Sprintf("%s/controller/alerting/rest/v1/applications/%d", c.BaseUrl, applicationId)
}

func (c *AppDClient) createAuthHeader() req.Header {
	return req.Header{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", c.Secret),
		"Accept":        "application/json",
	}
}

func createAccessTokenUrl(baseUrl string) string {
	return fmt.Sprintf("%s/controller/api/oauth/access_token", baseUrl)
}

type AppdTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func CreateAccessToken(controllerBaseUrl string, clientName string, clientSecret string) (*string, error) {
	authHeader := req.Header{
		"Content-Type": "application/vnd.appd.cntrl+protobuf;v=1",
	}

	param := req.Param{
		"grant_type":    "client_credentials",
		"client_id":     clientName,
		"client_secret": clientSecret,
	}

	resp, err := req.Post(createAccessTokenUrl(controllerBaseUrl), authHeader, param)
	if err != nil {
		return nil, err
	}
	appdTokenResponse := AppdTokenResponse{}
	err = resp.ToJSON(&appdTokenResponse)
	if err != nil {
		return nil, err
	}

	return &appdTokenResponse.AccessToken, nil
}
