package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"testing_go/warehouse"

	"github.com/stretchr/testify/assert"
)

func TestSubmitNewProduct(t *testing.T) {

	c := NewWarehouseController(warehouse.NewWarehouse(warehouse.NewInMemoryState()))

	form := url.Values{
		"Name": {"Blue Jeans"},
		"Qty":  {"10"},
	}

	req, err := http.NewRequest("POST", "/warehouse/products", strings.NewReader(form.Encode()))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res := httptest.NewRecorder()

	c.SubmitProduct(res, req)

	t.Run("another test", func(t *testing.T) {
		t.Parallel()
	})

	products := c.warehouse.ListProducts()

	created := products[0]

	assert.Equal(t, res.Code, http.StatusOK)
	assert.Len(t, products, 1)
	assert.Equal(t, "Blue Jeans", created.Name)
	assert.Equal(t, 10, created.Qty)
}
