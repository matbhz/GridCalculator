package GridCalculator

type Point struct{
    X int
    Y int
}

func (p *Point) canMoveRight(grid *Grid, visitedPoints []Point) bool {
    if grid.isRightWall(p) {
        return false
    }

    rightNeighbor := &Point{X:p.X, Y:p.Y}
    rightNeighbor.Y++;
    return !rightNeighbor.wasVisited(visitedPoints)
}

func (p *Point) canMoveLeft(grid *Grid, visitedPoints []Point) bool {
    if grid.isLeftWall(p) {
        return false;
    }

    leftNeighbor := &Point{X:p.X, Y:p.Y}
    leftNeighbor.Y--;
    return !leftNeighbor.wasVisited(visitedPoints);
}

func (p *Point) canMoveUp(grid *Grid, visitedPoints []Point) bool {
    if grid.isTopWall(p) {
        return false
    }

    upNeighbor := &Point{X:p.X, Y:p.Y}
    upNeighbor.X--;
    return !upNeighbor.wasVisited(visitedPoints);
}

func (p *Point) canMoveDown(grid *Grid, visitedPoints []Point) bool {
    if grid.isBottomWall(p) {
        return false
    }

    downNeighbor := &Point{X:p.X, Y:p.Y}
    downNeighbor.X++;
    return !downNeighbor.wasVisited(visitedPoints);
}

func (p *Point) canMoveUpRight(grid *Grid, visitedPoints []Point) bool {
    if grid.isTopRightWallCorner(p) {
        return false
    }

    upRightVisitor := &Point{X:p.X, Y:p.Y}
    upRightVisitor.X--
    upRightVisitor.Y++
    return !upRightVisitor.wasVisited(visitedPoints);
}

func (p *Point) canMoveUpLeft(grid *Grid, visitedPoints []Point) bool {
    if grid.isTopLeftWallCorner(p) {
        return false
    }

    upLeftNeighbor := &Point{X:p.X, Y:p.Y}
    upLeftNeighbor.X--
    upLeftNeighbor.Y--
    return !upLeftNeighbor.wasVisited(visitedPoints);
}

func (p *Point) canMoveDownRight(grid *Grid, visitedPoints []Point) bool {
    if grid.isBottomRightWallCorner(p) {
        return false;
    }

    downRightNeighbor := &Point{X:p.X, Y:p.Y}
    downRightNeighbor.X++
    downRightNeighbor.Y++
    return !downRightNeighbor.wasVisited(visitedPoints);
}

func (p *Point) canMoveDownLeft(grid *Grid, visitedPoints []Point) bool {
    if grid.isBottomLeftWallCorner(p) {
        return false
    }

    downLeftNeighbor := &Point{X:p.X, Y:p.Y}
    downLeftNeighbor.X++
    downLeftNeighbor.Y--
    return !downLeftNeighbor.wasVisited(visitedPoints);
}

func (p *Point) wasVisited(visitedPoints []Point) bool {
    for _, visitedPoint := range visitedPoints {
        if p.X == visitedPoint.X && p.Y == visitedPoint.Y{
            return true
        }
    }
    return false
}