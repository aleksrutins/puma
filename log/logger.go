package log

import (
	"fmt"

	"github.com/logrusorgru/aurora/v3"
)

func Warn(msg string) {
	fmt.Print(aurora.Sprintf(aurora.Yellow("%s %s\n"), aurora.Bold("WARN"), msg))
}
