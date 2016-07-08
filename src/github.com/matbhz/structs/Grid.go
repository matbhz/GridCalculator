package GridCalculator

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

func (g *Grid) NumberOfPathsFromAllPoints() int {
    total := 0

    for i := 0; i < len(g.Grid); i++ {
        for j := 0; j < len(g.Grid); j++ {
            total += g.NumberOfPathsFromPoint(Point{i, j})
        }
    }

    return total
}

func (g *Grid) NumberOfPathsFromPoint(p Point) int {
    startingVisitedNodes := []Point{p}
    g.ValidPaths = 0
    return g.NumberOfPaths(startingVisitedNodes);
}

func (g *Grid) NumberOfPaths(visitedNodes []Point) int {
    if (len(visitedNodes) > 3) {
        g.ValidPaths++
    }

    currentPoint := visitedNodes[len(visitedNodes)-1];

    if(currentPoint.canMoveRight(g, visitedNodes)) {
        next := append(visitedNodes, Point{currentPoint.X, currentPoint.Y+1});
        g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveLeft(g, visitedNodes)) {
        next := append(visitedNodes, Point{currentPoint.X, currentPoint.Y-1});
        g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveUp(g, visitedNodes)) {
        next := append(visitedNodes, Point{currentPoint.X-1, currentPoint.Y});
        g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveDown(g, visitedNodes)) {
        next := append(visitedNodes, Point{currentPoint.X+1, currentPoint.Y});
        g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveUpRight(g, visitedNodes)){
        next := append(visitedNodes, Point{currentPoint.X-1, currentPoint.Y+1});
        g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveUpLeft(g, visitedNodes)){
        next := append(visitedNodes, Point{currentPoint.X-1, currentPoint.Y-1});
        g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveDownRight(g, visitedNodes)){
        next := append(visitedNodes, Point{currentPoint.X+1, currentPoint.Y+1});
        g.NumberOfPaths(next)
    }

    if(currentPoint.canMoveDownLeft(g, visitedNodes)){
        next := append(visitedNodes, Point{currentPoint.X+1, currentPoint.Y-1});
        g.NumberOfPaths(next)
    }

    return g.ValidPaths
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