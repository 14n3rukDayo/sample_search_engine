package schema
type AddDocumentRequest struct {
	Description string `json:"description" validate:"required"`
}
type AddDocumentResponse struct {
	Message string `json:"message" validate:"required"`
}
