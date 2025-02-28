package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"

	"github.com/andreyxaxa/posts_comments_service/graph"
	"github.com/andreyxaxa/posts_comments_service/internal/models"
)

// Replies is the resolver for the replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *models.Comment) ([]*models.Comment, error) {
	panic(fmt.Errorf("not implemented: Replies - replies"))
}

// CreateComment is the resolver for the CreateComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input models.InputComment) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - CreateComment"))
}

// CommentsSubscription is the resolver for the CommentsSubscription field.
func (r *subscriptionResolver) CommentsSubscription(ctx context.Context, postID int) (<-chan *models.Comment, error) {
	panic(fmt.Errorf("not implemented: CommentsSubscription - CommentsSubscription"))
}

// Comment returns graph.CommentResolver implementation.
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type commentResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
