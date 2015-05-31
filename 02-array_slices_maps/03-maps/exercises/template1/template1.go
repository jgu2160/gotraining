// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/-JBSUoux-v

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import "fmt"

// main is the entry point for the application.
func main() {
	// Declare and make a map of integer type values.
  ints := make(map[int]int)
	// Intialize some data into the map.
    ints[1] = 1
    ints[2] = 2
    ints[3] = 3
    ints[4] = 4
	// Display each key/value pair.
  for i, value := range ints {
      fmt.Printf("key: %d value: %d\n", i, value)
  }
}
