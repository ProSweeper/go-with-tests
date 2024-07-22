package intergers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(3, 2)
	fmt.Println(sum)
	// Output: 5
}
