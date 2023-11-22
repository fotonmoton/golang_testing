package main

import "testing/warehouse/notifications"

func main() {
	worker := notifications.NewRabbitMQChannel()

	worker.ProcessNotifications()
}
