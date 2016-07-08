package main
import ("fmt"
    . "./src/github.com/matbhz/structs" // Ugly local import so you don't need to go-get this and clutter your GOPATH with this lib :)
)

func main() {
	myGrid := NewGrid(3,3);

	startingPoint := Point{0, 0};
    fmt.Println("Example from point", startingPoint)
	fmt.Println("Total possibilities:", myGrid.NumberOfPathsFromPoint(startingPoint))
    fmt.Println()
    fmt.Println("Example from all points")
    fmt.Println("Total possibilities:", myGrid.NumberOfPathsFromAllPoints())
}
