package kontrakto

type CreateShortLink struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

type UpdateShortLink struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

type DeleteShortLink struct {
	Token string `json:"token"`
}
