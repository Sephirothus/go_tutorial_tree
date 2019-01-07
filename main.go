package main

import (
	//"fmt"
	//"io/ioutil"
	"os"
	"strconv"
	//"path/filepath"
	"strings"
	"sort"
	//"reflect"
	//"bytes"
)

type ByName []os.FileInfo
func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByName) Less(i, j int) bool {
	return a[i].Name() < a[j].Name()
}

func main() {
	out := os.Stdout
	//fmt.Println(reflect.TypeOf(out))
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
	lvl := len(strings.Split(path, "/")) - 1
	//fmt.Println(lvl)
	file, _ := os.Open(path)
	dirs, _ := file.Readdir(0)
	sort.Sort(ByName(dirs))
	for key, dir := range dirs {
		isLast := key + 1 == len(dirs)
		if dir.IsDir() {
			printEl(out, dir, isLast, lvl)
			if dir.IsDir() {
				dirTree(out, path + string(os.PathSeparator) + dir.Name(), printFiles)
			}
		} else if printFiles {
			printEl(out, dir, isLast, lvl)
		}
	}
	return nil
}

func printEl(out *os.File, dir os.FileInfo, isLast bool, lvl int) {
	if !isLast && lvl > 0 {
		//for i 
		//fmt.Println(lvl)
		out.Write([]byte(shiftCursor(lvl - 1) + "│"))
	}
	prefix := "├───"
	if isLast {
		prefix = "└───"
	}
	output := shiftCursor(lvl) + prefix + dir.Name()
	if !dir.IsDir() {
		output += " (" + getSize(dir) + ")"
	}
	out.Write([]byte(output + "\n"))
}

func shiftCursor(lvl int) string {
	return strings.Repeat("\t", lvl)
}

func getSize(dir os.FileInfo) string {
	size := dir.Size()
	if size == 0 {
		return "empty"
	} else {
		return strconv.FormatInt(size, 10) + "b"
	}
}
