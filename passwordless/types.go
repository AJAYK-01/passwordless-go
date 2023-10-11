package passwordless

type Client struct {
	ApiSecret string
	BaseUrl   string
}

type RegisterTokenRequest struct {
	UserId      string `json:"userId"`
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
}

type APIErrorResponse struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	ErrorCode string `json:"errorCode"`
}

type RegisterTokenResponse struct {
	Token string `json:"token"`
}

type VerifySigninRequest struct {
	Token string `json:"token"`
}

type VerifySigninResponse struct {
	Success   bool   `json:"success"`
	UserId    string `json:"userId"`
	Timestamp string `json:"timestamp"`
	Rpid      string `json:"rpid"`
	Origin    string `json:"origin"`
	Device    string `json:"device"`
	Country   string `json:"country"`
	Nickname  string `json:"nickname"`
	ExpiresAt string `json:"expiresAt"`
	TokenId   string `json:"tokenId"`
	Type      string `json:"type"`
}

type AliasRequest struct {
	UserId  string   `json:"userId"`
	Aliases []string `json:"aliases"`
	Hashing bool     `json:"hashing"`
}

type CredentialDescriptor struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Credential struct {
	Descriptor       CredentialDescriptor `json:"descriptor"`
	PublicKey        string               `json:"publicKey"`
	UserHandle       string               `json:"userHandle"`
	SignatureCounter int                  `json:"signatureCounter"`
	CreatedAt        string               `json:"createdAt"`
	AaGuid           string               `json:"aaGuid"`
	LastUsedAt       string               `json:"lastUsedAt"`
	Rpid             string               `json:"rpid"`
	Origin           string               `json:"origin"`
	Country          string               `json:"country"`
	Device           string               `json:"device"`
	Nickname         string               `json:"nickname"`
	UserId           string               `json:"userId"`
}

type ListCredentialsResponse struct {
	Values []Credential `json:"values"`
}

type DeleteCredentialRequest struct {
	CredentialId string `json:"credentialId"`
}
