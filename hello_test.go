package main

import "testing"

//test hello function
func TestHello(t *testing.T) {

	//test for empty argument
	emptyResult := hello("")

	if emptyResult != "Hello Dude!" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Dude!", emptyResult)
	} else {
		t.Logf("hello(\"\") success, expected %v,got %v", "Hello Dude!", emptyResult)
	}

	//test for valid argument
	result := hello("Mike")

	if result != "Hello Mike!" {
		t.Errorf("hello(\"Mike\") failed, expected %v,got %v", "Hello Dude!", result)
	} else {
		t.Logf("hello(\"Mike\") success, expected %v,got %v", "Hello Dude!", result)
	}
}
