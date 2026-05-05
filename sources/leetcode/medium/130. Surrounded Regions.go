package medium

func solve(board [][]byte) {
	type pxl struct {
		x, y int
	}
	turns := [4][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

	islands := make(map[pxl]struct{})
	trav := func(root pxl) {
		queue := make([]pxl, 0)
		queue = append(queue, root)
		islands[root] = struct{}{}

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, t := range turns {
				x := node.x + t[0]
				y := node.y + t[1]

				if x >= 0 && x < (len(board)) &&
					y >= 0 && y < len(board[0]) {
					p := pxl{x, y}
					if _, ok := islands[p]; !ok && board[p.x][p.y] == 'O' {
						islands[p] = struct{}{}
						queue = append(queue, p)
					}
				}
			}
		}
	}

	// borders
	for i, y := range board[0] {
		if y == 'O' {
			p := pxl{0, i}
			if _, ok := islands[p]; !ok {
				trav(p)
			}
		}
	}
	for i, y := range board[len(board)-1] {
		if y == 'O' {
			p := pxl{len(board) - 1, i}
			if _, ok := islands[p]; !ok {
				trav(p)
			}
		}
	}

	for i, col := range board {
		if col[0] == 'O' {
			p := pxl{i, 0}
			if _, ok := islands[p]; !ok {
				trav(p)
			}
		}

		if col[len(col)-1] == 'O' {
			p := pxl{i, len(col) - 1}
			if _, ok := islands[p]; !ok {
				trav(p)
			}
		}
	}

	for x, col := range board {
		for y /* , _ */ := range col {
			if _, ok := islands[pxl{x, y}]; !ok && board[x][y] == 'O' {
				board[x][y] = 'X'
			}
		}
	}
}
