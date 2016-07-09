package tests

import (
    "fmt"
    "testing"
    . "../structs"
)

func TestGrid(t *testing.T) {

    testGrid := NewGrid(2, 2);

    expected := 24
    actual := testGrid.NumberOfPathsFromAllPoints()
    assert(expected, actual, t)

    expected = 6;
    actual = testGrid.NumberOfPathsFromPoint(Point{0,0})
    assert(expected, actual, t)

    testGrid = NewGrid(3, 3)
    expected = 1354;
    actual = testGrid.NumberOfPathsFromPoint(Point{0,0})
    assert(expected, actual, t)

    testGrid = NewGrid(1, 4)
    expected = 2;
    actual = testGrid.NumberOfPathsFromPoint(Point{0,0})
    assert(expected, actual, t)
    
}

func assert(expected, actual int, t *testing.T) {
    if (expected != actual){
        t.Error(fmt.Sprint("Expected ", expected, ". But was ", actual))
    }
}