package views

type PostResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	Excerpt       string    `json:"excerpt,omitempty"`
	Content       string    `json:"content"`
	FeaturedImage string    `json:"featuredImage,omitempty"`
	Status        string    `json:"status"`
	Featured      bool      `json:"featured"`
	Views         int       `json:"views"`
	ReadTime      int       `json:"readTime"`
	AuthorID      string    `json:"authorId"`
	CategoryID    string    `json:"categoryId,omitempty"`
	CreatedAt     string    `json:"createdAt"`
	UpdatedAt     string    `json:"updatedAt"`
	Author        UserResponse `json:"author,omitempty"`
	Category      CategoryResponse `json:"category,omitempty"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

type CategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}
