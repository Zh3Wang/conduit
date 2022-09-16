package commentsModel

type Comments struct {
	ID          int64  `gorm:"column:id;primary_key" json:"id"`         //
	CreatedAt   int64  `gorm:"column:created_at" json:"created_at"`     //
	UpdatedAt   int64  `gorm:"column:updated_at" json:"updated_at"`     //
	DeletedAt   int64  `gorm:"column:deleted_at" json:"deleted_at"`     //
	ArticleSlug string `gorm:"column:article_slug" json:"article_slug"` //
	Body        string `gorm:"column:body" json:"body"`                 //
	AuthorID    int64  `gorm:"column:author_id" json:"author_id"`       //
}

// TableName sets the insert table name for this struct type
func (c *Comments) TableName() string {
	return "comments"
}
