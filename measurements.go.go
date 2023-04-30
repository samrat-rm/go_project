// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"unicode"
)

func trimNonSpecialCharacter(input string) (string, error) {
	var result string

	if len(input) == 7 {
		return "", fmt.Errorf("the length of the string is invalid")
	}

	for _, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			result += string(char) + "_"
		}
	}

	// Remove the last underscore, if there is one
	if len(result) > 0 && result[len(result)-1] == '_' {
		result = result[:len(result)-1]
	}

	return result, nil
}

func main() {
	input := "sa@d3f$ghff%^#vbgh*(%)"
	output, error := trimNonSpecialCharacter(input)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(output) // Output: ",_!_"

}
