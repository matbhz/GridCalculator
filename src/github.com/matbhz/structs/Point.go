package GridCalculator

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

func (p *Point) WasVisited(visitedPoints []Point) bool {
    for _, visitedPoint := range visitedPoints {
        if p.X == visitedPoint.X && p.Y == visitedPoint.Y{
            return true
        }
    }
    return false
}