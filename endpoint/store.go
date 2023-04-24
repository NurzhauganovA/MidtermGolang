package endpoint

import "time"

type Category struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserCategory struct {
	Id         int
	UserId     int
	CategoryId int
}

type Rating struct {
	Id        int `json:"id" db:"id"`
	UserId    int `json:"user_id" db:"user_id"`
	ProductId int `json:"product_id" db:"product_id"`
	Rating    int `json:"rating" db:"rating"`
}

type Comment struct {
	Id        int    `json:"id" db:"id"`
	UserId    int    `json:"user_id" db:"user_id"`
	ProductId int    `json:"product_id" db:"product_id"`
	Comment   string `json:"comment" db:"comment"`
}

type Product struct {
	Id             int       `json:"id" db:"id"`
	Image          string    `json:"image" db:"image"`
	Title          string    `json:"title" db:"title" binding:"required"`
	Description    string    `json:"description" db:"description"`
	Cost           float64   `json:"cost" db:"cost"`
	CreatedCompany string    `json:"created_company" db:"created_company"`
	CreatedCountry string    `json:"created_country" db:"created_country"`
	CreatedDate    time.Time `json:"created_date" db:"created_date"`
	Ratings        []Rating  `json:"ratings,omitempty" db:"-"`  // add ratings field
	Comments       []Comment `json:"comments,omitempty" db:"-"` // add comments field
	Rating         int
	Comment        string
}

func (p Product) Error() string {
	//TODO implement me
	panic("implement me")
}

type CategoryProduct struct {
	Id         int
	CategoryId int
	ProductId  int
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
