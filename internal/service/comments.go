package service

import "github.com/andreyxaxa/posts_comments_service/internal/gateway"

type CommentsService struct {
	repo gateway.Comments
}

func NewCommentsService(repo gateway.Comments) *CommentsService {
	return &CommentsService{repo: repo}
}
