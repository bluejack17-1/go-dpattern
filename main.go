package main

import (
	"fmt"
	"github.com/bluejack17-1/go-dpattern/behavioral"
	"github.com/bluejack17-1/go-dpattern/behavioral/observer"
	"github.com/bluejack17-1/go-dpattern/creational"
	"github.com/bluejack17-1/go-dpattern/structural"
	adapter2 "github.com/bluejack17-1/go-dpattern/structural/adapter"
)

func main() {
	//singleton()
	//builder()
	//factory()
	//adapter()
	//strategy()
	//decorator()
	//templateMethod()
	observable()
}

func singleton() {
	a := creational.GetInstance()
	b := creational.GetInstance()

	fmt.Println(&a, &b)
}

func builder() {
	q := &creational.Query{}

	qs := q.
		From("users").
		Select("id", "name", "role").
		Where("role", "=", "creator").
		OrWhere("name", "ILIKE", "a%").
		Build()

	fmt.Println(qs)
}

func factory() {
	storage := creational.NewStorage("postgres")

	storage.Open()
}

func adapter() {
	cart1 := adapter2.ShopingCart{
		PaymentMethod: adapter2.PaypalAdapter{Paypal: adapter2.PaypalPayment{}},
	}

	cart2 := adapter2.ShopingCart{
		PaymentMethod: adapter2.BankAdapter{Bank: adapter2.BankPayment{}},
	}

	cart1.PaymentMethod.Pay("A", "B", 10000)
	fmt.Println("")
	cart2.PaymentMethod.Pay("C", "D", 50000)
}

func strategy() {
	impl := behavioral.Implementation{
		ReportStrat: behavioral.PDF{},
	}

	impl.Export()
}

func decorator() {
	structural.LogDecorator(func () {
		for i := 0 ; i < 100 ; i++ {
			fmt.Println( i )
		}
	})()
}

func templateMethod() {
	exporter := behavioral.NewExporter(behavioral.TemplateImpl{})
	exporter.ExportReport()
}

func observable() {
	service := observer.NewCheckoutService()

	service.Checkout()
}