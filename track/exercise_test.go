package track

import (
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExerciseSlug(t *testing.T) {
	path := filepath.FromSlash("../fixtures/fake-exercise")

	rgx, err := regexp.Compile("")
	assert.NoError(t, err)

	ex, err := NewExercise(path, rgx)
	assert.NoError(t, err)
	assert.Equal(t, "fake-exercise", ex.Slug)
}

func TestExerciseSolutionPaths(t *testing.T) {
	tests := []struct {
		solution string
		pattern  string
	}{
		{
			// It finds files in the root of the exercise directory.
			pattern:  "[Ee]xample",
			solution: "example.ext",
		},
		{
			// It finds files in a subdirectory.
			pattern:  "solution",
			solution: "subdir/solution.ext",
		},
		{
			// It only matches files, not directories.
			pattern:  "subdir",
			solution: "subdir/solution.ext",
		},
		// It finds hidden files.
		{
			pattern:  "secret",
			solution: "subdir/.secret-solution.ext",
		},
		// it finds files in hidden directories
		{
			pattern:  "hidden.file\\.ext",
			solution: ".hidden/file.ext",
		},
	}

	for _, test := range tests {
		path := filepath.FromSlash("../fixtures/fake-exercise")

		rgx, err := regexp.Compile(test.pattern)
		assert.NoError(t, err)

		ex, err := NewExercise(path, rgx)
		assert.NoError(t, err)

		assert.Equal(t, test.solution, ex.SolutionPath)
	}
}
func TestExerciseLoadTestSuitePath(t *testing.T) {
	tests := []struct {
		pattern   string
		testsuite string
	}{
		{
			// Default pattern finds files in a subdirectory named tests.
			pattern:   "(?i)test",
			testsuite: "tests/file.ext",
		},
		{
			// It finds files with an exact match.
			pattern:   "[^_]+_test.ext$",
			testsuite: "fake_test.ext",
		},
		{
			// It finds files in a subdirectory.
			pattern:   "secret",
			testsuite: "subdir/.secret-solution.ext",
		},
	}

	for _, test := range tests {
		path := filepath.FromSlash("../fixtures/fake-exercise")

		rgx, err := regexp.Compile("[Ee]xample")
		assert.NoError(t, err)

		ex, err := NewExercise(path, rgx)
		assert.NoError(t, err)

		trgx, err := regexp.Compile(test.pattern)
		assert.NoError(t, err)

		err = ex.LoadTestSuitePath(path, trgx)
		assert.NoError(t, err)

		assert.Equal(t, test.testsuite, ex.TestSuitePath)
	}
}
