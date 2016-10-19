package hello

import "testing"

func TestHello(t *testing.T){
    expected := "Hello Go xx!"
    actual := hello("Hello Go xx!")
    if actual != expected {
        t.Errorf("Test failed, expected '%s', got: '%s'", expected, actual)
    }
}
