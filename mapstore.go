package mapstore

import (
	"github.com/shijuvar/gokit/examples/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore{
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (s *MapStore) Create(cust domain.Customer) error {
	s.store[cust.Id] = cust
	return nil
}
func (s *MapStore) GetAll() []domain.Customer{
	var customers []domain.Customer
	for _ , v := range s.store{
		customers = append ( customers, v)
	}
	return customers
}
