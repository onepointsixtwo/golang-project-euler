package eleven

import (
	"testing"
)

func TestGetGrid(t *testing.T) {
	grid := getGrid()
	if grid[0][1] != 2 || grid[4][2] != 16 {
		t.Error("Grid has not parsed correctly")
	}
}
