package window

// Window enables iteration through a rolling window size of string tokens.
func Window(elements []string, n int) [][]string {
	if len(elements) == 0 || len(elements) < n || n <= 0 {
		return nil
	}

	var (
		accum = make([][]string, len(elements)+1-n)
		j     int
	)

	for i := 0; i < len(elements)+1-n; i++ {
		win := make([]string, n)
		win[0] = elements[i]
		for j = 0; j+1 < n; j++ {
			win[j+1] = elements[i+j+1]
		}
		accum[i] = win
	}

	return accum
}
