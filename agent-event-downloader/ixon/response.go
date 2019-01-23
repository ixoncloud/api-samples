package ixon

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type ApiResponse struct {
	Links  []Link `json:"links"`
	Type   string `json:"type"`
	Count  int    `json:"count"`
	Status string `json:"status"`
}

type ApiErrorResponse struct {
	ApiResponse
	Data []struct {
		Message string `json:"message"`
	}
}

type AccessTokenResponse struct {
	ApiResponse
	Data struct {
		SecretId string `json:"secretId"`
		Message  string `json:"message"`
	} `json:"data"`
}

type Resource struct {
	Type     string `json:"type"`
	PublicId string `json:"publicId"`
	Name     string `json:"name"`
}
