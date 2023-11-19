package structs
import (
	"time"
	"ecomerce/model"
)

// type User struct {
// 	Id       int
// 	Nama     string
// 	Username string
// 	Email    string
// 	Password string
// 	Auth     string
// }

// type Role struct {
// 	Id          int
// 	Role        string
// 	Description string
// }

// type Previlage struct {
// 	Id          int
// 	Previlage   string
// 	Description string
// }

// type RolePrevilage struct {
// 	RoeleId     int
// 	PrevilageId int
// }

// type UserRole struct {
// 	UserId int
// 	RoleId int
// }

// type BlogCategory struct {
// 	Id       int
// 	Category string
// }

// type Blog struct {
// 	Id         int
// 	Judul      string
// 	Body       string
// 	Foto       string
// 	CreatedBy  string
// 	CreatedAt  date
// 	UpdatedBy  string
// 	UpdatedAt  date
// 	categoryId int
// }

// type Catalog struct {
// 	Id          int
// 	ProductName string
// 	Harga       int
// 	Detail      string
// 	InfoPenting string
// 	Foto        string
// 	SisaStok    int
// 	CreatedBy   string
// 	CreatedAt   date
// 	UpdatedBy   string
// 	UpdatedAt   date
// }

// type CategoryProduct struct {
// 	Id      int
// 	Product string
// }

// type Merek struct {
// 	Id        int
// 	Merek     string
// 	ProductId int
// }

// type About struct {
// }

type UserDataResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Roles     []model.Role   `json:"roles"`
	CreatedAt time.Time `json:"created_At"`
	Status bool `json:"status"`
}
