package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

// --------------------------------------------------------------------------------------------------
func (s *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}

// --------------------------------------------------------------------------------------------------
func (s *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (*v1.EmptyReply, error) {
	return &v1.EmptyReply{}, nil
}

// --------------------------------------------------------------------------------------------------
func (s *RealWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (*v1.SingleCommentReply, error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, req *v1.GetCommentsRequest) (*v1.MultipleCommentsReply, error) {
	return &v1.MultipleCommentsReply{}, nil
}

func (s *RealWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (*v1.EmptyReply, error) {
	return &v1.EmptyReply{}, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}

// --------------------------------------------------------------------------------------------------
func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (*v1.TagListReply, error) {
	return &v1.TagListReply{}, nil
}
