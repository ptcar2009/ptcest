package pkg

import (
	"io"
)

type T struct {
	name string

	level      int
	failed     bool
	writer     io.Writer
	writerBody io.Writer
	writerLogs io.Writer
}
