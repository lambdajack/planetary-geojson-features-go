package region

import (
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	a, err := EuropeFranceUnmarshaled()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)

	b, err := AfricaUnmarshaled()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}