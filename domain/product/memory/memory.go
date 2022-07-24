package memory

import (
	"github.com/google/uuid"
	"sync"

	"github.com/Sraik25/ddd-go/aggregate"
	"github.com/Sraik25/ddd-go/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: map[uuid.UUID]aggregate.Product{},
	}
}

// GetAll returns all products as a slice
// Yes, it never returns an error, but
// A database implementation could return an error for instance
func (mrp *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range mrp.products {
		products = append(products, product)
	}
	return products, nil
}

// GetByID searches for a product based on it's ID
func (mrp *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mrp.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

// Add will add a new product to the repository
func (mrp *MemoryProductRepository) Add(newProduct aggregate.Product) error {
	mrp.Lock()
	defer mrp.Unlock()

	if _, ok := mrp.products[newProduct.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}
	mrp.products[newProduct.GetID()] = newProduct
	return nil
}

// Update will change all values for a product based on it's ID
func (mrp *MemoryProductRepository) Update(uppprod aggregate.Product) error {
	mrp.Lock()
	defer mrp.Unlock()

	if _, ok := mrp.products[uppprod.GetID()]; !ok {
		return product.ErrProductNotFound
	}
	mrp.products[uppprod.GetID()] = uppprod
	return nil
}

// Delete remove an product from the repository
func (mrp *MemoryProductRepository) Delete(id uuid.UUID) error {
	mrp.Lock()
	defer mrp.Unlock()

	if _, ok := mrp.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mrp.products, id)
	return nil
}
