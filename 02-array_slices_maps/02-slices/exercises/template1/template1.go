// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/mPKmyGNR2L

// Declare a nil slice of integers. Declare a nil slice of integers. Create a
// loop that appends 10 values to the slice. Iterate over the slice and display
// each value. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
  var aSlice []int
	// Appends numbers to the slice.
    aSlice = append(aSlice, 6)
    aSlice = append(aSlice, 7)
    aSlice = append(aSlice, 9)
	// Display each value in the slice.
  for _,name := range aSlice {
     fmt.Printf("%d\n", name)
  }
	// Declare a slice of strings and populate the slice with names.
  stringSlice := []string{"hello", "moto", "goober", "boober"}
	// Display each index position and slice value.
  for _, value := range stringSlice {
      fmt.Printf(value + "\n")
  }
	// // Take a slice of index 1 and 2 of the slice of strings.
  slice3 := stringSlice[1:2]
  slice4 := stringSlice[2:4]
	// Display each index position and slice values for the new slice.
  for i, value := range slice3 {
      fmt.Printf("%d %s\n", i, value)
  }

  for i, value := range slice4 {
      fmt.Printf("index: %d name: %s\n", i, value)
  }
}
