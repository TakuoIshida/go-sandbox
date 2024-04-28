package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/go-multierror"
)

func checkFileOpen(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.New("fail to open file: " + path)
	}
	defer file.Close()
	return nil
}

func main() {
	paths := []string{"not-exist1.txt", "not-exist2.txt"}
	var result *multierror.Error
	for _, path := range paths {
		if err := checkFileOpen(path); err != nil {
			result = multierror.Append(result, err)
		}
	}
	// multierror.ErrorFormatFunc
	err := errors.New("added error")
	f := result.ErrorFormat([]error{err})
	fmt.Println(f)
	if err := result.ErrorOrNil(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// Output:
	// 2 errors occurred:
	//     * error! : open not-exist1.txt: no such file or directory
	//     * error! : open not-exist2.txt: no such file or directory
}
