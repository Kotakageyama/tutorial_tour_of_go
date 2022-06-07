package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rt rot13Reader) Read(b []byte) (int, error) {
	defaultAlphabet := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	ChangedAlphabet := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
	n, err := rt.r.Read(b)
	if err != nil {
		return n, err
	}
	// 一つ一つ文字を取り出して該当するものに置き換える
	for k, v := range b {
		for key, value := range defaultAlphabet {
			if v == value {
				b[k] = ChangedAlphabet[key]
			}
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
