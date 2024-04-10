package dto

type CustomerDTO struct {
	Name      string     `json:"name"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	LoggedIn  bool       `json:"loggedin"`
	Orders    []OrderDTO `json:"orders"`
}
