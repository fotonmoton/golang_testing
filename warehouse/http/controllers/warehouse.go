package controllers

import (
	"net/http"
	"strconv"
	"testing_go/warehouse"
	"testing_go/warehouse/http/views"
)

type WarehouseController struct {
	warehouse *warehouse.Warehouse
}

func NewWarehouseController(warehouse *warehouse.Warehouse) *WarehouseController {
	return &WarehouseController{warehouse}
}

func (w *WarehouseController) ListProducts(res http.ResponseWriter, req *http.Request) {
	views.HtmlProducts(res, w.warehouse.ListProducts())
}

func (w *WarehouseController) NewProduct(res http.ResponseWriter, req *http.Request) {
	views.NewHtmlProduct(res)
}

func (w *WarehouseController) SubmitProduct(res http.ResponseWriter, req *http.Request) {
	Qty, _ := strconv.ParseInt(req.PostFormValue("Qty"), 10, 64)
	Name := req.PostFormValue("Name")
	finsh := make(chan bool)
	go func() {
		w.warehouse.AddProduct(warehouse.NewProduct(Name, int(Qty)))
		finsh <- true
	}()

	<-finsh
	views.HtmlProducts(res, w.warehouse.ListProducts())
}
