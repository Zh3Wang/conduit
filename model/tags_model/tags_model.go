package tagsModel

type Tags struct {
	ID        int64  `gorm:"column:id;primary_key" json:"id"`     //
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` //
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` //
	DeletedAt int64  `gorm:"column:deleted_at" json:"deleted_at"` //
	Name      string `gorm:"column:name" json:"name"`             //
}

// TableName sets the insert table name for this struct type
func (t *Tags) TableName() string {
	return "tags"
}
