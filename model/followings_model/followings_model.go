package followingsModel

type Followings struct {
	ID          int64 `gorm:"column:id;primary_key" json:"id"`         //
	CreatedAt   int64 `gorm:"column:created_at" json:"created_at"`     //
	UpdatedAt   int64 `gorm:"column:updated_at" json:"updated_at"`     //
	DeletedAt   int64 `gorm:"column:deleted_at" json:"deleted_at"`     //
	UserID      int64 `gorm:"column:user_id" json:"user_id"`           //
	FollowingID int64 `gorm:"column:following_id" json:"following_id"` //
}

// TableName sets the insert table name for this struct type
func (f *Followings) TableName() string {
	return "followings"
}
