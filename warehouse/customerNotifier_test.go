package warehouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoSubscribers(t *testing.T) {

	state := NewInMemoryState()

	notifier := NewCustomerNotifier(state, state)

	notifier.Notify(NewProduct("Blue Jeans", 10))

	assert.Len(t, state.customerNotifications, 0)
}

func TestSubscriberWithWrongProduct(t *testing.T) {

	state := NewInMemoryState()

	customer := NewCustomer("Greg")

	subscription := NewCustomerSubscription(customer, "Blue t-shirt")

	state.AddSubscriptions([]CustomerSubscription{subscription})

	notifier := NewCustomerNotifier(state, state)

	notifier.Notify(NewProduct("Blue Jeans", 10))

	assert.Lenf(t, state.customerNotifications, 0, "we want that notifications be populated")
}

func TestSubscriberWithCorrectProduct(t *testing.T) {

	state := NewInMemoryState()

	customer := NewCustomer("Greg")

	subscription := NewCustomerSubscription(customer, "Blue Jeans")

	state.AddSubscriptions([]CustomerSubscription{subscription})

	notifier := NewCustomerNotifier(state, state)

	notifier.Notify(NewProduct("Blue Jeans", 10))

	assert.Len(t, state.customerNotifications, 1)
}

func TestMultipleSubscribersOneGetNotification(t *testing.T) {

	state := NewInMemoryState()

	greg := NewCustomer("Greg")
	bob := NewCustomer("Bob")

	gregsSubscription := NewCustomerSubscription(greg, "Blue Jeans")
	bobsSubscription := NewCustomerSubscription(bob, "Gray T-shirt")

	state.AddSubscriptions([]CustomerSubscription{gregsSubscription, bobsSubscription})

	notifier := NewCustomerNotifier(state, state)

	notifier.Notify(NewProduct("Blue Jeans", 10))

	assert.Len(t, state.customerNotifications, 1)
	assert.Equal(t, state.customerNotifications[0].Customer.ID, greg.ID)
}

func TestWithWarehouseObservation(t *testing.T) {

	state := NewInMemoryState()

	greg := NewCustomer("Greg")
	bob := NewCustomer("Bob")
	gregsSubscription := NewCustomerSubscription(greg, "Blue Jeans")
	bobsSubscription := NewCustomerSubscription(bob, "Gray T-shirt")

	state.AddSubscriptions([]CustomerSubscription{gregsSubscription, bobsSubscription})

	wh := NewWarehouse(state)

	notifier := NewCustomerNotifier(state, state)

	wh.Register(notifier)

	wh.AddProduct(NewProduct("Blue Jeans", 10))

	assert.Len(t, state.customerNotifications, 1)
	assert.Equal(t, state.customerNotifications[0].Customer.ID, greg.ID)
}
