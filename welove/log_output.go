package welove

import (
	"io"
	"os"
)

type CustomOutput struct {
	out []io.Writer
}

func New(out ...io.Writer) CustomOutput {
	return CustomOutput{out}
}

func (c *CustomOutput) Add(out io.Writer) {
	c.out = append(c.out, out)
}

func (c *CustomOutput) Write(p []byte) (int, error) {
	var n int = 0
	var err error = nil
	for _, v := range c.out {
		n, err = v.Write(p)
	}
	return n, err
}

func DefaultLog(path string) CustomOutput {
	var file, _ = os.OpenFile(path, os.O_APPEND | os.O_CREATE, os.ModeAppend)
	return New(os.Stdout, file)
}