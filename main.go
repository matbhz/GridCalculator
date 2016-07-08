package main
import "fmt"

func main() {
	myGrid := NewGrid(3,3);

	startingPoint := Point{0, 0};
	fmt.Println(myGrid.NumberOfPathsFromOrigin(startingPoint))
}

type Grid struct {
	Grid [][]*Point
	ValidPaths int;
}

func NewGrid(rows, columns int) *Grid {

	grid := [][]*Point{}

	for i := 0; i < rows; i++ {
		row := make([]*Point, columns);
		grid = append(grid, row);
	}

	return &Grid{Grid: grid}
}

func (g *Grid) NumberOfPathsFromOrigin(p Point) int {
	total := 0;
	g.ValidPaths = 0;

	startingVisitedNode := []Point{p}
	total += g.NumberOfPathsFromPoint(startingVisitedNode)

	return total;
}

func (g *Grid) NumberOfPathsFromPoint(visitedNodes []Point) int {

	fmt.Println("Visited: ", visitedNodes)

	if (len(visitedNodes) > 3) {
		g.ValidPaths++
		fmt.Println("Found one!")
	}

	currentPoint := visitedNodes[len(visitedNodes)-1];

	if(currentPoint.canMoveRight(g, visitedNodes)) {
		next := append(visitedNodes, Point{currentPoint.X, currentPoint.Y+1});
		g.NumberOfPathsFromPoint(next)
	}

	if(currentPoint.canMoveLeft(g, visitedNodes)) {
		next := append(visitedNodes, Point{currentPoint.X, currentPoint.Y-1});
		g.NumberOfPathsFromPoint(next)
	}

	if(currentPoint.canMoveUp(g, visitedNodes)) {
		next := append(visitedNodes, Point{currentPoint.X-1, currentPoint.Y});
		g.NumberOfPathsFromPoint(next)
	}

	if(currentPoint.canMoveDown(g, visitedNodes)) {
		next := append(visitedNodes, Point{currentPoint.X+1, currentPoint.Y});
		g.NumberOfPathsFromPoint(next)
	}

	return g.ValidPaths
}

type Point struct{
	X int
	Y int
}

func (p Point) canMoveRight(grid *Grid, visitedPoints []Point) bool {
	if !(len(grid.Grid[p.X]) - 1 > p.Y) {
		return false
	}

	rightNeighbor := p
	rightNeighbor.Y++;
	return !rightNeighbor.WasVisited(visitedPoints)
}

func (p Point) canMoveLeft(grid *Grid, visitedPoints []Point) bool {
	if !(0 < p.Y){
		return false;
	}

	leftNeighbor := p
	leftNeighbor.Y--;
	return !leftNeighbor.WasVisited(visitedPoints);
}

func (p Point) canMoveUp(grid *Grid, visitedPoints []Point) bool {
	if !(0 < p.X){
		return false
	}

	upNeighbor := p
	upNeighbor.X--;
	return !upNeighbor.WasVisited(visitedPoints);
}

func (p Point) canMoveDown(grid *Grid, visitedPoints []Point) bool {
	if !(len(grid.Grid) - 1 > p.X){
		return false
	}

	downNeighbor := p
	downNeighbor.X++;
	return !downNeighbor.WasVisited(visitedPoints);
}

func (p *Point) WasVisited(visitedPoints []Point) bool {
	for _, visitedPoint := range visitedPoints {
		if p.X == visitedPoint.X && p.Y == visitedPoint.Y{
			return true
		}
	}
	return false
}