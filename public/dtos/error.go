package dtos

type ErrorResponseDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  struct {
		ID int `json:"id"`
	} `json:"errors"`
}
