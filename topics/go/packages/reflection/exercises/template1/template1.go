// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type that represents a request for a customer invoice. Include a CustomerID and InvoiceID field. Define
// tags that can be used to validate the request. Define tags that specify both the length and range for the ID to be valid.
// Declare a function named validate that accepts values of any type and processes the tags. Display the results of the validation.
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Declare a struct type named Customer. Add the fields CustomerID of type int
// with the tag `length:"3" range:"100-300"`, and field InvoiceID of type int
// with tag `length:"5" range:"60000-99999"`.
type customer struct {
	customerID int `length:"3" range:"100-300"`
	invoiceID  int `length:"5" range:"60000-99999"`
}

// validate performs data validation on any struct type value.
func validate(value interface{}) {

	// Retrieve the value that the interface contains or points to.
	rv := reflect.ValueOf(value).Elem()

	// Iterate over the fields of the struct value.
	for i := 0; i < rv.NumField(); i++ {
		// Retrieve the field information.
		typeField := rv.Type().Field(i)

		// Get the value as an int, string and the length.
		valueInt := int(rv.Field(i).Int())
		valueString := strconv.Itoa(valueInt)
		valueLength := len(valueString)

		// Test the length of the value based on the tag setting.
		length, err := strconv.Atoi(typeField.Tag.Get("length"))
		if err != nil {
			panic(fmt.Errorf("invalid length definition: %w", err))
		}
		if length != valueLength {
			fmt.Printf("invalid length: field[%s] value[%d] - length[%d]\n", typeField.Name, valueInt, length)
			continue
		}

		// Test the range of the value based on the tag setting.
		splitRange := strings.Split(typeField.Tag.Get("range"), "-")
		if len(splitRange) != 2 {
			panic(fmt.Errorf("invalid range definition: %w", err))
		}
		start, err := strconv.Atoi(splitRange[0])
		if err != nil {
			panic(fmt.Errorf("invalid range definition: %w", err))
		}
		end, err := strconv.Atoi(splitRange[1])
		if err != nil {
			panic(fmt.Errorf("invalid range definition: %w", err))
		}
		if start > end {
			panic(fmt.Sprintf("invalid range definition: start[%d] can't be greater than end[%d]", start, end))
		}

		if valueInt < start || valueInt > end {
			fmt.Printf("invalid range: field[%s] value[%d] - start[%d] end[%d]\n", typeField.Name, valueInt, start, end)
			continue
		}

		fmt.Printf("valid: field[%s] - value[%d]\n", typeField.Name, valueInt)
	}
}

func main() {

	// Declare a variable of type Customer and initialize it.
	c := customer{
		customerID: 123,
		invoiceID:  99998,
	}

	// Validate the value and display the results.
	validate(&c)
}
