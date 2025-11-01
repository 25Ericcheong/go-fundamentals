package go_fundamentals

import (
	dependencyinjection "go-fundamentals/dependency-injection"
	"os"
)

func main() {
	dependencyinjection.Greet(os.Stdout, "test")
}
