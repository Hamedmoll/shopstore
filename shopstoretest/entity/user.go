package entity

type User struct {
	ID          uint   `json:"id"`
	Role        Role   `json:"role"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Credit      int    `json:"credit"`
	PhoneNumber string `json:"phone_number"`
}
