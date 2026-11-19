package note

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/GLAD-DEV/dn/config"
	"github.com/GLAD-DEV/dn/constants"
)

func Search(needle string, caseSensitive bool) error {
	path, err := config.BasePath()
	if err != nil {
		return err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	pattern := regexp.QuoteMeta(needle)
	if !caseSensitive {
		pattern = "(?i)" + pattern
	}
	re := regexp.MustCompile(pattern)

	var out strings.Builder
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Warning: Encountered a directory: %s\n", file.Name())

			continue
		}

		if file.Name() == constants.ConfigName {
			continue
		}

		//nolint:gosec
		// Name comes from os.ReadDir(config.BasePath()).
		contentByte, err := os.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			fmt.Printf("Warning: Reading the content of %s failed: %s\n", file.Name(), err)
			fmt.Print("Skipped file\n\n")

			continue
		}

		searchContent(&out, file.Name(), string(contentByte), re)
	}

	if out.Len() == 0 {
		return errors.New("no match found")
	}

	fmt.Println(strings.TrimSpace(out.String()))

	return nil
}

func searchContent(out *strings.Builder, fileName string, content string, re *regexp.Regexp) {
	if !re.MatchString(content) {
		return
	}

	for i, line := range strings.Split(content, "\n") {
		if !re.MatchString(line) {
			continue
		}

		line = re.ReplaceAllString(line, "\033[0;31m${0}\033[0m")
		// error can be ignored since strings.Builder.Write never returns an error
		_, _ = fmt.Fprintf(out, "%s:%d: %s\n", fileName, i+1, line)
	}
}
