//go:build !darwin || !arm64 || !cgo

package m1cpu

// IsAppleSilicon return false on this platform.
func IsAppleSilicon() bool {
	return false
}

// PCoreHZ returns the max frequency in Hertz of the P-Core of an Apple Silicon CPU.
func PCoreHz() uint64 {
	panic("m1cpu: not a darwin/arm64 system")
}

// ECoreHZ returns the max frequency in Hertz of the E-Core of an Apple Silicon CPU.
func ECoreHz() uint64 {
	panic("m1cpu: not a darwin/arm64 system")
}

// PCoreGHz returns the max frequency in Gigahertz of the P-Core of an Apple Silicon CPU.
func PCoreGHz() float64 {
	panic("m1cpu: not a darwin/arm64 system")
}

// ECoreGHz returns the max frequency in Gigahertz of the E-Core of an Apple Silicon CPU.
func ECoreGHz() float64 {
	panic("m1cpu: not a darwin/arm64 system")
}
