package main

import "testing_go/warehouse/notifications"

func main() {
	worker := notifications.NewRabbitMQChannel()

	worker.ProcessNotifications()
}
