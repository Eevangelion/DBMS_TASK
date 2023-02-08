package models

type UserResponseSearch struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Role             string `json:"role"`
	PostsCount       int    `json:"posts_count"`
	SubscribersCount int    `json:"subscribers_count"`
}
