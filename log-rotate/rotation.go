// Package rotation ...for now, it works for current directory only
package rotation

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func removeLastNumber(s string) string {
	ss := strings.Split(s, ".")
	return strings.Join(ss[:len(ss)-1], ".")
}

// Rotate ...you fill in path and targte files, and Rotate rotates the files
// for you out of the box
func Rotate(enable bool, path string, targets ...string) {
	if len(targets) == 0 {
		log.Warn("Oops, you gotta fill in some target files before I can start handling....")
		return
	}

	originfile := ""
	// define a slice to store files name
	foundFiles := []string{}
	// define a map to store the file name as key, and the index of file as value
	ns := map[string][]int{}

	// parse the directory to get file names
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		foundFiles = append(foundFiles, path)
		return nil
	})

	for _, file := range foundFiles {
		for _, target := range targets {
			if file == target {
				// fmt.Println("!!!!!!!!!!!")
				ns[target] = append(ns[target], 0)
			}
		}

		// split the file name into slice
		fileslice := strings.Split(file, ".")

		// do the valid files only
		if len(fileslice) > 0 {
			// extract the index
			lastInString := fileslice[len(fileslice)-1]

			// turn the index from type string into int
			i, err := strconv.Atoi(lastInString)
			if err == nil {
				// get the file name
				r := removeLastNumber(file)

				for _, target := range targets {
					if r == target {
						// append the indices to the key as file name
						ns[target] = append(ns[target], i)
					}
				}
			}
		}
	}

	if len(ns) == 0 {
		log.Warn("there is no matches found...")
		return
	}

	if enable {
		defer log.Info("done!")
	}

	// sort the file index so that wouldn't overwrite the file
	for _, v := range ns {
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
	}

	if enable {
		log.Info("doing the rotation...")
	} else {
		log.Info("just dry-run: it will be happend as the below")
		fmt.Println()
	}

	// do the rotating
	for k, v := range ns {
		for _, n := range v {
			// fmt.Println(k, n)
			var old, new string
			if n == 0 {
				old = k
				new = k + "." + strconv.Itoa(n+1)
				originfile = k

			} else {
				old = k + "." + strconv.Itoa(n)
				new = k + "." + strconv.Itoa(n+1)
			}

			if enable {
				os.Rename(old, new)
			}
			defer log.Infof("Rotated file: %s -> %s\n", old, new)
		}
	}

	if originfile != "" {
		f, err := os.OpenFile(originfile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Warn(err)
		}
		defer f.Close()

		f.WriteString("hello")
		log.Infof("Created new file: %s\n", originfile)
	}
}
