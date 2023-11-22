package notifications

import (
	"log"
	"testing/warehouse"
	"testing/warehouse/db"
)

type DatabaseChannel struct {
	db *db.GormState
}

func NewDatabaseChannel(db *db.GormState) *DatabaseChannel {
	return &DatabaseChannel{db}
}

func (ch *DatabaseChannel) NotifyCustomers(notifications []warehouse.CustomerNotification) {
	ch.db.NotifyCustomers(notifications)
	log.Println("Customers notifications has been written!")
}
