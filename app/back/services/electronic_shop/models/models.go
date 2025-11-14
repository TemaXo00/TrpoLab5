package models

import (
	"time"
)

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"category_id"`
	Name string `gorm:"size:50;not null" json:"category_name"`
}

type Supplier struct {
	ID     uint    `gorm:"primaryKey" json:"supplier_id"`
	Name   string  `gorm:"size:50;not null" json:"supplier_name"`
	Phone  string  `gorm:"size:30;not null" json:"phone"`
	Email  string  `gorm:"size:90" json:"email"`
	Rating float64 `gorm:"type:decimal(2,1);not null" json:"rating"`
}

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"product_id"`
	Name        string    `gorm:"size:100;not null" json:"product_name"`
	Description string    `gorm:"type:text" json:"description"`
	CategoryID  uint      `gorm:"not null" json:"category_id"`
	SupplierID  uint      `gorm:"not null" json:"supplier_id"`
	Price       float64   `gorm:"type:decimal(10,1);not null" json:"price"`
	Rating      float64   `gorm:"type:decimal(2,1)" json:"rating"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at"`
	Category    *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Supplier    *Supplier `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
}

type Position struct {
	ID       uint   `gorm:"primaryKey" json:"position_id"`
	Name     string `gorm:"size:50;not null" json:"position_name"`
	Category string `gorm:"size:50;not null" json:"position_category"`
}

type Store struct {
	ID          uint   `gorm:"primaryKey" json:"store_id"`
	Name        string `gorm:"size:100;not null" json:"store_name"`
	Address     string `gorm:"type:text;not null" json:"address"`
	Email       string `gorm:"size:90" json:"email"`
	OpeningTime string `gorm:"type:time;not null" json:"opening_time"`
	ClosingTime string `gorm:"type:time;not null" json:"closing_time"`
}

type Employee struct {
	ID         uint      `gorm:"primaryKey" json:"employee_id"`
	LastName   string    `gorm:"size:50;not null" json:"last_name"`
	FirstName  string    `gorm:"size:50;not null" json:"first_name"`
	FatherName string    `gorm:"size:50" json:"father_name"`
	PositionID uint      `gorm:"not null" json:"position_id"`
	StoreID    uint      `gorm:"not null" json:"store_id"`
	Phone      string    `gorm:"size:30;not null" json:"phone"`
	Email      string    `gorm:"size:90" json:"email"`
	HireDate   string    `gorm:"type:date;not null" json:"hire_date"`
	Position   *Position `gorm:"foreignKey:PositionID" json:"position,omitempty"`
	Store      *Store    `gorm:"foreignKey:StoreID" json:"store,omitempty"`
}

type Client struct {
	ID               uint      `gorm:"primaryKey" json:"client_id"`
	FirstName        string    `gorm:"size:50;not null" json:"first_name"`
	LastName         string    `gorm:"size:50;not null" json:"last_name"`
	Phone            string    `gorm:"size:30;not null" json:"phone"`
	Email            string    `gorm:"size:90" json:"email"`
	Address          string    `gorm:"type:text" json:"address"`
	RegistrationDate time.Time `gorm:"default:current_timestamp" json:"registration_date"`
}

type Sale struct {
	ID            uint        `gorm:"primaryKey" json:"sale_id"`
	StoreID       uint        `gorm:"not null" json:"store_id"`
	CustomerID    uint        `gorm:"not null" json:"customer_id"`
	EmployeeID    uint        `gorm:"not null" json:"employee_id"`
	SaleDate      time.Time   `gorm:"default:current_timestamp" json:"sale_date"`
	PaymentMethod string      `gorm:"size:50;not null" json:"payment_method"`
	Status        string      `gorm:"size:50;not null" json:"status"`
	Store         *Store      `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	Customer      *Client     `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Employee      *Employee   `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	SaleItems     []*SaleItem `gorm:"foreignKey:SaleID" json:"sale_items,omitempty"`
}

type SaleItem struct {
	ID        uint     `gorm:"primaryKey" json:"sale_item_id"`
	SaleID    uint     `gorm:"not null" json:"sale_id"`
	ProductID uint     `gorm:"not null" json:"product_id"`
	Quantity  int      `gorm:"not null" json:"quantity"`
	UnitPrice float64  `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Sale      *Sale    `gorm:"foreignKey:SaleID" json:"sale,omitempty"`
}

type Review struct {
	ID         uint      `gorm:"primaryKey" json:"review_id"`
	CustomerID uint      `gorm:"not null" json:"customer_id"`
	ProductID  uint      `gorm:"not null" json:"product_id"`
	EmployeeID uint      `gorm:"not null" json:"employee_id"`
	StoreID    uint      `gorm:"not null" json:"store_id"`
	Rating     int       `gorm:"not null" json:"rating"`
	ReviewText string    `gorm:"type:text" json:"review_text"`
	ReviewDate time.Time `gorm:"default:current_timestamp" json:"review_date"`
	Customer   *Client   `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Product    *Product  `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	Store      *Store    `gorm:"foreignKey:StoreID" json:"store,omitempty"`
}
