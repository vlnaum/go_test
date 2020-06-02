package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	input, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer input.Close()

	inputInfo, err := input.Stat()
	if err != nil {
		return err
	}

	switch {
	case inputInfo.Size() == 0 || !inputInfo.Mode().IsRegular():
		return ErrUnsupportedFile
	case inputInfo.Size() < offset:
		return ErrOffsetExceedsFileSize
	case limit == 0 || inputInfo.Size() < limit:
		limit = inputInfo.Size()
	}

	output, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer output.Close()

	if _, err := input.Seek(offset, io.SeekStart); err != nil {
		return err
	}

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(input)

	if _, err := io.CopyN(output, barReader, limit); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil
}
