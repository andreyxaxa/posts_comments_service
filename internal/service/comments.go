package service

import (
	"github.com/andreyxaxa/posts_comments_service/internal/gateway"
	"github.com/andreyxaxa/posts_comments_service/pkg/logger"
)

type CommentsService struct {
	repo   gateway.Comments
	logger *logger.Logger
}

func NewCommentsService(repo gateway.Comments, logger *logger.Logger) *CommentsService {
	return &CommentsService{repo: repo, logger: logger}
}
