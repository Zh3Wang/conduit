package articlesModel

type Articles struct {
	ID          int64  `gorm:"column:id;primary_key" json:"id"`       //
	CreatedAt   int64  `gorm:"column:created_at" json:"created_at"`   //
	UpdatedAt   int64  `gorm:"column:updated_at" json:"updated_at"`   //
	DeletedAt   int64  `gorm:"column:deleted_at" json:"deleted_at"`   //
	Slug        string `gorm:"column:slug" json:"slug"`               //
	Title       string `gorm:"column:title" json:"title"`             //
	Description string `gorm:"column:description" json:"description"` //
	Body        string `gorm:"column:body" json:"body"`               //
	AuthorID    int64  `gorm:"column:author_id" json:"author_id"`     //
}

// TableName sets the insert table name for this struct type
func (a *Articles) TableName() string {
	return "articles"
}
