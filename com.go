package kontrakto

type ValidationResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Paginator struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
