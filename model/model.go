package model

import (
	"sync"
)

type BaseModel struct {
	Id uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	// CreatedAt time.Time `gorm:"column:createdAt" json:"-"`
	// UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	// DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserModel
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}

type UserInfo struct {
	UserID        uint64             `json:"user_id"`
	Username      string             `json:"username"`
	Email         string             `json:"email"`
	AvatarURL     string             `json:"avatar_url"`
	StudentID     string             `json:"student_id"`
	IsMuxiMember  bool               `json:"is_muxi_member"`
	Roles         []string           `json:"roles"`
	MemberProfile *MemberProfileInfo `json:"member_profile"`
	Info          string             `json:"info"`
}

type MemberProfileInfo struct {
	RealName     string `json:"real_name"`
	Group        string `json:"group"`
	JoinYear     int    `json:"join_year"`
	PersonalBlog string `json:"personal_blog"`
	Github       string `json:"github"`
	Zhihu        string `json:"zhihu"`
}
