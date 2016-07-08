package main
import "fmt"

func main() {
	myGrid := NewGrid(5,5);

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
	total += g.NumberOfPathsFromPoint(startingVisitedNode, p)

	return total;
}

func (g *Grid) NumberOfPathsFromPoint(visitedNodes []Point, p Point) int {

	fmt.Println("Visited: ", visitedNodes)

	if (len(visitedNodes) > 3) {
		g.ValidPaths++
		fmt.Println("Found one!")
	}

	if(g.canMoveRight(visitedNodes, p)) {
		pNext := p;
		pNext.Y++
		next := append(visitedNodes, pNext);

		g.NumberOfPathsFromPoint(next, pNext)
	}

	if(g.canMoveLeft(visitedNodes, p)) {
		pNext := p;
		pNext.Y--
		next := append(visitedNodes, pNext);

		g.NumberOfPathsFromPoint(next, pNext)
	}

	if(g.canMoveUp(visitedNodes, p)) {
		pNext := p;
		pNext.X--
		next := append(visitedNodes, pNext);

		g.NumberOfPathsFromPoint(next, pNext)
	}

	if(g.canMoveDown(visitedNodes, p)) {
		pNext := p;
		pNext.X++
		next := append(visitedNodes, pNext);

		g.NumberOfPathsFromPoint(next, pNext)
	}

	return g.ValidPaths
}

func (g *Grid) canMoveRight(visitedPoints []Point, currentPoint Point) bool {
	neighborNode := currentPoint
	neighborNode.Y++;

	return len(g.Grid[currentPoint.X]) - 1 > currentPoint.Y && !neighborNode.IsVisited(visitedPoints);
}

func (g *Grid) canMoveLeft(visitedPoints []Point, currentPoint Point) bool {
	neighborNode := currentPoint
	neighborNode.Y--;

	return 0 < currentPoint.Y && !neighborNode.IsVisited(visitedPoints);
}

func (g *Grid) canMoveUp(visitedPoints []Point, currentPoint Point) bool {
	neighborNode := currentPoint
	neighborNode.X--;

	return 0 < currentPoint.X && !neighborNode.IsVisited(visitedPoints);
}

func (g *Grid) canMoveDown(visitedPoints []Point, currentPoint Point) bool {
	neighborNode := currentPoint
	neighborNode.X++;

	return len(g.Grid) - 1 > currentPoint.X && !neighborNode.IsVisited(visitedPoints);
}

///////

type Point struct{
	X int
	Y int
}

func (p *Point) IsVisited(visitedPoints []Point) bool {
	for _, visitedPoint := range visitedPoints {
		if p.X == visitedPoint.X && p.Y == visitedPoint.Y{
			return true
		}
	}
	return false
}