package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealworldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (rep *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealworldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (rep *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealworldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (rep *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealworldService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealworldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealworldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealworldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealworldService) AddCommentsToAnArticle(ctx context.Context, req *v1.AddCommentsToanArticleRequest) (rep *v1.SingleCommentReply, err error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *RealworldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealworldService) UnFavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleRequest) (rep *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealworldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (rep *v1.ListOfTagsReply, err error) {
	return &v1.ListOfTagsReply{}, nil
}
