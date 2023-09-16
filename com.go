package kontrakto

type Error struct {
	Message string `json:"message"`
}

type ValidationResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Paginator struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
