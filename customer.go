package domain

type Customer struct {
	Id string `json:id`
	Name string `json:name`
	Email string `json:email`
}

type CustomerStore interface {
	Create(Customer) error
	GetAll() [] Customer
	// Update(string, Customer) error
	// Delete(string) error
	// GetById(string)(Customer, error)
}