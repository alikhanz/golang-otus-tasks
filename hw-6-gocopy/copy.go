package main

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"io"
	"os"
)

// ReadChunkSize size of each chunk
const ReadChunkSize = 1024 * 4 // 4KB

// GoCopy structure
type GoCopy struct {
}

// NewGoCopy constructor
func NewGoCopy() GoCopy {
	return GoCopy{}
}

// Copy copy function
func (g *GoCopy) Copy(src, dst string, limit, offset int64, progressBarEnabled bool) (err error) {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0755)
	defer srcFile.Close()

	if err != nil {
		return err
	}

	stat, err := srcFile.Stat()

	if err != nil {
		return err
	}

	if stat.IsDir() {
		return errors.New(`cannot copy directory`)
	}

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0755)
	defer dstFile.Close()

	if err != nil {
		return err
	}

	if offset > stat.Size() {
		return errors.New("offset cannot be larger than file size")
	}

	if limit < 0 {
		return errors.New("limit cannot be negative")
	}

	if limit == 0 {
		limit = stat.Size() - offset
	}

	if limit <= 0 {
		return errors.New("incorrect size")
	}

	_, err = srcFile.Seek(offset, io.SeekStart)

	reader := io.LimitReader(srcFile, limit)

	buf := make([]byte, ReadChunkSize)

	writer := io.Writer(dstFile)

	if progressBarEnabled {
		bar := pb.New64(limit)
		defer bar.Finish()

		bar.Start()
		writer = bar.NewProxyWriter(writer)
	}

	for {
		n, err := reader.Read(buf)
		buf := buf[:n]

		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			return err
		}

		_, err = writer.Write(buf)

		if err != nil {
			return err
		}
	}

	return nil
}
