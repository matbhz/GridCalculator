package tests

import (
    "fmt"
    "testing"
    . "../structs"
)

func TestGrid(t *testing.T) {

    // Squared grid
    testGrid := NewGrid(2, 2);

    expected := 6;
    actual := testGrid.NumberOfPathsFromPoint(Point{0,0})
    assert(expected, actual, t)

    expected = 24
    actual = testGrid.NumberOfPathsFromAllPoints()
    assert(expected, actual, t)

//    actual = testGrid.NumberOfPathsFromPoint(Point{0,0}) // Custom panic for accessing invalid position

    // Squared grid
    testGrid = NewGrid(3, 3)

    expected = 1354;
    actual = testGrid.NumberOfPathsFromPoint(Point{0,0})
    assert(expected, actual, t)

    expected = 10096
    actual = testGrid.NumberOfPathsFromAllPoints()
    assert(expected, actual, t)

    // Rectangular grid
    testGrid = NewGrid(2, 4)

    expected = 434;
    actual = testGrid.NumberOfPathsFromPoint(Point{0,0})
    assert(expected, actual, t)

    // "Horizontal" grid
    testGrid = NewGrid(1, 4)
    expected = 1;
    actual = testGrid.NumberOfPathsFromPoint(Point{0,0})
    assert(expected, actual, t)

    expected = 1;
    actual = testGrid.NumberOfPathsFromPoint(Point{0,3})
    assert(expected, actual, t)

    expected = 2;
    actual = testGrid.NumberOfPathsFromAllPoints()
    assert(expected, actual, t)

    // "Vertical grid"
    testGrid = NewGrid(4, 1)

    expected = 1;
    actual = testGrid.NumberOfPathsFromPoint(Point{0,0})
    assert(expected, actual, t)

    expected = 1;
    actual = testGrid.NumberOfPathsFromPoint(Point{3,0})
    assert(expected, actual, t)

    expected = 0;
    actual = testGrid.NumberOfPathsFromPoint(Point{1,0})
    assert(expected, actual, t)

    expected = 0;
    actual = testGrid.NumberOfPathsFromPoint(Point{2,0})
    assert(expected, actual, t)

    // "Point" grid
    testGrid = NewGrid(1, 1)

    expected = 0;
    actual = testGrid.NumberOfPathsFromPoint(Point{0,0})
    assert(expected, actual, t)

    // Zero grid
//    testGrid = NewGrid(0, 0) // Panics for having no points
}

func assert(expected, actual int, t *testing.T) {
    if (expected != actual){
        t.Error(fmt.Sprint("Expected ", expected, ". But was ", actual))
    }
}