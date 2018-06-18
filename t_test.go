package hilbert

import (
	"fmt"
	"math"
	"testing"

	"github.com/irifrance/snd"
	"github.com/irifrance/snd/freq"
)

type sgen struct {
	rps float64
	ph  float64
}

func newSinGen(fa, fs freq.T) snd.Source {
	return &sgen{ph: 0.0, rps: fs.RadsPer(fa)}
}

func (s *sgen) Samples(dst []float64) (int, error) {
	ph, rps := s.ph, s.rps
	for i := range dst {
		dst[i] = math.Sin(ph)
		ph += rps
	}
	s.ph = ph
	return len(dst), nil
}

func TestT(t *testing.T) {
	testTSin(20*freq.Hertz, t)
}

func testTSin(fa freq.T, t *testing.T) {
	src := newSinGen(fa, 44100*freq.Hertz)
	var err error
	N := 2048
	Q := 8
	hlb := New(snd.Discard, src, N, Q)
	for err == nil {
		err = hlb.Process()
		org, rot := hlb.Org(), hlb.Rotated()
		fmt.Printf("block:\n")
		for i := 0; i < N; i++ {
			fmt.Printf("%d in %f out %f\n", i, org[i], rot[i])
		}
	}
}
