package views

type UserStatsResponse struct {
	TotalPosts     int64 `json:"totalPosts"`
	PublishedPosts int64 `json:"publishedPosts"`
	DraftPosts     int64 `json:"draftPosts"`
	TotalViews     int64 `json:"totalViews"`
}
