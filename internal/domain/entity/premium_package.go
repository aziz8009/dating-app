package entity

type PremiumPackage struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
