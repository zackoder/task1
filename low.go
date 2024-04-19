package moudel

func Low(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] <= 'Z' && r[i] >= 'A' {
			r[i] += 32
		}
	}
	return string(r)
}