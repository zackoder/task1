package moudel

func Up(s *string) string {
	r := []rune(*s)
	for i := 0; i < len(r); i++ {
		if r[i] <= 'z' && r[i] >= 'a' {
			r[i] -= 32
		}
	}
	*s=string(r)
	return string(r)
}