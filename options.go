package readability

type Options struct {
	// MaxElemsToParse is the max number of nodes supported by this
	// parser. Default: 0 (no limit)
	MaxElemsToParse int
	// NTopCandidates is the number of top candidates to consider when
	// analysing how tight the competition is among candidates.
	NTopCandidates int
	// CharThresholds is the default number of chars an article must
	// have in order to return a result
	CharThresholds int
	// ClassesToPreserve are the classes that readability sets itself.
	ClassesToPreserve []string
	// KeepClasses specify whether the classes should be stripped or not.
	KeepClasses bool
	// TagsToScore is element tags to score by default.
	TagsToScore []string
	// Debug determines if the log should be printed or not. Default: false.
	Debug bool
	// DisableJSONLD determines if metadata in JSON+LD will be extracted
	// or not. Default: false.
	DisableJSONLD bool
	// MinCharactersOfParagraph when the number of paragraph is less than MinCharactersOfParagraph, it will not be counted.
	MinCharactersOfParagraph int
	// BasePointOfParagraph base point of paragraph.
	BasePointOfParagraph int
	// MaxDepthToScore the max depth of ancestors node to score.
	MaxDepthToScore int
	// CommasWeight weight of commas to score.
	CommasWeight float64
	//MaxScoreOfParagraph max point of anys paragraph.
	MaxPointOfParagraph float64
	// CharactersPerScore how many characters in on point.
	CharactersPerPoint float64
}

var defaultOptions = Options{
	MaxElemsToParse:          0,
	NTopCandidates:           5,
	CharThresholds:           500,
	ClassesToPreserve:        []string{"page"},
	KeepClasses:              false,
	TagsToScore:              []string{"section", "h2", "h3", "h4", "h5", "h6", "p", "td", "pre"},
	Debug:                    false,
	BasePointOfParagraph:     1,
	MinCharactersOfParagraph: 25,
	MaxDepthToScore:          5,
	CommasWeight:             1,
	MaxPointOfParagraph:      3,
	CharactersPerPoint:       100,
}

func NewDefaultOptions() Options {
	return defaultOptions
}
