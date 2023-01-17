package types

const (
	VerbosityNormal      Verbose = 0
	VerbosityVerbose     Verbose = 1
	VerbosityVeryVerbose Verbose = 2
)

type Verbose byte

func (v Verbose) IsVerbose() bool {
	return v > VerbosityNormal
}
