package service

import (
	"database/sql"
	"errors"

	"github.com/andreyxaxa/posts_comments_service/internal/consts"
	"github.com/andreyxaxa/posts_comments_service/internal/gateway"
	"github.com/andreyxaxa/posts_comments_service/internal/models"
	"github.com/andreyxaxa/posts_comments_service/pkg/logger"
	"github.com/andreyxaxa/posts_comments_service/pkg/pagination"
	re "github.com/andreyxaxa/posts_comments_service/pkg/responce_errors"
)

type CommentsService struct {
	repo       gateway.Comments
	logger     *logger.Logger
	PostGetter PostGetter
}

type PostGetter interface {
	GetPostById(id int) (models.Post, error)
}

func NewCommentsService(repo gateway.Comments, logger *logger.Logger, getter PostGetter) *CommentsService {
	return &CommentsService{repo: repo, logger: logger, PostGetter: getter}
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

	post, err := c.PostGetter.GetPostById(comment.Post)
	if err != nil {
		c.logger.Error.Println(consts.GettingPostError, err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return models.Comment{}, re.ResponseError{
				Message: consts.PostNotFountError,
				Type:    consts.NotFoundType,
			}
		}
	}

	if !post.CommentsAllowed {
		c.logger.Error.Println(consts.CommentsNotAllowedError, err.Error())
		return models.Comment{}, re.ResponseError{
			Message: consts.CommentsNotAllowedError,
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

func (c CommentsService) GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error) {

	if postId <= 0 {
		c.logger.Error.Println(consts.WrongIdError, postId)
		return nil, re.ResponseError{
			Message: consts.WrongIdError,
			Type:    consts.BadRequestType,
		}
	}

	offset, limit := pagination.GetOffsetAndLimit(page, pageSize)

	comments, err := c.repo.GetCommentsByPost(postId, limit, offset)
	if err != nil {
		c.logger.Error.Println(consts.GettingCommentError, postId, err.Error())
		return nil, re.ResponseError{
			Message: consts.GettingCommentError,
			Type:    consts.InternalErrorType,
		}
	}

	return comments, nil
}

func (c CommentsService) GetRepliesOfComment(commentId int) ([]*models.Comment, error) {

	if commentId <= 0 {
		c.logger.Error.Println(consts.WrongIdError, commentId)
		return nil, re.ResponseError{
			Message: consts.WrongIdError,
			Type:    consts.BadRequestType,
		}
	}

	comments, err := c.repo.GetRepliesOfComment(commentId)
	if err != nil {
		c.logger.Error.Println(consts.GettingRepliesError, commentId, err.Error())
		return nil, re.ResponseError{
			Message: consts.GettingRepliesError,
			Type:    consts.InternalErrorType,
		}
	}

	return comments, nil

}
