package moudel

func Cap(s string) string {
	r := []rune(s)
	if r[0] <= 'z' && r[0] >= 'a' {
		r[0] -= 32
	}
	return string(r)
}