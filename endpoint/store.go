package endpoint

import "time"

type Category struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserCategory struct {
	Id int
	UserId int
	CategoryId int
}

type Product struct {
	Id             int       `json:"id"`
	CategoryId     int       `json:"category_id"`
	Image          string    `json:"image"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Cost           float64   `json:"cost"`
	CreatedCompany string    `json:"created_company"`
	CreatedCountry string    `json:"created_country"`
	CreatedDate    time.Time `json:"created_date"`
}

type CategoryProduct struct {
	Id int
	CategoryId int
	ProductId int
}

type CartProduct struct {
	Id        int     `json:"id"`
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}

type Cart struct {
	Id            int       `json:"id"`
	CartProductId int       `json:"cart_product_id"`
	CreatedAt     time.Time `json:"created_at"`
	Total         float64   `json:"total"`
}

type Order struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	CartId    int    `json:"cart_id"`
	Status    string `json:"status"`
	Confirmed bool   `json:"confirmed"`
}

type Review struct {
	Id         int    `json:"id"`
	OrderId    int    `json:"order_id"`
	Image      string `json:"image"`
	ReviewText string `json:"review_text"`
	Rating     int    `json:"rating"`
}
