package track

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

// Track is a collection of Exercism exercises for a programming language.
type Track struct {
	path      string
	Config    Config
	Exercises []Exercise
}

// New loads a track.
func New(path string) (Track, error) {
	track := Track{
		path: filepath.FromSlash(path),
	}

	c, err := NewConfig(filepath.Join(path, "config.json"))
	if err != nil {
		return track, err
	}
	track.Config = c

	dir := filepath.Join(track.path, "exercises")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return track, err
	}

	srgx, _ := regexp.Compile(track.Config.SolutionPattern)
	trgx, _ := regexp.Compile(track.Config.TestPattern)
	for _, file := range files {
		if file.IsDir() {
			fp := filepath.Join(dir, file.Name())
			ex, err := NewExercise(fp, srgx)
			if err != nil {
				return track, err
			}

			ex.LoadTestSuitePath(fp, trgx)
			track.Exercises = append(track.Exercises, ex)
		}
	}
	return track, nil
}
