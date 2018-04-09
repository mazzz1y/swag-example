package models

type UserFamilies struct {
	ID        uint `json:"id" gorm:"primary_key;unique"`
	UserID    uint `json:"owner_id"`
	FamilyID  uint `json:"family_id"`
	Confirmed bool `json:"confirmed" gorm:"not null"`
}
