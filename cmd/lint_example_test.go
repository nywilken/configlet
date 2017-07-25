package cmd

func ExampleLint() {
	lintTrack("../fixtures/numbers")
	// Output:
	// -> An exercise with slug 'bajillion' is referenced in config.json, but no implementation was found.
	// -> The implementation for 'three' is missing an example solution.
	// -> The implementation for 'two' is missing an accompanying test suite.
	// -> An implementation for 'zero' was found, but config.json specifies that it should be foregone (not implemented).
}
