package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/poin4003/yourVibes_GoApi/internal/dtos/post_dto"
	"github.com/poin4003/yourVibes_GoApi/internal/model"
	"github.com/poin4003/yourVibes_GoApi/internal/query_object"
	"github.com/poin4003/yourVibes_GoApi/pkg/response"
	"mime/multipart"
)

type (
	IPostUser interface {
		CreatePost(ctx context.Context, postModel *model.Post, inMedia []multipart.File) (post *model.Post, resultCode int, httpStatusCode int, err error)
		UpdatePost(ctx context.Context, postId uuid.UUID, updateData map[string]interface{}, deleteMediaIds []uint, inMedia []multipart.File) (post *model.Post, resultCode int, httpStatusCode int, err error)
		DeletePost(ctx context.Context, postId uuid.UUID) (resultCode int, httpStatusCode int, err error)
		GetPost(ctx context.Context, postId uuid.UUID, userId uuid.UUID) (postDto *post_dto.PostDto, resultCode int, httpStatusCode int, err error)
		GetManyPosts(ctx context.Context, query *query_object.PostQueryObject, userId uuid.UUID) (postDtos []*post_dto.PostDto, resultCode int, httpStatusCode int, response *response.PagingResponse, err error)
	}
	IPostLike interface {
		LikePost(ctx context.Context, likeUserPost *model.LikeUserPost, userId uuid.UUID) (postDto *post_dto.PostDto, resultCode int, httpStatusCode int, err error)
		GetUsersOnLikes(ctx context.Context, postId uuid.UUID, query *query_object.PostLikeQueryObject) (users []*model.User, resultCode int, httpStatusCode int, pagingResponse *response.PagingResponse, err error)
	}
	IPostShare interface {
		SharePost(ctx context.Context, postId uuid.UUID, userId uuid.UUID, shareInput *post_dto.SharePostInput) (post *model.Post, resultCode int, httpStatusCode int, err error)
	}
)

var (
	localPostUser     IPostUser
	localLikeUserPost IPostLike
	localPostShare    IPostShare
)

func PostUser() IPostUser {
	if localPostUser == nil {
		panic("repository_implement localPostUser not found for interface IPostUser")
	}

	return localPostUser
}

func LikeUserPost() IPostLike {
	if localLikeUserPost == nil {
		panic("repository_implement localLikeUserPost not found for interface ILikeUserPost")
	}

	return localLikeUserPost
}

func PostShare() IPostShare {
	if localPostShare == nil {
		panic("repository_implement localPostShare not found for interface IPostShare")
	}

	return localPostShare
}

func InitPostUser(i IPostUser) {
	localPostUser = i
}

func InitLikeUserPost(i IPostLike) {
	localLikeUserPost = i
}

func InitPostShare(i IPostShare) {
	localPostShare = i
}
