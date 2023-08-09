package dto

import "time"

type WebhookForm struct {
	ID           int       `json:"id"`
	Owner        string    `json:"owner"`
	Name         string    `json:"name"`
	CreatedTime  time.Time `json:"createdTime"`
	Organization string    `json:"organization"`
	ClientIP     string    `json:"clientIp"`
	User         string    `json:"user"`
	Method       string    `json:"method"`
	RequestURI   string    `json:"requestUri"`
	Action       string    `json:"action"`
	IsTriggered  bool      `json:"isTriggered"`
	Object       string    `json:"object"`
}

type WebhookObject struct {
	Owner               string    `json:"owner"`
	Name                string    `json:"name"`
	Organization        string    `json:"organization"`
	CreatedTime         time.Time `json:"createdTime"`
	DisplayName         string    `json:"displayName"`
	Logo                string    `json:"logo"`
	EnablePassword      bool      `json:"enablePassword"`
	EnableSignUp        bool      `json:"enableSignUp"`
	EnableSigninSession bool      `json:"enableSigninSession"`
	EnableCodeSignin    bool      `json:"enableCodeSignin"`
	EnableSamlCompress  bool      `json:"enableSamlCompress"`
}
