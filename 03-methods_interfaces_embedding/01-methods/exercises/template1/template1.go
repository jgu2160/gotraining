// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Rj0QfwVPhX

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initalize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main
import "fmt"

// Add imports.
type player struct {
    name string
    atBats float64
    hits float64
}

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.

// Declare a method that calculates the batting average for a batter.
func (p *player) average() float64 {
    return p.hits/p.atBats
}

// main is the entry point for the application.
func main() {

	// Create a slice of players and populate each player
	// with field values.
    players := []player{
        {"yogi berra", 20, 15},
        {"mace windu", 7, 3},
    }

	// Display the batting average for each player in the slice.
  for _, player := range players {
      average := player.average() * 1000
      fmt.Printf("player: %s avg: 0.%.f \n", player.name, average)
  }
}
