func surfaceArea(grid [][]int) int {
	N := len(grid)
	ans := 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			c := grid[i][j]
			if c == 0 {
				continue
			}
			ans += 2
			if i == 0 {
				ans += c
			} else if grid[i-1][j] < c {
				ans += c - grid[i-1][j]
			}
			if j == 0 {
				ans += c
			} else if grid[i][j-1] < c {
				ans += c - grid[i][j-1]
			}
			if i == N-1 {
				ans += c
			} else if grid[i+1][j] < c {
				ans += c - grid[i+1][j]
			}
			if j == N-1 {
				ans += c
			} else if grid[i][j+1] < c {
				ans += c - grid[i][j+1]
			}
		}
	}
	return ans
}