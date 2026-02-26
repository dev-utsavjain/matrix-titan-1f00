package views

// PostListResponse represents paginated post list response
type PostListResponse struct {
	Posts      interface{} `json:"posts"`
	Pagination interface{} `json:"pagination"`
}

// PostResponse represents a single post response
type PostResponse struct {
	Post interface{} `json:"post"`
}
