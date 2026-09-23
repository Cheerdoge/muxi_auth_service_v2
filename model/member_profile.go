package model

import "time"

type MemberProfile struct {
	BaseModel
	UserID       uint64    `json:"user_id" column:"user_id"`
	RealName     string    `json:"real_name" column:"real_name"`
	StudentID    string    `json:"student_id" column:"student_id"`
	Group        string    `json:"group" column:"group"`
	JoinYear     int       `json:"join_year" column:"join_year"`
	PersonalBlog string    `json:"personal_blog" column:"personal_blog"`
	Github       string    `json:"github" column:"github"`
	Zhihu        string    `json:"zhihu" column:"zhihu"`
	CreatedAt    time.Time `json:"created_at" column:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" column:"updated_at"`
}

func (MemberProfile) TableName() string {
	return "member_profiles"
}
