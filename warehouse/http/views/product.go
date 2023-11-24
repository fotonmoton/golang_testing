package views

import (
	"html/template"
	"io"
	"testing_go/warehouse"
)

var productsTemplate = template.Must(template.ParseFiles("../../http/views/products.html"))
var newProduct = template.Must(template.ParseFiles("../../http/views/newProduct.html"))

func HtmlProducts(w io.Writer, products []warehouse.Product) {
	productsTemplate.Execute(w, products)
}

func NewHtmlProduct(w io.Writer) {
	newProduct.Execute(w, nil)
}
