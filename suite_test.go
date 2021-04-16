package pkg

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_Suite(t *testing.T) {
	rw := bytes.NewBufferString("")
	rwb := bytes.NewBufferString("")
	tt := &T{
		name:       "",
		level:      -1,
		writer:     rw,
		writerBody: rwb,
		writerLogs: rwb,
	}
	baseSuite.run(tt)
	if tt.failed {
		fmt.Print(rwb.String())
		t.Fail()
	}
}
