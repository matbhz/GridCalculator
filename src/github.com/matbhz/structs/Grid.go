package GridCalculator

import (
    "fmt"
    "sync"
)

type Grid struct {
    Grid [][]*Point
}

func NewGrid(rows, columns int) *Grid {

    if(rows <= 0 || columns <= 0) {
        panic("A grid must have at least 1 row and 1 column")
    }

    grid := [][]*Point{}

    for i := 0; i < rows; i++ {
        row := make([]*Point, columns);
        grid = append(grid, row);
    }

    return &Grid{Grid: grid}
}

func (g* Grid) NumberOfPathsFromAllPointsParallel() int {

    goRoutines := g.Rows() * g.Columns(0)
    subtotalsChan := make(chan int, goRoutines)

    var wg sync.WaitGroup
    wg.Add(goRoutines) // TODO: Currently only working for rectangular Grids

    for i := 0; i < g.Rows(); i++ {
        for j := 0; j < g.Columns(i); j++ {
            go func(x, y int){
                defer wg.Done()
                total := g.NumberOfPathsFromPoint(Point{x, y})
                subtotalsChan <- total
            }(i, j)
        }
    }

    wg.Wait()
    close(subtotalsChan)

    total := 0
    for subtotal := range subtotalsChan{
        total += subtotal
    }

    return total

}

func (g *Grid) NumberOfPathsFromAllPoints() int {
    total := 0

    for i := 0; i < g.Rows(); i++ {
        for j := 0; j < g.Columns(i); j++ {
            total += g.NumberOfPathsFromPoint(Point{i, j})
        }
    }

    return total
}

func (g *Grid) NumberOfPathsFromPoint(p Point) int {

    if(p.X > g.Rows() || p.Y > g.Columns(p.X) || p.X < 0 || p.Y < 0){
        panic("Invalid point")
    }

    startingVisitedNodes := []Point{p}
    return g.NumberOfPaths(startingVisitedNodes);
}

func (g *Grid) NumberOfPaths(visitedNodes []Point) int {

    validPaths := 0

    if (len(visitedNodes) > 3) {
        validPaths++
    }

    currentPoint := visitedNodes[len(visitedNodes)-1];

    if(currentPoint.canMoveRight(g, visitedNodes)) {
        next := append(visitedNodes, Point{currentPoint.X, currentPoint.Y+1});
        validPaths += g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveLeft(g, visitedNodes)) {
        next := append(visitedNodes, Point{currentPoint.X, currentPoint.Y-1});
        validPaths += g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveUp(g, visitedNodes)) {
        next := append(visitedNodes, Point{currentPoint.X-1, currentPoint.Y});
        validPaths += g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveDown(g, visitedNodes)) {
        next := append(visitedNodes, Point{currentPoint.X+1, currentPoint.Y});
        validPaths += g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveUpRight(g, visitedNodes)){
        next := append(visitedNodes, Point{currentPoint.X-1, currentPoint.Y+1});
        validPaths += g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveUpLeft(g, visitedNodes)){
        next := append(visitedNodes, Point{currentPoint.X-1, currentPoint.Y-1});
        validPaths += g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveDownRight(g, visitedNodes)){
        next := append(visitedNodes, Point{currentPoint.X+1, currentPoint.Y+1});
        validPaths += g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveDownLeft(g, visitedNodes)){
        next := append(visitedNodes, Point{currentPoint.X+1, currentPoint.Y-1});
        validPaths += g.NumberOfPaths(next)
    }

    return validPaths
}

func (g *Grid) isLeftWall(p *Point) bool {
    return !(p.Y > 0)
}

func (g *Grid) isRightWall(p *Point) bool {
    return !(len(g.Grid[p.X]) - 1 > p.Y)
}

func (g *Grid) isTopWall(p *Point) bool {
    return !(p.X > 0)
}

func (g *Grid) isBottomWall(p *Point) bool { // TODO Make p a pointer
    return !(len(g.Grid) - 1 > p.X)
}

func (g *Grid) isTopRightWallCorner(p *Point) bool {
    return g.isRightWall(p) || g.isTopWall(p)
}

func (g *Grid) isTopLeftWallCorner(p *Point) bool {
    return g.isLeftWall(p) || g.isTopWall(p)
}

func (g *Grid) isBottomRightWallCorner(p *Point) bool {
    return g.isRightWall(p) || g.isBottomWall(p)
}

func (g *Grid) isBottomLeftWallCorner(p *Point) bool {
    return g.isLeftWall(p) || g.isBottomWall(p)
}

func (g *Grid) Rows() int {
    return len(g.Grid)
}

func (g *Grid) Columns(rowNumber int) int {
    if rowNumber > len(g.Grid){
        panic(fmt.Sprintf("Grid only has", rowNumber, len(g.Grid), "rows"))
    }
    return len(g.Grid[rowNumber])
}