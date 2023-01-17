package types

const (
	VerbosityNormal Verbose = iota
	VerbosityVerbose
	VerbosityVeryVerbose
)

type Verbose byte

func (v Verbose) IsVerbose() bool {
	return v > VerbosityNormal
}
