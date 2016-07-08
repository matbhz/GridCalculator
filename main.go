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
//
//func (p *Point) canMoveTopRight(grid *Grid, visitedPoints []Point) bool{
//	if !(len(grid.Grid[p.X]) - 1 > p.Y && 0 < p.X) { // If it's on top or right edges you cant move
//		return false
//	}
//
//	topRightNeighbor := p
//	topRightNeighbor.Y++;
//	topRightNeighbor.X--;
//	return !topRightNeighbor.WasVisited(visitedPoints);
//}

func (g *Grid) canMoveRight(visitedPoints []Point, currentPoint Point) bool {
	if !(len(g.Grid[currentPoint.X]) - 1 > currentPoint.Y) {
		return false
	}

	rightNeighbor := currentPoint
	rightNeighbor.Y++;
	return !rightNeighbor.WasVisited(visitedPoints)
}

func (g *Grid) canMoveLeft(visitedPoints []Point, currentPoint Point) bool {
	if !(0 < currentPoint.Y){
		return false;
	}

	leftNeighbor := currentPoint
	leftNeighbor.Y--;
	return !leftNeighbor.WasVisited(visitedPoints);
}

func (g *Grid) canMoveUp(visitedPoints []Point, currentPoint Point) bool {
	if !(0 < currentPoint.X){
		return false
	}

	upNeighbor := currentPoint
	upNeighbor.X--;
	return !upNeighbor.WasVisited(visitedPoints);
}

func (g *Grid) canMoveDown(visitedPoints []Point, currentPoint Point) bool {
	if !(len(g.Grid) - 1 > currentPoint.X){
		return false
	}

	downNeighbor := currentPoint
	downNeighbor.X++;
	return !downNeighbor.WasVisited(visitedPoints);
}

///////

type Point struct{
	X int
	Y int
}

func (p *Point) WasVisited(visitedPoints []Point) bool {
	for _, visitedPoint := range visitedPoints {
		if p.X == visitedPoint.X && p.Y == visitedPoint.Y{
			return true
		}
	}
	return false
}