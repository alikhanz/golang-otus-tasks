package main

import (
	"errors"
	"flag"
	"github.com/cheggaaa/pb/v3"
	"github.com/gookit/color"
	"io"
	"os"
)

func main() {
	var from, to string
	var limit, offset int64

	flag.StringVar(&from, "from", "", "/path/to/file")
	flag.StringVar(&to, "to", "", "/path/to/file")
	flag.Int64Var(&limit, "limit", 0, "1024")
	flag.Int64Var(&offset, "offset", 0, "1024")

	flag.Parse()

	c := NewGoCopy()

	err := c.Copy(from, to, limit, offset)

	if err != nil {
		color.SetOutput(os.Stderr)
		color.Error.Println(err.Error())
	}
}

const ReadChunkSize = 1024 * 4 // 4KB

type GoCopy struct {
	barEnabled bool
}

func NewGoCopy() GoCopy {
	return GoCopy{barEnabled: true}
}

func (g *GoCopy) Copy(src, dst string, limit, offset int64) (err error) {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0755)
	defer func() {
		cerr := srcFile.Close()

		if err == nil {
			err = cerr
		}
	}()

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
	defer func() {
		cerr := dstFile.Close()

		if err == nil {
			err = cerr
		}
	}()

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
	bar := pb.New64(limit)

	if g.barEnabled {
		bar.Start()
	}

	buf := make([]byte, ReadChunkSize)
	writer := bar.NewProxyWriter(dstFile)

	for {
		n, err := reader.Read(buf)
		buf := buf[:n]

		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				err = nil
				break
			}
			return err
		}

		if err != nil && err != io.EOF {
			bar.Finish()
			return err
		}

		_, err = writer.Write(buf)

		if err != nil {
			bar.Finish()
			return err
		}
	}

	bar.Finish()

	return nil
}
