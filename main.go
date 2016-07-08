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

	if (len(visitedNodes) > 3) {
		g.ValidPaths++
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

	if(currentPoint.canMoveUpRight(g, visitedNodes)){
		next := append(visitedNodes, Point{currentPoint.X-1, currentPoint.Y+1});
		g.NumberOfPathsFromPoint(next)
	}

	if(currentPoint.canMoveUpLeft(g, visitedNodes)){
		next := append(visitedNodes, Point{currentPoint.X-1, currentPoint.Y-1});
		g.NumberOfPathsFromPoint(next)
	}

	if(currentPoint.canMoveDownRight(g, visitedNodes)){
		next := append(visitedNodes, Point{currentPoint.X+1, currentPoint.Y+1});
		g.NumberOfPathsFromPoint(next)
	}

	if(currentPoint.canMoveDownLeft(g, visitedNodes)){
		next := append(visitedNodes, Point{currentPoint.X+1, currentPoint.Y-1});
		g.NumberOfPathsFromPoint(next)
	}

	return g.ValidPaths
}

type Point struct{
	X int
	Y int
}

func (p Point) canMoveRight(grid *Grid, visitedPoints []Point) bool {
	if grid.isRightWall(p) {
		return false
	}

	rightNeighbor := p
	rightNeighbor.Y++;
	return !rightNeighbor.WasVisited(visitedPoints)
}

func (p Point) canMoveLeft(grid *Grid, visitedPoints []Point) bool {
	if grid.isLeftWall(p) {
		return false;
	}

	leftNeighbor := p
	leftNeighbor.Y--;
	return !leftNeighbor.WasVisited(visitedPoints);
}

func (p Point) canMoveUp(grid *Grid, visitedPoints []Point) bool {
	if grid.isTopWall(p) {
		return false
	}

	upNeighbor := p
	upNeighbor.X--;
	return !upNeighbor.WasVisited(visitedPoints);
}

func (p Point) canMoveDown(grid *Grid, visitedPoints []Point) bool {
	if grid.isBottomWall(p) {
		return false
	}

	downNeighbor := p
	downNeighbor.X++;
	return !downNeighbor.WasVisited(visitedPoints);
}

func (p Point) canMoveUpRight(grid *Grid, visitedPoints []Point) bool {
	if grid.isTopRightWallCorner(p) {
		return false
	}

	upRightVisitor := p
	upRightVisitor.X--
	upRightVisitor.Y++
	return !upRightVisitor.WasVisited(visitedPoints);
}
func (p Point) canMoveUpLeft(grid *Grid, visitedPoints []Point) bool {
	if grid.isTopLeftWallCorner(p) {
		return false
	}

	upLeftNeighbor := p
	upLeftNeighbor.X--
	upLeftNeighbor.Y--
	return !upLeftNeighbor.WasVisited(visitedPoints);
}
func (p Point) canMoveDownRight(grid *Grid, visitedPoints []Point) bool {
	if grid.isBottomRightWallCorner(p) {
		return false;
	}

	downRightNeighbor := p
	downRightNeighbor.X++
	downRightNeighbor.Y++
	return !downRightNeighbor.WasVisited(visitedPoints);
}
func (p Point) canMoveDownLeft(grid *Grid, visitedPoints []Point) bool {
	if grid.isBottomLeftWallCorner(p) {
		return false
	}

	downLeftNeighbor := p
	downLeftNeighbor.X++
	downLeftNeighbor.Y--
	return !downLeftNeighbor.WasVisited(visitedPoints);
}

func (g *Grid) isLeftWall(p Point) bool {
	return !(p.Y > 0)
}

func (g *Grid) isRightWall(p Point) bool {
	return !(len(g.Grid[p.X]) - 1 > p.Y)
}

func (g *Grid) isTopWall(p Point) bool {
	return !(p.X > 0)
}

func (g *Grid) isBottomWall(p Point) bool { // TODO Make p a pointer
	return !(len(g.Grid) - 1 > p.X)
}

func (g *Grid) isTopRightWallCorner(p Point) bool {
	return g.isRightWall(p) || g.isTopWall(p)
}

func (g *Grid) isTopLeftWallCorner(p Point) bool {
	return g.isLeftWall(p) || g.isTopWall(p)
}

func (g *Grid) isBottomRightWallCorner(p Point) bool {
	return g.isRightWall(p) || g.isBottomWall(p)
}

func (g *Grid) isBottomLeftWallCorner(p Point) bool {
	return g.isLeftWall(p) || g.isBottomWall(p)
}

func (p *Point) WasVisited(visitedPoints []Point) bool {
	for _, visitedPoint := range visitedPoints {
		if p.X == visitedPoint.X && p.Y == visitedPoint.Y{
			return true
		}
	}
	return false
}