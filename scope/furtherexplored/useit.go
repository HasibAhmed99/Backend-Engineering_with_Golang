package furtherexplored

import "fmt"

// this is also "package block" scope
// this is exported because the identifier "Fasincating" is capitalized

func Fascinating() {
	fmt.Println("Planet speed:", planetSpeed)

	planetRotationSpeed := 1000
	fmt.Println("Planer spinning speed:", planetRotationSpeed)
}
