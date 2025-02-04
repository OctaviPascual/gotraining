// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Setup a new program to use the log package. Set the Prefix to your first name and on each log line show
// the date and long path for the code file.
package main

import (
	"fmt"
	"log"
	"os"
)

func init() {

	// Change the output device from the default stderr to stdout.
	log.SetOutput(os.Stdout)

	// Set the prefix string for each log line.
	log.SetPrefix("OCTAVI: ")

	// Set the extra log info.
	setFlags()
}

// setFlags adds extra information on each log line.
func setFlags() {
	/*
	   Ldate		   // the date: 2009/01/23
	   Ltime           // the time: 01:23:23
	   Lmicroseconds   // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	   Llongfile       // full file name and line number: /a/b/c/d.go:23
	   Lshortfile      // final file name element and line number: d.go:23. overrides Llongfile
	   LstdFlags       // Ldate | Ltime // initial values for the standard logger
	*/

	// Set the flags.
	log.SetFlags(log.Ldate | log.Llongfile)
}

func main() {

	// Use the Println function.
	fmt.Println("Hello world!")

	// Create a slice of strings and initialize with names.
	// Log the values of the slice.
	names := []string{"Alice", "Bob", "Charlie"}
	log.Println(names)

	// Use the Fatalln function.
	log.Fatalln("Ending program")
}
