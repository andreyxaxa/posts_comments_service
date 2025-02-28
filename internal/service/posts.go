package service

import "github.com/andreyxaxa/posts_comments_service/internal/gateway"

type PostsService struct {
	repo gateway.Posts
}

func NewPostsService(repo gateway.Posts) *PostsService {
	return &PostsService{repo: repo}
}
