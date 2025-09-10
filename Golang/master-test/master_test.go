package master

import (
	"testing"
)

func TestADD(t *testing.T) {

	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2,3) failed, expected %d, got %d", expected, result)
	}
}

func TestEven(t *testing.T) {

	if !IsEven(4) {
		t.Errorf("IsEven(4) failed: expected true,but false")
	}


}
