package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"errors"

	"github.com/andreyxaxa/posts_comments_service/graph"
	"github.com/andreyxaxa/posts_comments_service/internal/consts"
	"github.com/andreyxaxa/posts_comments_service/internal/models"
	re "github.com/andreyxaxa/posts_comments_service/pkg/responce_errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Replies is the resolver for the replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *models.Comment) ([]*models.Comment, error) {
	comments, err := r.CommentsService.GetRepliesOfComment(obj.ID)
	if err != nil {
		var rErr re.ResponseError
		if errors.As(err, &rErr) {
			return nil, &gqlerror.Error{
				Extensions: rErr.Extensions(),
			}
		}
	}

	return comments, nil
}

// CreateComment is the resolver for the CreateComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input models.InputComment) (*models.Comment, error) {
	newComment, err := r.CommentsService.CreateComment(input.FromInput())
	if err != nil {
		var rErr re.ResponseError
		if errors.As(err, &rErr) {
			return nil, &gqlerror.Error{
				Extensions: rErr.Extensions(),
			}
		}
	}

	if err := r.CommentsObservers.NotifyObservers(newComment.Post, newComment); err != nil {
		if err.Error() != consts.ThereIsNoObserversError {
			return nil, &gqlerror.Error{
				Extensions: map[string]interface{}{
					"message": err,
					"type":    consts.InternalErrorType,
				},
			}
		}
	}

	return &newComment, nil
}

// CommentsSubscription is the resolver for the CommentsSubscription field.
func (r *subscriptionResolver) CommentsSubscription(ctx context.Context, postID int) (<-chan *models.Comment, error) {
	id, ch, err := r.CommentsObservers.CreateObserver(postID)

	if err != nil {
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"message": err,
				"type":    consts.InternalErrorType,
			},
		}
	}

	go func() {
		<-ctx.Done()
		err := r.CommentsObservers.DeleteObserver(id, postID)
		if err != nil {
			// TODO: error log
		}
	}()

	return ch, nil
}

// Comment returns graph.CommentResolver implementation.
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type commentResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
