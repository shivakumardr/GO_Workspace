package Saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome my dear friend James" {
		t.Error("got", s, "Expected", "Welcome my dear friend James")
	}
}
func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear friend James

}
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
