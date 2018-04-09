package models

type Family struct {
	ID      uint   `json:"id" gorm:"primary_key;unique"`
	OwnerID uint   `json:"owner_id"`
	Name    string `json:"name"`
	Deleted bool   `json:"is_deleted"`
	Users   []User `gorm:"many2many:user_families;"`
}

func (Family) TableName() string {
	return "families"
}
