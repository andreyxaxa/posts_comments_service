package service

import (
	"github.com/andreyxaxa/posts_comments_service/internal/gateway"
	"github.com/andreyxaxa/posts_comments_service/internal/models"
	"github.com/andreyxaxa/posts_comments_service/pkg/logger"
)

type Services struct {
	Posts
	Comments
}

func NewServices(gateways *gateway.Gateways, logger *logger.Logger) *Services {
	return &Services{
		Posts:    NewPostsService(gateways.Posts, logger),
		Comments: NewCommentsService(gateways.Comments, logger),
	}
}

type Posts interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPostById(id int) (models.Post, error)
	GetAllPosts(page, pageSize *int) ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetCommentsByPost(postId int) ([]*models.Comment, error)
	GetRepliesOfComment(commentId int) ([]*models.Comment, error)
}
