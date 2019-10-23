package cat

import (
	"testing"
)

func TestRun(t *testing.T) {
	Run(111)
}

func BenchmarkTransport1(b *testing.B) {
	Run(222)
}
