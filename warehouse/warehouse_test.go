package warehouse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWarehouseAddProduct(t *testing.T) {

	warehouse := NewWarehouse(NewInMemoryState())

	warehouse.AddProduct(NewProduct("Blue Jeans", 1))

	assert.Equal(t, warehouse.ListProducts()[0].Name, "Blue Jeans")
	assert.Equal(t, warehouse.ListProducts()[0].Qty, 1)
}

type testObserver struct {
	observations []any
}

func (o *testObserver) Observe(sbj any) {
	o.observations = append(o.observations, sbj)
}

func TestWarehouseNotifications(t *testing.T) {

	warehouse := NewWarehouse(NewInMemoryState())

	observer := &testObserver{}

	warehouse.Register(observer)

	warehouse.AddProduct(NewProduct("Blue Jeans", 1))

	assert.Len(t, observer.observations, 1)
	assert.Equal(t, observer.observations[0].(Product).Name, "Blue Jeans")
}

func BenchmarkAddProductWithoutSubscribers(b *testing.B) {
	warehouse := NewWarehouse(NewInMemoryState())

	for i := 0; i < b.N; i++ {
		warehouse.AddProduct(NewProduct("Blue Jeans", 1))
	}
}

func BenchmarkAddProductWith10Subscribers(b *testing.B) {
	warehouse := NewWarehouse(NewInMemoryState())

	for range make([]struct{}, 10) {
		warehouse.Register(&testObserver{})
	}

	for i := 0; i < b.N; i++ {
		warehouse.AddProduct(NewProduct("Blue Jeans", 1))
	}
}

func BenchmarkAddProductWith1000Subscribers(b *testing.B) {
	warehouse := NewWarehouse(NewInMemoryState())

	for range make([]struct{}, 1000) {
		warehouse.Register(&testObserver{})
	}

	for i := 0; i < b.N; i++ {
		warehouse.AddProduct(NewProduct("Blue Jeans", 1))
	}
}

func ExampleWarehouse() {
	warehouse := NewWarehouse(NewInMemoryState())
	warehouse.AddProduct(NewProduct("Blue Jeans", 1))
	fmt.Printf("%+v", warehouse.ListProducts())
	// Output: [{ID:1 Name:Blue Jeans Qty:1 Active:true}]
}
