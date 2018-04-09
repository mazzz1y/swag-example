package models

type User struct {
	ID            uint     `json:"id" gorm:"primary_key;unique"`
	Email         string   `json:"email" gorm:"unique_index"`
	PassHash      string   `json:"pass_hash"`
	Role          string   `json:"role"`
	Wallets       []Wallet `gorm:"foreignkey:OwnerID;association_foreignkey:ID"`
	FamiliesOwner []Family `json:"families" gorm:"foreignkey:OwnerID;association_foreignkey:ID"`
	Families      []Family `gorm:"many2many:user_families;"`
}

func (User) TableName() string {
	return "users"
}

func (u User) Update() (User, error) {
	err := DB.Save(&u).Error
	return u, err
}
