package main

import (
	"log"
	"net/http"
	"testing_go/warehouse"
	"testing_go/warehouse/db"
	"testing_go/warehouse/http/controllers"
	"testing_go/warehouse/notifications"

	"github.com/gorilla/mux"
)

func main() {

	dbConnection := db.NewGormConnection()

	state := db.NewGormState(dbConnection)
	queueChannel := notifications.NewRabbitMQChannel()
	// dbChannel := notifications.NewDatabaseChannel(state)

	greg := warehouse.NewCustomer("Greg")
	bob := warehouse.NewCustomer("Bob")

	gregInterests := warehouse.NewCustomerSubscription(greg, "blue jeans")
	bobInterests := warehouse.NewCustomerSubscription(bob, "gray t-shirt")

	state.AddCustomers([]warehouse.Customer{greg, bob})
	state.AddSubscriptions([]warehouse.CustomerSubscription{gregInterests, bobInterests})

	wh := warehouse.NewWarehouse(state)

	customerNotifier := warehouse.NewCustomerNotifier(state, queueChannel)

	wh.Register(customerNotifier)

	warehouseController := controllers.NewWarehouseController(wh)

	r := mux.NewRouter()
	r.HandleFunc("/warehouse/products/list", warehouseController.ListProducts).Methods("GET")
	r.HandleFunc("/warehouse/products/new", warehouseController.NewProduct).Methods("GET")
	r.HandleFunc("/warehouse/products", warehouseController.SubmitProduct).Methods("POST")

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}

}
