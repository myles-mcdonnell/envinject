package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	tokenStart = "$${{"
	tokenEnd = "}}"
)

func Inject(root string) error {
	fmt.Printf("Walking %v to find envvar replacement tokens\r\n", root)
	return filepath.WalkDir(root, func(path string, dirEntry os.DirEntry, err error) error {
		if !dirEntry.IsDir() {
			contents, err := ioutil.ReadFile(path)
			if err!=nil {
				return err
			}

			envVars := make([]string, 0)
			contentsStr := string(contents)
			offset := -1
			for offset < len(contents) {
				offset++
				remaining := contentsStr[offset:]
				i := strings.Index(remaining, tokenStart)
				if i == -1 {
					break
				}

				if len(contentsStr)-i < 6 {
					break
				}

				y := strings.Index(remaining, tokenEnd)
				if y == -1 {
					continue
				}

				envVars =  append(envVars, remaining[i+4: y] )
				offset = offset+y
			}

			for _, envVar  := range envVars {
				fmt.Printf("Injecting %v into %v\r\n", envVar, path)
				contentsStr = strings.ReplaceAll(contentsStr, tokenStart+envVar+tokenEnd, os.Getenv(envVar))
			}

			err = ioutil.WriteFile(path, []byte(contentsStr), 0644)
			if err!=nil {
				fmt.Printf("error writing file : %v : %v", path,  err.Error())
			}
		}

		if err != nil {
			log.Fatalln(err)
		}
		return nil
	})
}

