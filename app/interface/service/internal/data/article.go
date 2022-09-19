package data

import (
	"conduit/api/article/v1"
	"conduit/api/interface/v1"
	"conduit/app/interface/service/internal/biz"
	"conduit/pkg/format"
	"conduit/pkg/middleware/auth"
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &ArticleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (a *ArticleRepo) GetArticleBySlug(ctx context.Context, slug string) (*interfacePb.SingleArticle, error) {
	reply, err := a.data.ac.GetArticleBySlug(ctx, &articlePb.GetArticleBySlugRequest{Slug: slug})
	if err != nil {
		return nil, err
	}

	return convertSingleArticle(reply), nil
}

func (a *ArticleRepo) ListArticles(ctx context.Context, tag, author, favorited string, limit, offset int64) ([]*interfacePb.SingleArticle, error) {
	reply, err := a.data.ac.ListArticles(ctx, &articlePb.ListArticlesRequest{
		Tag:       tag,
		Author:    author,
		Favorited: favorited,
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		return nil, err
	}

	var res = make([]*interfacePb.SingleArticle, 0, limit)
	for _, v := range reply.Article {
		res = append(res, &interfacePb.SingleArticle{
			Slug:           v.Slug,
			Title:          v.Title,
			Description:    v.Description,
			Body:           v.Body,
			TagList:        v.TagList,
			CreatedAt:      format.ConvertTime(v.CreatedAt),
			UpdatedAt:      format.ConvertTime(v.UpdatedAt),
			Favorited:      v.Favorited,
			FavoritesCount: int32(v.FavoritesCount),
			Author: &interfacePb.Profile{
				Username:  v.Author.UserName,
				Bio:       v.Author.Bio,
				Image:     v.Author.Image,
				Following: v.Author.Following,
			},
		})
	}
	return res, nil
}

func (a *ArticleRepo) FeedArticles(ctx context.Context, limit, offset int64) ([]*interfacePb.SingleArticle, error) {
	userId := auth.GetUserIdFromContext(ctx)
	reply, err := a.data.ac.FeedArticles(ctx, &articlePb.FeedArticlesRequest{
		Limit:  limit,
		Offset: offset,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	var res = make([]*interfacePb.SingleArticle, 0, limit)
	for _, v := range reply.Article {
		res = append(res, &interfacePb.SingleArticle{
			Slug:           v.Slug,
			Title:          v.Title,
			Description:    v.Description,
			Body:           v.Body,
			TagList:        v.TagList,
			CreatedAt:      format.ConvertTime(v.CreatedAt),
			UpdatedAt:      format.ConvertTime(v.UpdatedAt),
			Favorited:      v.Favorited,
			FavoritesCount: int32(v.FavoritesCount),
			Author: &interfacePb.Profile{
				Username:  v.Author.UserName,
				Bio:       v.Author.Bio,
				Image:     v.Author.Image,
				Following: v.Author.Following,
			},
		})
	}
	return res, nil
}

func (a *ArticleRepo) CreateArticle(ctx context.Context, req *interfacePb.CreateArticleRequest) (*interfacePb.SingleArticle, error) {
	userId := auth.GetUserIdFromContext(ctx)
	reply, err := a.data.ac.CreateArticle(ctx, &articlePb.CreateArticleRequest{
		Title:       req.Article.Title,
		Description: req.Article.Description,
		Body:        req.Article.Body,
		TagList:     req.Article.TagList,
		AuthorId:    userId,
	})
	if err != nil {
		return nil, err
	}
	return convertSingleArticle(reply), nil
}

func (a *ArticleRepo) UpdateArticle(ctx context.Context, req *interfacePb.UpdateArticleRequest) (*interfacePb.SingleArticle, error) {
	reply, err := a.data.ac.UpdateArticle(ctx, &articlePb.UpdateArticleRequest{
		Title: req.Article.Title,
	})
	if err != nil {
		return nil, err
	}
	return convertSingleArticle(reply), nil
}

func (a *ArticleRepo) DeleteArticle(ctx context.Context, req *interfacePb.DeleteArticleRequest) error {
	_, err := a.data.ac.DeleteArticle(ctx, &articlePb.DeleteArticleRequest{
		Slug: req.Slug,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepo) GetTags(ctx context.Context) ([]string, error) {
	reply, err := a.data.ac.GetTags(ctx, &empty.Empty{})
	if err != nil {
		return []string{}, err
	}
	return reply.Tags, nil
}

func (a *ArticleRepo) AddComment(ctx context.Context, slug, body string) (*interfacePb.Comment, error) {
	userId := auth.GetUserIdFromContext(ctx)
	comment, err := a.data.ac.AddComment(ctx, &articlePb.AddCommentRequest{
		Slug:   slug,
		Body:   body,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return &interfacePb.Comment{
		Id:        int32(comment.Comment.Id),
		CreatedAt: format.ConvertTime(comment.Comment.CreatedAt),
		UpdatedAt: format.ConvertTime(comment.Comment.UpdatedAt),
		Body:      comment.Comment.Body,
		Author: &interfacePb.Profile{
			Username:  comment.Comment.Author.UserName,
			Bio:       comment.Comment.Author.Bio,
			Image:     comment.Comment.Author.Image,
			Following: comment.Comment.Author.Following,
		},
	}, nil
}

func (a *ArticleRepo) GetComments(ctx context.Context, slug string) ([]*interfacePb.Comment, error) {
	res, err := a.data.ac.GetComments(ctx, &articlePb.GetCommentsRequest{Slug: slug})
	if err != nil {
		return nil, err
	}

	var comments = make([]*interfacePb.Comment, 0, len(res.Comments))
	for _, v := range res.GetComments() {
		comments = append(comments, &interfacePb.Comment{
			Id:        int32(v.Id),
			CreatedAt: format.ConvertTime(v.CreatedAt),
			UpdatedAt: format.ConvertTime(v.UpdatedAt),
			Body:      v.Body,
			Author: &interfacePb.Profile{
				Username:  v.Author.UserName,
				Bio:       v.Author.Bio,
				Image:     v.Author.Image,
				Following: v.Author.Following,
			},
		})
	}

	return comments, nil
}

func (a *ArticleRepo) DeleteComment(ctx context.Context, slug string, commentId int64) error {
	_, err := a.data.ac.DeleteComment(ctx, &articlePb.DeleteCommentRequest{
		Slug:      slug,
		CommentId: commentId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleRepo) FavoriteArticle(ctx context.Context, slug string) (*interfacePb.SingleArticle, error) {
	userId := auth.GetUserIdFromContext(ctx)
	res, err := a.data.ac.FavoriteArticle(ctx, &articlePb.FavoriteArticleRequest{
		Slug:   slug,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return convertSingleArticle(res), nil
}

func (a *ArticleRepo) UnFavoriteArticle(ctx context.Context, slug string) (*interfacePb.SingleArticle, error) {
	userId := auth.GetUserIdFromContext(ctx)
	res, err := a.data.ac.UnFavoriteArticle(ctx, &articlePb.UnFavoriteArticleRequest{
		Slug:   slug,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return convertSingleArticle(res), nil
}

func convertSingleArticle(reply *articlePb.GetArticleReply) *interfacePb.SingleArticle {
	return &interfacePb.SingleArticle{
		Slug:           reply.Article.Slug,
		Title:          reply.Article.Title,
		Description:    reply.Article.Description,
		Body:           reply.Article.Body,
		TagList:        reply.Article.TagList,
		CreatedAt:      format.ConvertTime(reply.Article.CreatedAt),
		UpdatedAt:      format.ConvertTime(reply.Article.UpdatedAt),
		Favorited:      reply.Article.Favorited,
		FavoritesCount: int32(reply.Article.FavoritesCount),
		Author: &interfacePb.Profile{
			Username:  reply.Article.Author.UserName,
			Bio:       reply.Article.Author.Bio,
			Image:     reply.Article.Author.Image,
			Following: reply.Article.Author.Following,
		},
	}
}
