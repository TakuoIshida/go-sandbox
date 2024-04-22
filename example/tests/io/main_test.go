package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Double(stdin io.Reader, stdout io.Writer) error {
	buf, err := io.ReadAll(stdin)
	if err != nil {
		return err
	}

	stdout.Write(buf)
	stdout.Write(buf)

	return nil
}

func TestDouble(t *testing.T) {
	stdin := bytes.NewBufferString("foo\n")
	stdout := new(bytes.Buffer)

	err := Double(stdin, stdout)
	assert.NoError(t, err)

	expected := []byte("foo\nfoo\n")
	assert.Equal(t, expected, stdout.Bytes())
}
