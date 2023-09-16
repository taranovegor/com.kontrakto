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
	WithToken string `json:"with_token"`
	Location  string `json:"location"`
}

type UpdateShortLinkResult struct {
	Success            bool
	Message            string
	LocationValidation ValidateShortLinkLocation
	ShortLink          ShortLink
}

type DeleteShortLink struct {
	WithToken string `json:"with_token"`
}

type GetShortLink struct {
	WithToken string `json:"with_token"`
}

type PaginateShortLinks struct {
	Paginator `json:"paginator"`
}

type ShortLinksPagination struct {
	Items []ShortLink `json:"items"`
}
