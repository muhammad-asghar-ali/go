package services

import (
	"matrix/internal/pb"
	"sync"

	"github.com/jinzhu/copier"
)

type (
	LaptopStore interface {
		Save(laptop *pb.Laptop) error
	}
)

type (
	InMemoryLaptopStore struct {
		mutax sync.RWMutex
		data  map[string]*pb.Laptop
	}
)

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(lp *pb.Laptop) error {
	store.mutax.Lock()
	defer store.mutax.Unlock()

	if store.data[lp.Id] != nil {
		return ErrAlreadyExists
	}

	o := &pb.Laptop{}
	err := copier.Copy(o, lp)
	if err != nil {
		return err
	}

	store.data[o.Id] = o

	return nil
}
