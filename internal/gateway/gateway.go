package gateway

import "github.com/andreyxaxa/posts_comments_service/internal/models"

type Gateways struct {
	Posts
	Comments
}

func NewGateways(posts Posts, comments Comments) *Gateways {
	return &Gateways{
		Posts:    posts,
		Comments: comments,
	}
}

type Posts interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPostById(id int) (models.Post, error)
	GetAllPosts(limit, offset int) ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetCommentsByPost(postId, limit, offset int) ([]*models.Comment, error)
	GetRepliesOfComment(commentId int) ([]*models.Comment, error)
}
