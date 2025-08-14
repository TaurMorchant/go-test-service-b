package go_test_service_b

import (
	"fmt"
	gotestservicea "github.com/taurmorchant/go-test-service-a"
)

func Hello() {
	fmt.Println("Hello, World from go-test-service-a")
	gotestservicea.Hello()
}
