package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}

		if err != nil {
			return nil, err
		}
	}
}

func buildGzipReader(filename string) (*gzip.Reader, func(), error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}

	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func main() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(counts)

	r, closer, err := buildGzipReader("./gras.txt.gz")
	if err != nil {
		os.Exit(1)
	}
	defer closer()
	gzippedCounts, err := countLetters(r)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(gzippedCounts)
}
