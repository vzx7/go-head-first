package S

type S1 struct {
	F1 int
	F2 string
	F3 int
}

type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

func (s S2slice) Len() int {
	return len(s)
}

func (s S2slice) Less(i, j int) bool {
	return s[i].F3.F1 < s[j].F3.F1
}

func (s S2slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
