package dir

import (
	"io"
	"io/fs"
	"os"
	"sort"
)

func IsEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}

	_, err = f.ReadDir(1)
	f.Close()
	// ReadDirの引数が0より大きい場合は、成功時にerrがio.EOFになる
	if err == io.EOF {
		return true, nil
	} else {
		return false, err
	}
}

func Read(dirname string) ([]fs.DirEntry, []fs.DirEntry, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, nil, err
	}

	entries, err := f.ReadDir(-1)
	f.Close()
	if err != nil {
		return nil, nil, err
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	var (
		dirs  []fs.DirEntry
		files []fs.DirEntry
	)

	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry)
		} else {
			files = append(files, entry)
		}
	}

	return dirs, files, nil
}
