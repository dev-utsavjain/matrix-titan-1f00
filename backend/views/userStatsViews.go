package views

type UserStatsResponse struct {
	TotalPosts      int `json:"totalPosts"`
	PublishedPosts  int `json:"publishedPosts"`
	DraftPosts      int `json:"draftPosts"`
	TotalViews      int `json:"totalViews"`
	TotalComments   int `json:"totalComments"`
}
