package resonacsv

import (
	"bytes"
	"io"
)

// lineCounter ファイルの行数を数える。
// 参考URL：https://www.it-swarm-ja.tech/ja/go/golang：ファイル内の行数を効率的に決定するにはどうすればよいですか？/1046642158/
func lineCounter(filesize int64, r io.Reader) (int64, error) {
	//	buf := make([]byte, 32*1024)
	buf := make([]byte, filesize)
	var count int64 = 0
	lineStep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += int64(bytes.Count(buf[:c], lineStep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
