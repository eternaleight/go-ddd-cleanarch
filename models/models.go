package models

import "time"

type Post struct {
	ID        uint `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	AuthorID  uint      `gorm:"column:authorId"`
}

// TableName overrides the table name
func (Post) TableName() string {
	return "Post"
}

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Posts    []Post `gorm:"foreignKey:AuthorID"`
	Profile  Profile
}

func (User) TableName() string {
	return "User"
}

type Profile struct {
	ID              uint   `gorm:"primaryKey"`
	Bio             string `gorm:"size:1000"`
	ProfileImageUrl string `gorm:"column:profileImageUrl"`
	UserID          uint   `gorm:"column:userId"`
}

func (Profile) TableName() string {
	return "Profile"
}

type Product struct {
	ID          uint       `gorm:"primarykey;column:id"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	Price       int        `gorm:"column:price"`
	ImageURL    string     `gorm:"column:imageUrl"`
	VideoURL    string     `gorm:"column:videoUrl"`
	CreatedAt   time.Time  `gorm:"column:createdAt"`
	SellerID    uint       `gorm:"column:sellerId"`
	Seller      User       `gorm:"foreignKey:SellerID;references:ID"`
	Purchases   []Purchase `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "Product"
}

type Purchase struct {
	ID              uint      `gorm:"primarykey;column:id"`
	ProductID       uint      `gorm:"column:productId"`
	BuyerID         uint      `gorm:"column:buyerId"`
	PurchaseDate    time.Time `gorm:"column:purchaseDate"`
	StripePaymentID string    `gorm:"column:stripePaymentId"`
}

func (Purchase) TableName() string {
	return "Purchase"
}
