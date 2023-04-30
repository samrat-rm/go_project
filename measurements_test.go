
package main 

import "testing"

func TestStringConversion(t *testing.T) {
	input := "sample*(1%23@s#tr*i^ng"
	output := "*_(_%_@_#_*_^"
	acutal , err := trimNonSpecialCharacter(input)
	if err != nil  {
		t.Errorf("Expected no error but got %v" , err)
	}
	if acutal != output {
		t.Errorf("Expected %s output but got %s " , output , acutal)
	}

}