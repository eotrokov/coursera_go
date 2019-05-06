package main

import (
	"fmt"
	//"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out *os.File, path string, printFiles bool) error {
	var oldlvl = -1
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		lvl := strings.Count(path, string(os.PathSeparator))
		rep := strings.Repeat("│\t", lvl)
		var skip = lvl
		
		fmt.Print(rep)
		if oldlvl >= lvl {
			fmt.Print("└")

		} else if oldlvl < lvl {
			fmt.Print("├")
			oldlvl = lvl
			skip = 0
		}
		
		fmt.Println(getInfo(info, printFiles, skip))
		oldlvl = lvl
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func getInfo(fileInfo os.FileInfo, printFiles bool, lvl int) string {
	var fileSize = ""
	rep := strings.Repeat("───", 1)
	fileS := fileInfo.Size()
	if fileS == 0 {
		fileSize = "empty"
	} else {
		fileSize = fmt.Sprintf("%db", fileS)
	}
	if !fileInfo.IsDir() && printFiles {
		return fmt.Sprintf("%s %s (%s)", rep, fileInfo.Name(), fileSize)
	} else {
		return fmt.Sprintf("%s %s", rep, fileInfo.Name())
	}
}
