package kontrakto

const (
	RegexShortLinkToken = "[\\w\\-]{4,64}"
)

type ValidateShortLinkToken struct {
	Token string `json:"token"`
}

type ValidateShortLinkUrl struct {
	Url string `json:"url"`
}

type ShortLink struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

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
