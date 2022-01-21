package iteration

import "strings"


func Repeate(character string, repetitions int) string {

	var returnValue string


	
	defer func() {

		if r := recover(); r != nil {
			returnValue = "";
		} 
		
	} ()
	returnValue = strings.Repeat(character, repetitions)

	return returnValue

}