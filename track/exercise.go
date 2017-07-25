package track

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Exercise is an implementation of an Exercism exercise.
type Exercise struct {
	Slug          string
	SolutionPath  string
	TestSuitePath string
}

// NewExercise loads an exercise.
func NewExercise(root string, rgx *regexp.Regexp) (Exercise, error) {
	ex := Exercise{
		Slug: filepath.Base(root),
	}

	err := filepath.Walk(root, fieldPath(root, rgx, &ex.SolutionPath))
	return ex, err
}

// LoadTestSuitePath loads the test suite path for exercise using the provided file pattern.
func (ex *Exercise) LoadTestSuitePath(root string, rgx *regexp.Regexp) error {
	err := filepath.Walk(root, fieldPath(root, rgx, &ex.TestSuitePath))
	return err
}

// fieldPath returns a WalkFunc type that can be used to update the value of field
// with the string matched by rgx for the given root path.
func fieldPath(root string, rgx *regexp.Regexp, field *string) filepath.WalkFunc {

	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if rgx.Match([]byte(path)) {
			prefix := fmt.Sprintf("%s%s", root, string(filepath.Separator))
			*field = strings.Replace(path, prefix, "", 1)
		}
		return nil
	}
}

// HasTestSuite checks that an exercise has an accompanying test suite.
func (ex Exercise) HasTestSuite() bool {
	return ex.TestSuitePath != ""
}

// IsValid checks that an exercise has a sample solution.
func (ex Exercise) IsValid() bool {
	return ex.SolutionPath != ""
}
