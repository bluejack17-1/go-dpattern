package observer

import "fmt"

type shipmentPaymentSubscriber struct {
	shipmentService ShipmentServiceInterface
}

func NewShipmentSubscriber() ISubscriber {
	return &shipmentPaymentSubscriber{
		shipmentService: NewShipmentService(),
	}
}

func (ps shipmentPaymentSubscriber) OnSubscribe() {
	fmt.Println("FROM SHIPMENT SUBSCRIBE")

	ps.shipmentService.Shipment()
}
