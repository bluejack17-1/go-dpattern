package structural

import "fmt"

type Endpoint func ()

func LogDecorator(fn Endpoint) Endpoint {
	return func() {
		fmt.Println("Start")

		fn()

		fmt.Println("End")
	}
}