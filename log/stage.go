package log

import "fmt"

type StageSpinner struct {
	Finished int
	Total    int
	Message  string
}

func (s StageSpinner) Print() {
	fmt.Printf("\x1b[1A\x1b[2K[%d/%d] %s\n", s.Finished, s.Total, s.Message)
}
