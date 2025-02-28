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

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, post models.InputPost) (*models.PostGraph, error) {
	panic(fmt.Errorf("not implemented: CreatePost - CreatePost"))
}

// GetAllPosts is the resolver for the GetAllPosts field.
func (r *queryResolver) GetAllPosts(ctx context.Context, page *int, pageSize *int) ([]*models.PostGraph, error) {
	panic(fmt.Errorf("not implemented: GetAllPosts - GetAllPosts"))
}

// GetPostByID is the resolver for the GetPostById field.
func (r *queryResolver) GetPostByID(ctx context.Context, id *int) (*models.Post, error) {
	panic(fmt.Errorf("not implemented: GetPostByID - GetPostById"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
