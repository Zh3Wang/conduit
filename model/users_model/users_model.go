package usersModel

type Users struct {
	ID           int64  `gorm:"column:id;primary_key" json:"id"`           //
	CreatedAt    int64  `gorm:"column:created_at" json:"created_at"`       //
	UpdatedAt    int64  `gorm:"column:updated_at" json:"updated_at"`       //
	DeletedAt    int64  `gorm:"column:deleted_at" json:"deleted_at"`       //
	Email        string `gorm:"column:email" json:"email"`                 //
	Username     string `gorm:"column:username" json:"username"`           //
	Bio          string `gorm:"column:bio" json:"bio"`                     //
	Image        string `gorm:"column:image" json:"image"`                 //
	PasswordHash string `gorm:"column:password_hash" json:"password_hash"` //
	Following    int    `gorm:"column:following" json:"following"`         //
}

// TableName sets the insert table name for this struct type
func (u *Users) TableName() string {
	return "users"
}
