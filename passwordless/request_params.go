package passwordless

import "time"

// Required parameters are user ID, username and display name, others are optional.
type RegisterRequest struct {
	UserId            string    `json:"userId"`
	Username          string    `json:"username"`
	Displayname       string    `json:"displayname"`
	Attestation       string    `json:"attestation,omitempty"`
	AuthenticatorType string    `json:"authenticatorType,omitempty"`
	Discoverable      bool      `json:"discoverable,omitempty"`
	UserVerification  string    `json:"userVerification,omitempty"`
	ExpiresAt         time.Time `json:"expiresAt,omitempty"`
	Aliases           []string  `json:"aliases,omitempty"`
	AliasHashing      bool      `json:"aliasHashing,omitempty"`
}

// Hashing parameter optional and defaults to true.
type AliasRequest struct {
	UserId  string   `json:"userId"`
	Aliases []string `json:"aliases"`
	Hashing bool     `json:"hashing,omitempty"`
}

type DeleteCredentialRequest struct {
	CredentialId string `json:"credentialId"`
}
