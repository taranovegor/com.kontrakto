package kontrakto

type ValidateShortLinkToken struct {
	Token string `json:"token"`
}

type ValidateShortLinkLocation struct {
	Location string `json:"location"`
}

type ShortLink struct {
	Token    string `json:"token"`
	Location string `json:"location"`
	Url      string `json:"url"`
}

type CreateShortLink struct {
	Token    string `json:"token"`
	Location string `json:"location"`
}

type CreateShortLinkResult struct {
	Success            bool
	Message            string
	TokenValidation    ValidationResult
	LocationValidation ValidationResult
	ShortLink          ShortLink
}

type UpdateShortLink struct {
	WithToken string `json:"token"`
	Location  string `json:"location"`
}

type UpdateShortLinkResult struct {
	Success            bool
	Message            string
	LocationValidation ValidateShortLinkLocation
	ShortLink          ShortLink
}

type DeleteShortLink struct {
	WithToken string `json:"token"`
}
