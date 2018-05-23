package utils

import (
	"bytes"
	"io"
	"os"
)

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
// Source: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// Create a file in a given dest from some buffer
func Create(buffer bytes.Buffer, dst string) error {
	reader := bytes.NewReader(buffer.Bytes())
	imgFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer imgFile.Close()
	_, err = io.Copy(imgFile, reader)
	if err != nil {
		return err
	}
	return nil
}
