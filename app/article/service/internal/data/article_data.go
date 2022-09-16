package data

import (
	"conduit/app/article/service/internal/biz"
	articleTagsModel "conduit/model/article_tags_model"
	"conduit/model/articles_model"
	tagsModel "conduit/model/tags_model"
	"conduit/pkg/middleware/auth"
	"conduit/pkg/mysql"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) GetArticle(ctx context.Context, slug string) (*articlesModel.Articles, error) {
	var d = articlesModel.Articles{}
	result := r.data.db.WithContext(ctx).Where("slug = ?", slug).First(&d)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &d, nil
}

func (r *articleRepo) CreateArticle(ctx context.Context, article *articlesModel.Articles) error {
	return r.data.db.WithContext(ctx).Create(article).Error
}

func (r *articleRepo) UpdateArticle(ctx context.Context, oldSlug string, article *articlesModel.Articles) error {
	return r.data.db.WithContext(ctx).Where("slug = ?", oldSlug).Updates(article).Error
}

func (r *articleRepo) DeleteArticle(ctx context.Context, slug string) error {
	err := r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		m := &articlesModel.Articles{}
		err := tx.Where("slug = ?", slug).First(m).Error
		if err != nil {
			return err
		}
		// delete article table
		err = tx.Where("slug = ?", slug).Delete(m).Error
		if err != nil {
			return err
		}

		// delete tags 关联表
		err = tx.Where("article_id = ?", m.ID).Delete(&articleTagsModel.ArticleTags{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *articleRepo) BatchGetArticles(ctx context.Context, opts ...mysql.QueryOption) ([]*articlesModel.Articles, error) {
	var res []*articlesModel.Articles
	db := r.data.db.WithContext(ctx)
	query := mysql.QueryOptions{}
	for _, o := range opts {
		o(&query)
	}
	err := createQuery(ctx, db, query).Debug().Find(&res).Error
	return res, err
}

func createQuery(ctx context.Context, db *gorm.DB, query mysql.QueryOptions) *gorm.DB {
	if len(query.AuthorName) > 0 {
		db = db.Joins("join users on users.id = articles.author_id and users.username = ?", query.AuthorName)
	}

	if len(query.Tag) > 0 {
		db = db.Joins("join tags on tags.name = ?", query.Tag).Joins("join article_tags on tags.id = article_tags.tag_id and article_tags.article_id = articles.id")
	}

	if len(query.Favorited) > 0 {
		userId := auth.GetUserIdFromContext(ctx)
		db = db.Joins("join article_favorites on article_favorites.article_id = articles.id and article_favorites.user_id = ?", userId)
	}

	if query.Offset > 0 {
		db = db.Offset(int(query.Offset))
	}

	if query.Offset > 0 {
		db = db.Limit(int(query.Limit))
	}
	return db
}

func (r *articleRepo) GetTagsFromArticleId(ctx context.Context, articleId int64) ([]string, error) {
	var res []tagsModel.Tags
	err := r.data.db.WithContext(ctx).Joins("join article_tags on article_tags.tag_id = tags.id").Where("article_tags.article_id = ?", articleId).Find(&res).Error
	if err != nil {
		return nil, err
	}
	var tagList []string
	if len(res) > 0 {
		for _, vv := range res {
			tagList = append(tagList, vv.Name)
		}
	}
	return tagList, nil
}

func (r *articleRepo) CreateArticleTags(ctx context.Context, articleId int64, tags []string) error {
	err := r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		tagList := make([]tagsModel.Tags, 0, len(tags))
		// tag是否已存在
		err := tx.Where("name in ?", tags).Find(&tagList).Error
		if err != nil {
			return err
		}

		tagMap := make(map[string]struct{}, len(tags))
		if len(tagList) > 0 {
			for _, v := range tagList {
				tagMap[v.Name] = struct{}{}
			}
		}

		newTagList := make([]tagsModel.Tags, 0, len(tags))
		if len(tags) > 0 {
			for _, v := range tags {
				_, ok := tagMap[v]
				if !ok {
					// 新的tag要单独创建
					newTagList = append(newTagList, tagsModel.Tags{
						CreatedAt: time.Now().Unix(),
						UpdatedAt: time.Now().Unix(),
						Name:      v,
					})
				}
			}
		}

		if len(newTagList) > 0 {
			// 创建新的tag
			err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&newTagList).Error
			if err != nil {
				return nil
			}
			tagList = append(tagList, newTagList...)
		}

		articleTags := make([]articleTagsModel.ArticleTags, 0, len(tagList))
		for _, v := range tagList {
			articleTags = append(articleTags, articleTagsModel.ArticleTags{
				TagID:     v.ID,
				ArticleID: articleId,
			})
		}

		err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&articleTags).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *articleRepo) FeedArticles(ctx context.Context, limit, offset, userId int64) ([]*articlesModel.Articles, error) {
	var res []*articlesModel.Articles
	err := r.data.db.WithContext(ctx).Joins("join followings on articles.author_id = followings.following_id").Where("followings.user_id = ?", userId).Limit(int(limit)).Offset(int(offset)).Order("created_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *articleRepo) GetTags(ctx context.Context) ([]tagsModel.Tags, error) {
	var tagList = make([]tagsModel.Tags, 0)
	err := r.data.db.WithContext(ctx).Find(&tagList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tagList, nil
}
