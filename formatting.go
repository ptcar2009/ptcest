package pkg

import (
	"fmt"
	"strings"
)

func (t *T) failHeader() {
	fmt.Fprintf(t.writer, "\033[31;1m FAIL\033[0m\n")
}
func (t *T) okHeader() {
	fmt.Fprintf(t.writer, "\033[32;1m OK\033[0m\n")
}

func (t *T) Infof(format string, values ...interface{}) {
	errStr := ""
	for i := 0; i < t.level+1; i++ {
		errStr += "    "
	}
	errStr += "\033[33;1mINFO\033[0m "
	fmt.Fprintf(t.writerLogs, "%v%v\n", errStr, fmt.Sprintf(format, values...))
}

func (t *T) Errorf(format string, values ...interface{}) {
	errStr := ""
	for i := 0; i < t.level+1; i++ {
		errStr += "    "
	}

	format = strings.TrimSuffix(format, "\n")
	format = strings.Join(strings.Split(format, "\n"), "\n"+errStr)
	errStr += "\033[31;1mFAIL\033[0m "
	fmt.Fprintf(t.writerLogs, fmt.Sprintf("%v%v\n", errStr, format), values...)
	t.failed = true
}
