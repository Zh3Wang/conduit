package articleFavModel

type ArticleFavorites struct {
	ID        int64 `gorm:"column:id;primary_key" json:"id"`     //
	CreatedAt int64 `gorm:"column:created_at" json:"created_at"` //
	UpdatedAt int64 `gorm:"column:updated_at" json:"updated_at"` //
	DeletedAt int64 `gorm:"column:deleted_at" json:"deleted_at"` //
	UserID    int64 `gorm:"column:user_id" json:"user_id"`       //
	ArticleID int64 `gorm:"column:article_id" json:"article_id"` //
}

// TableName sets the insert table name for this struct type
func (a *ArticleFavorites) TableName() string {
	return "article_favorites"
}
