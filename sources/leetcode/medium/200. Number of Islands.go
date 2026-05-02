package medium

func numIslands(grid [][]byte) int {
	type pxl struct {
		x, y int
	}
	// X := func(p pxl) int {
	// 	return int(p.x)
	// }
	// Y := func(p pxl) int {
	// 	return int(p.y)
	// }

	turns := [4][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	// seen := make(map[pxl]struct{})

	var trav func(p pxl)
	trav = func(p pxl) {
		// seen[p] = struct{}{}
		grid[p.x][p.y] = '0'
		for _, t := range turns {
			if p.x+t[0] >= 0 && p.x+t[0] < len(grid) &&
				p.y+t[1] >= 0 && p.y+t[1] < len(grid[0]) &&
				grid[p.x+t[0]][p.y+t[1]] == '1' {
				new := pxl{p.x + t[0], p.y + t[1]}
				// if _, ok := seen[new]; !ok {
				trav(new)
				// }
			}
		}
	}

	counter := 0

	for x, col := range grid {
		for y := range col {
			cur := pxl{x, y}
			if /* _, ok := seen[cur]; !ok && */ grid[x][y] == '1' {
				counter++
				trav(cur)
			}
		}
	}

	return counter
}
