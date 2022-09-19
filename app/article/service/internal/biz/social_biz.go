package biz

import (
	articlePb "conduit/api/article/v1"
	commentsModel "conduit/model/comments_model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SocialRepo interface {
	AddComment(ctx context.Context, body, slug string, authorId int64) (commentsModel.Comments, error)
	GetComments(ctx context.Context, slug string) ([]commentsModel.Comments, error)
	DeleteComment(ctx context.Context, slug string, commentId int64) error
	FavoriteArticle(ctx context.Context, userId, articleId int64) error
	UnFavoriteArticle(ctx context.Context, userId, articleId int64) error

	GetProfile(ctx context.Context, userId int64) (*articlePb.Profile, error)
	GetFavoritesCount(ctx context.Context, aritcleID int64) (int64, error)
	GetFavorited(ctx context.Context, userId, articleId int64) (bool, error)
}

type SocialUsecase struct {
	repo        SocialRepo
	articleRepo ArticleRepo
	log         *log.Helper
}

func NewSocialUsecase(repo SocialRepo, articleRepo ArticleRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{
		repo:        repo,
		articleRepo: articleRepo,
		log:         log.NewHelper(logger),
	}
}

func (s *SocialUsecase) AddComment(ctx context.Context, req *articlePb.AddCommentRequest) (*articlePb.Comment, error) {
	userId := req.GetUserId()
	res, err := s.repo.AddComment(ctx, req.Body, req.Slug, userId)
	if err != nil {
		return nil, err
	}
	// 作者信息
	author, err := s.repo.GetProfile(ctx, userId)
	if err != nil {
		s.log.WithContext(ctx).Errorf("AddComment GetProfile err: %s", err.Error())
	}
	return &articlePb.Comment{
		Id:        res.ID,
		Body:      res.Body,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Author:    author,
	}, nil
}

func (s *SocialUsecase) GetComments(ctx context.Context, req *articlePb.GetCommentsRequest) ([]*articlePb.Comment, error) {
	// 获取文章评论
	res, err := s.repo.GetComments(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}
	var rr = make([]*articlePb.Comment, 0)
	for _, v := range res {
		// 评论者的信息
		author, err := s.repo.GetProfile(ctx, v.AuthorID)
		if err != nil {
			s.log.WithContext(ctx).Errorf("AddComment GetProfile err: %s", err.Error())
		}
		rr = append(rr, &articlePb.Comment{
			Id:        v.ID,
			Body:      v.Body,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Author:    author,
		})
	}

	return rr, nil
}

func (s *SocialUsecase) DeleteComment(ctx context.Context, req *articlePb.DeleteCommentRequest) error {
	return s.repo.DeleteComment(ctx, req.Slug, req.CommentId)
}

func (s *SocialUsecase) FavoriteArticle(ctx context.Context, req *articlePb.FavoriteArticleRequest) (*articlePb.ArticleData, error) {
	userId := req.GetUserId()
	// get article id by slug
	articles, err := s.articleRepo.GetArticle(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}
	err = s.repo.FavoriteArticle(ctx, userId, articles.ID)
	if err != nil {
		return nil, err
	}

	// get article
	// 作者信息
	author, err := s.repo.GetProfile(ctx, articles.AuthorID)
	if err != nil {
		s.log.WithContext(ctx).Errorf("AddComment GetProfile err: %s", err.Error())
	}
	// 获取tag
	tags, err := s.articleRepo.GetTagsFromArticleId(ctx, articles.ID)
	if err != nil {
		s.log.WithContext(ctx).Errorf("GetTagsFromArticleId err: %s", err.Error())
	}

	favoritesCount, err := s.repo.GetFavoritesCount(ctx, articles.ID)
	if err != nil {
		s.log.WithContext(ctx).Errorf("GetFavoritesCount err: %s", err.Error())
	}

	//favorited, err := s.repo.GetFavorited(ctx, userId, articles.ID)
	//if err != nil {
	//	s.log.WithContext(ctx).Errorf("GetFavorited err: %s", err.Error())
	//}

	return &articlePb.ArticleData{
		Slug:           articles.Slug,
		Title:          articles.Title,
		Description:    articles.Description,
		Body:           articles.Body,
		TagList:        tags,
		CreatedAt:      articles.CreatedAt,
		UpdatedAt:      articles.UpdatedAt,
		Author:         author,
		FavoritesCount: favoritesCount,
		Favorited:      true,
	}, nil
}

func (s *SocialUsecase) UnFavoriteArticle(ctx context.Context, req *articlePb.UnFavoriteArticleRequest) (*articlePb.ArticleData, error) {
	articles, err := s.articleRepo.GetArticle(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}
	err = s.repo.UnFavoriteArticle(ctx, req.GetUserId(), articles.ID)
	if err != nil {
		return nil, err
	}
	// 作者信息
	author, err := s.repo.GetProfile(ctx, articles.AuthorID)
	if err != nil {
		s.log.WithContext(ctx).Errorf("AddComment GetProfile err: %s", err.Error())
	}
	// 获取tag
	tags, err := s.articleRepo.GetTagsFromArticleId(ctx, articles.ID)
	if err != nil {
		s.log.WithContext(ctx).Errorf("GetTagsFromArticleId err: %s", err.Error())
	}

	favoritesCount, err := s.repo.GetFavoritesCount(ctx, articles.ID)
	if err != nil {
		s.log.WithContext(ctx).Errorf("GetFavoritesCount err: %s", err.Error())
	}

	//favorited, err := s.repo.GetFavorited(ctx, req.GetUserId(), articles.ID)
	//if err != nil {
	//	s.log.WithContext(ctx).Errorf("GetFavorited err: %s", err.Error())
	//}
	return &articlePb.ArticleData{
		Slug:           articles.Slug,
		Title:          articles.Title,
		Description:    articles.Description,
		Body:           articles.Body,
		TagList:        tags,
		CreatedAt:      articles.CreatedAt,
		UpdatedAt:      articles.UpdatedAt,
		Author:         author,
		FavoritesCount: favoritesCount,
		Favorited:      false,
	}, nil
}
