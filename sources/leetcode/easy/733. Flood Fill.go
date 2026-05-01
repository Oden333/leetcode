package easy

/*
	m == image.length
	n == image[i].length
	1 <= m, n <= 50
	0 <= image[i][j], color < 216
	0 <= sr < m
	0 <= sc < n
*/
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	type pxl struct {
		x, y int
	}
	start := pxl{sr, sc}

	old := image[start.x][start.y]
	image[start.x][start.y] = color

	getCh := func(x, y int) (pxl, bool) {
		var res pxl
		if x >= 0 && x < len(image) &&
			(y >= 0 && y < len(image[0])) &&
			image[x][y] != color &&
			old == image[x][y] {

			image[x][y] = color
			res = pxl{x, y}
			return res, true
		}
		return res, false
	}

	queue := make([]pxl, 0, len(image))
	queue = append(queue, start)

	for len(queue) > 0 {
		l := len(queue)
		for range l {
			node := queue[0]
			queue = queue[1:]

			if ch1, ok := getCh(node.x+1, node.y); ok {
				queue = append(queue, ch1)
			}
			if ch2, ok := getCh(node.x, node.y+1); ok {
				queue = append(queue, ch2)
			}
			if ch3, ok := getCh(node.x-1, node.y); ok {
				queue = append(queue, ch3)
			}
			if ch4, ok := getCh(node.x, node.y-1); ok {
				queue = append(queue, ch4)
			}

		}
	}

	return image
}
