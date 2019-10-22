package observer

import "fmt"

type ShipmentServiceInterface interface {
	Shipment()
}

type shipmentService struct {}

func NewShipmentService() ShipmentServiceInterface {
	return &shipmentService{}
}

func (shipmentService) Shipment() {
	fmt.Println("DO SHIPMENT LOGIC")
}