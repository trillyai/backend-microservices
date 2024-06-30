package shared

import "time"

type Feed struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Username         string `json:"userName"`
	UserProfileImage string `json:"userProfileImage"`
	Description      string `json:"description"`

	LikeCount    uint       `json:"likeCount"`
	CommentCount uint       `json:"commentCount"`
	CreatedDate  *time.Time `json:"createdDate"`
}