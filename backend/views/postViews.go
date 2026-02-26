package views

type PostListResponse struct {
	Posts []PostResponse `json:"posts"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

type PostResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	Excerpt       string    `json:"excerpt"`
	FeaturedImage string    `json:"featuredImage"`
	ReadTime      int       `json:"readTime"`
	Views         int       `json:"views"`
	Featured      bool      `json:"featured"`
	Status        string    `json:"status"`
	CreatedAt     string    `json:"createdAt"`
	Author        AuthorResponse `json:"author"`
	Category      CategoryResponse `json:"category,omitempty"`
}

type AuthorResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}