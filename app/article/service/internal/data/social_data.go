package data

import (
	articlePb "conduit/api/article/v1"
	userPb "conduit/api/user/v1"
	"conduit/app/article/service/internal/biz"
	articleFavModel "conduit/model/article_favorites_model"
	commentsModel "conduit/model/comments_model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type socialRepo struct {
	data *Data
	log  *log.Helper
}

// NewSocialRepo .
func NewSocialRepo(data *Data, logger log.Logger) biz.SocialRepo {
	return &socialRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *socialRepo) AddComment(ctx context.Context, body, slug string, authorId int64) (commentsModel.Comments, error) {
	var comment = commentsModel.Comments{
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		ArticleSlug: slug,
		Body:        body,
		AuthorID:    authorId,
	}
	err := r.data.db.WithContext(ctx).Create(&comment).Error
	if err != nil {
		return commentsModel.Comments{}, err
	}
	return comment, nil
}

func (r *socialRepo) GetComments(ctx context.Context, slug string) ([]commentsModel.Comments, error) {
	var comments = make([]commentsModel.Comments, 0)
	err := r.data.db.WithContext(ctx).Where("article_slug = ?", slug).Find(&comments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return comments, nil
}

func (r *socialRepo) DeleteComment(ctx context.Context, slug string, commentId int64) error {
	return r.data.db.WithContext(ctx).Where("article_slug = ? and id = ?", slug, commentId).Delete(&commentsModel.Comments{}).Error
}

func (r *socialRepo) FavoriteArticle(ctx context.Context, userId, articleId int64) error {
	var fav = &articleFavModel.ArticleFavorites{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		UserID:    userId,
		ArticleID: articleId,
	}
	err := r.data.db.WithContext(ctx).Create(&fav).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *socialRepo) UnFavoriteArticle(ctx context.Context, userId, articleId int64) error {
	return r.data.db.WithContext(ctx).Where("user_id = ? and article_id = ?", userId, articleId).Delete(&articleFavModel.ArticleFavorites{}).Error
}

func (r *socialRepo) GetProfile(ctx context.Context, userId int64) (*articlePb.Profile, error) {
	author, err := r.data.userService.GetProfileById(ctx, &userPb.GetProfileByIdRequest{Id: int32(userId)})
	if err != nil {
		return nil, err
	}
	var authorInfo = new(articlePb.Profile)
	if author != nil {
		authorInfo.Bio = author.Profile.Bio
		authorInfo.UserName = author.Profile.UserName
		authorInfo.Image = author.Profile.Image
		authorInfo.Following = author.Profile.Following
	}
	return authorInfo, nil
}

func (r *socialRepo) GetFavoritesCount(ctx context.Context, aritcleID int64) (int64, error) {
	var count int64
	err := r.data.db.WithContext(ctx).Where("article_id = ?", aritcleID).Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}

func (r *socialRepo) GetFavorited(ctx context.Context, userId, articleId int64) (bool, error) {
	var af articleFavModel.ArticleFavorites
	err := r.data.db.WithContext(ctx).Where("article_id = ? and user_id = ?", articleId, userId).First(&af).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
