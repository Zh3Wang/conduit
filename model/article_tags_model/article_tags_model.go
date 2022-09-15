package articleTagsModel

type ArticleTags struct {
	TagID     int64 `gorm:"column:tag_id;primary_key" json:"tag_id"`         //
	ArticleID int64 `gorm:"column:article_id;primary_key" json:"article_id"` //
}

// TableName sets the insert table name for this struct type
func (a *ArticleTags) TableName() string {
	return "article_tags"
}
