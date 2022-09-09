package main

import (
	"fmt"
	"os"
	"strconv"
)

func getPids(path string) ([]int32, error) {
	fmt.Println("Try to get pids for path ", path)
	var ret []int32

	d, err := os.Open(path)
	if err != nil {
		fmt.Println("Cannot open path:", path)
		return nil, err
	}
	defer d.Close()

	fnames, err := d.Readdirnames(-1)

	fmt.Println("Number of files: ", len(fnames))

	for _, fname := range fnames {
		pid, err := strconv.ParseInt(fname, 10, 32)
		if err != nil {
			continue
		}
		fmt.Println("pid:", pid)
		ret = append(ret, int32(pid))
	}

	return ret, nil
}

func main() {
	pids, err := getPids("/proc")
	fmt.Println("Pids count:", len(pids))
	if err != nil {
		fmt.Errorf("Error %v while ger pids", err)
	}
}
