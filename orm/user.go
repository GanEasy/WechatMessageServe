package orm

import "time"

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	ID         uint   `gorm:"primary_key"`
	Email      string `gorm:"type:varchar(100);unique_index"` // `type` set sql type, `unique_index` will create unique index for this column
	OpenID     string `gorm:"size:255;sql:index"`
	Invited    bool
	Registered bool
	Subscribed bool `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

// GetUserByID 获取用户关注的书籍
func (user *User) GetUserByID(id int) {
	DB().First(user, id)
}

// GetUserByEmail 通过openID获取用户信息 如果没有的话进行初始化
func (user *User) GetUserByEmail(email string) {
	DB().Where(User{Email: email}).FirstOrCreate(user)
}

// GetUserByOpenID 通过openID获取用户信息 如果没有的话进行初始化
func (user *User) GetUserByOpenID(openID string) {
	DB().Where(User{OpenID: openID}).FirstOrInit(user)
}

// Save 保存用户信息
func (user *User) Save() {
	DB().Save(&user)
}
