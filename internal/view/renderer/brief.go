package renderer

type Brief struct {
	No          int
	Timestamp   int64
	Length      int
	Source      string
	Destination string
	Protocol    string
	Info        string
	Phony       bool
}
