package main

import (
	"archive/tar"
	"log"
	"os"
)

func main() {
	Writer()
}

// Writer writes a list an array files to a tar file
func Writer() {
	file, err := os.Create("tarfile.tar")
	if err != nil {
		log.Fatal(err)
	}
	w := tar.NewWriter(file)
	files := []struct {
		name string
		body string
		mode int64
	}{
		{
			name: "file1",
			body: "this is file1 content",
			mode: 0777,
		},
		{
			name: "file2",
			body: "this is file2 content",
			mode: 0777,
		},
		{
			name: "file3",
			body: "this is file3 content",
			mode: 0777,
		},
	}
	for _, f := range files {
		h := &tar.Header{
			Name: f.name,
			Size: int64(len(f.body)),
			Mode: f.mode,
		}
		if err := w.WriteHeader(h); err != nil {
			log.Fatal(err)
		}
		if _, err := w.Write([]byte(f.body)); err != nil {
			log.Fatal(err)
		}
	}
}
