package note

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/GLAD-DEV/dn/config"
)

func View(slug string) error {
	slug = strings.ReplaceAll(slug, ".", "-")
	if !isValidDateSlug(slug) {
		return fmt.Errorf("the passed date slug '%s' is invalid", slug)
	}

	path, err := config.BasePath()
	if err != nil {
		return err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	out := ""
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Warning: Encountered a directory: %s\n", file.Name())

			continue
		}

		// All files have a prefix of ""
		if !strings.HasPrefix(file.Name(), slug) || file.Name() == "config.toml" {
			continue
		}

		content, err := os.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			fmt.Printf("Warning: Reading the content of %s failed: %s\n", file.Name(), err)
			fmt.Print("Skipped file\n\n")
		}

		// Add space before every bullet point
		split := strings.Split(string(content), "\n")
		for i, str := range split {
			split[i] = " " + str
		}

		out += strings.TrimSpace(fmt.Sprintf("%s\n%s", file.Name(), strings.Join(split, "\n"))) + "\n"
	}

	out = strings.TrimSpace(out)
	if len(out) == 0 {
		return fmt.Errorf("no matching note found for date %s", slug)
	}

	fmt.Print(out)

	return nil
}

func isValidDateSlug(dateStr string) bool {
	if len(dateStr) == 0 {
		return true
	}

	// Check for ISO8601
	_, err := time.Parse("2006-01-02", dateStr)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01", dateStr)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006", dateStr)
	if err == nil {
		return true
	}

	// Check for dates like 01.02.2003
	_, err = time.Parse("02-01-2006", dateStr)
	if err == nil {
		return true
	}

	_, err = time.Parse("01-2006", dateStr)

	return err == nil
}