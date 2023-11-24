package views

import (
	"strings"
	"testing"
	"testing_go/warehouse"

	"github.com/stretchr/testify/assert"
)

func TestEmptyHtmlProductsTemplate(t *testing.T) {

	rendered := strings.Builder{}

	wanted := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Products</title>
  </head>
  <body>
    <h2>Products</h2>
    <a href="/warehouse/products/new">Add new!</a>
    <ul>
      
    </ul>
  </body>
</html>
`

	HtmlProducts(&rendered, []warehouse.Product{})

	assert.Equal(t, wanted, rendered.String())
}

func TestOneProductHtmlProductsTemplate(t *testing.T) {

	rendered := strings.Builder{}

	wanted := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Products</title>
  </head>
  <body>
    <h2>Products</h2>
    <a href="/warehouse/products/new">Add new!</a>
    <ul>
      
      <li>Blue Jeans - 2</li>
      
    </ul>
  </body>
</html>
`

	products := []warehouse.Product{warehouse.NewProduct("Blue Jeans", 2)}

	HtmlProducts(&rendered, products)

	assert.Equal(t, wanted, rendered.String())
}
