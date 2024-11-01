package comment_dto

import (
	"github.com/google/uuid"
	"github.com/poin4003/yourVibes_GoApi/internal/dtos/user_dto"
	"time"
)

type CommentDto struct {
	ID              uuid.UUID                `json:"id"`
	PostId          uuid.UUID                `json:"post_id"`
	UserId          uuid.UUID                `json:"user_id"`
	ParentId        *uuid.UUID               `json:"parent_id"`
	Content         string                   `json:"content"`
	LikeCount       int                      `json:"like_count"`
	RepCommentCount int                      `json:"rep_comment_count"`
	IsLiked         bool                     `json:"is_liked"`
	CreatedAt       time.Time                `json:"created_at"`
	UpdatedAt       time.Time                `json:"updated_at"`
	User            user_dto.UserDtoShortVer `json:"user"`
}

type NewCommentDto struct {
	ID              uuid.UUID  `json:"id"`
	PostId          uuid.UUID  `json:"post_id"`
	UserId          uuid.UUID  `json:"user_id"`
	ParentId        *uuid.UUID `json:"parent_id"`
	Content         string     `json:"content"`
	LikeCount       int        `json:"like_count"`
	RepCommentCount int        `json:"rep_comment_count"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type UpdatedCommentDto struct {
	ID              uuid.UUID                `json:"id"`
	PostId          uuid.UUID                `json:"post_id"`
	UserId          uuid.UUID                `json:"user_id"`
	ParentId        *uuid.UUID               `json:"parent_id"`
	Content         string                   `json:"content"`
	LikeCount       int                      `json:"like_count"`
	RepCommentCount int                      `json:"rep_comment_count"`
	CreatedAt       time.Time                `json:"created_at"`
	UpdatedAt       time.Time                `json:"updated_at"`
	User            user_dto.UserDtoShortVer `json:"user"`
}
