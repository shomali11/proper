package main

import (
	"fmt"
	"github.com/shomali11/proper"
)

func main() {
	parameters := make(map[string]string)
	parameters["boolean"] = "true"
	parameters["float"] = "1.2"
	parameters["integer"] = "11"
	parameters["string"] = "value"

	properties := proper.NewProperties(parameters)

	fmt.Println(properties.BooleanParam("boolean", false))
	fmt.Println(properties.FloatParam("float", 0))
	fmt.Println(properties.IntegerParam("integer", 0))
	fmt.Println(properties.StringParam("string", ""))
}
