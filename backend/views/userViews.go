package views

type UserStatsResponse struct {
	TotalPosts int64 `json:"totalPosts"`
	TotalViews int64 `json:"totalViews"`
}
