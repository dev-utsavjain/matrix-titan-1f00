package views

// UserProfileResponse represents user profile response
type UserProfileResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar,omitempty"`
	Bio       string `json:"bio,omitempty"`
	CreatedAt string `json:"createdAt"`
}

// UserStatsResponse represents user statistics response
type UserStatsResponse struct {
	TotalPosts     int64 `json:"totalPosts"`
	PublishedPosts int64 `json:"publishedPosts"`
	DraftPosts     int64 `json:"draftPosts"`
	TotalViews     int64 `json:"totalViews"`
}
