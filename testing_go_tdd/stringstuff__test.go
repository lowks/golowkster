package main

import "testing"

func TestPrintStuff(t *testing.T) {
    if actual, expected := PrintSomeStuff("lowks"), "hello lowks"; actual != expected {
        t.Errorf("Expected %s to Equal %s", actual, expected)
    } else {
        t.Log("Successfully tested")
    }
}
