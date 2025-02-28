package service

import (
	"github.com/andreyxaxa/posts_comments_service/internal/consts"
	"github.com/andreyxaxa/posts_comments_service/internal/gateway"
	"github.com/andreyxaxa/posts_comments_service/internal/models"
	"github.com/andreyxaxa/posts_comments_service/pkg/logger"
	re "github.com/andreyxaxa/posts_comments_service/pkg/responce_errors"
)

type CommentsService struct {
	repo   gateway.Comments
	logger *logger.Logger
}

func NewCommentsService(repo gateway.Comments, logger *logger.Logger) *CommentsService {
	return &CommentsService{repo: repo, logger: logger}
}

func (c CommentsService) CreateComment(comment models.Comment) (models.Comment, error) {
	if len(comment.Author) == 0 {
		c.logger.Error.Println(consts.EmptyAuthorError)
		return models.Comment{}, re.ResponseError{
			Message: consts.EmptyAuthorError,
			Type:    consts.BadRequestType,
		}
	}

	if len(comment.Content) >= consts.MaxContentLength {
		c.logger.Error.Println(consts.TooLongContentError, len(comment.Content))
		return models.Comment{}, re.ResponseError{
			Message: consts.TooLongContentError,
			Type:    consts.BadRequestType,
		}
	}

	newComment, err := c.repo.CreateComment(comment)
	if err != nil {
		c.logger.Error.Println(consts.CreatingCommentError, err.Error())
		return models.Comment{}, re.ResponseError{
			Message: consts.CreatingCommentError,
			Type:    consts.InternalErrorType,
		}
	}

	return newComment, nil
}

func (c CommentsService) GetCommentsByPost(postId int) ([]*models.Comment, error) {
	//TODO implement me
	panic("implement me")
}
