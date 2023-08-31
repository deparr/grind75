

// BAD hashmap way
// I really need to go to bed
type point struct {
    r int
    c int
}
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    origColor := image[sr][sc]
    m, n := len(image), len(image[0])
    vis := map[point]struct{}{}

    var fill func(int, int)
    fill = func (r, c int) {
        currpoint := point{r,c}

        if _, prs := vis[currpoint]; prs {
            return
        }

        vis[currpoint] = struct{}{}
        
        if image[r][c] == origColor {
            image[r][c] = color
        } else {
            return
        }


        prevCGood, nextCGood := c-1 >= 0, c+1 < n
        prevRGood, nextRGood := r-1 >= 0, r+1 < m

        if prevRGood {
            fill(r-1, c)
        }
        if nextRGood {
            fill(r+1, c)
        }
        if prevCGood {
            fill(r, c-1)
        }
        if nextCGood {
            fill(r, c+1)
        }
    }

    fill(sr, sc)
    return image
}

