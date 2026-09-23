package model

import (
	"time"

	"github.com/Muxi-X/muxi_auth_service_v2/pkg/constvar"
	"github.com/jinzhu/gorm"
)

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

// Create inserts a new member profile.
func (profile *MemberProfile) Create() error {
	return DB.Self.Create(profile).Error
}

// Update saves the mutable fields of an existing member profile.
func (profile *MemberProfile) Update() error {
	return DB.Self.Save(profile).Error
}

// GetMemberProfileByID returns the member profile by its primary key.
func GetMemberProfileByID(id uint64) (*MemberProfile, error) {
	profile := &MemberProfile{}
	d := DB.Self.Where("id = ?", id).First(profile)
	return profile, d.Error
}

// GetMemberProfileByUserID returns the member profile bound to a local user.
func GetMemberProfileByUserID(userID uint64) (*MemberProfile, error) {
	profile := &MemberProfile{}
	d := DB.Self.Where("user_id = ?", userID).First(profile)
	return profile, d.Error
}

// ListMemberProfiles returns a page of member profiles and the total count.
func ListMemberProfiles(offset, limit int) ([]*MemberProfile, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	var count uint64
	if err := DB.Self.Model(&MemberProfile{}).Count(&count).Error; err != nil {
		return nil, count, err
	}

	profiles := make([]*MemberProfile, 0, limit)
	if err := DB.Self.Offset(offset).Limit(limit).Find(&profiles).Error; err != nil {
		return nil, count, err
	}
	return profiles, count, nil
}

// DeleteMemberProfileByUserID removes the member profile bound to a local user.
// It returns gorm.ErrRecordNotFound when no profile exists for the user.
func DeleteMemberProfileByUserID(userID uint64) error {
	res := DB.Self.Where("user_id = ?", userID).Delete(&MemberProfile{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
