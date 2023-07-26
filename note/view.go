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
	slug, isValid := isValidDateSlug(slug)
	if !isValid {
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

	if len(files) == 0 || (len(files) == 1 && files[0].Name() == "config.toml") {
		fmt.Println("You have no notes.")

		return nil
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

func isValidDateSlug(slug string) (string, bool) {
	if len(slug) == 0 {
		return "", true
	}

	slug = strings.ReplaceAll(slug, ".", "-")

	// Check for ISO8601
	t, err := time.Parse("2006-01-02", slug)
	if err == nil {
		return t.Format("2006-01-02"), true
	}

	t, err = time.Parse("2006-01", slug)
	if err == nil {
		return t.Format("2006-01"), true
	}

	t, err = time.Parse("2006", slug)
	if err == nil {
		return t.Format("2006-01"), true
	}

	// Check for dates like 01.02.2003
	t, err = time.Parse("02-01-2006", slug)
	if err == nil {
		return t.Format("2006-01-02"), true
	}

	t, err = time.Parse("01-2006", slug)
	if err == nil {
		return t.Format("2006-01"), true
	}

	return "", false
}
