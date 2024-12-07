package schema
type SearcherRequest struct {
	SearchWords string `json:"searchWord" validate:"required"`
}
type DocumentResponse struct {
	DocumentId  int    `json:"documentId" valdiate:"required"`
	Description string `json:"desciption" valdiate:"required"`
}
type SearchResponse struct {
	Documents []DocumentResponse `json:"documents" valdiate:"required"`
}
