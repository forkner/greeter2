package greeter1

import (
	"fmt"

	"github.com/forkner/greetings"
)

func SayHello() {
	fmt.Printf("Greeter2: %s!\n", greetings.Greeting)
}
