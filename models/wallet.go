package models

type Wallet struct {
	ID        uint    `json:"id" gorm:"primary_key;unique"`
	OwnerID   uint    `json:"owner_id"`
	FamilyID  uint    `json:"family_id"`
	Value     float64 `json:"value"`
	IsDeleted bool    `json:"is_deleted"`
}

func (Wallet) TableName() string {
	return "wallets"
}
