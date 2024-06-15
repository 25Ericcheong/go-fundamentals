package integers

import (
	"fmt"
	"go-fundamentals/utils"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	utils.AssertCorrectNumbersMessage(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
