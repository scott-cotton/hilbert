// Copyright 2018 Scott Cotton. All rights reserved.  Use of this source
// code is governed by a license that can be found in the License file.

package hilbert

import (
	"fmt"
	"math"
	"testing"

	snd "zikichombo.org/sound"
	"zikichombo.org/sound/freq"
)

type sgen struct {
	rps float64
	ph  float64
	sr  freq.T
}

func (s *sgen) Channels() int {
	return 1
}

func (s *sgen) Close() error {
	return nil
}

func newSinGen(fa, fs freq.T) snd.Source {
	return &sgen{sr: fs, ph: 0.0, rps: fs.RadsPer(fa)}
}

func (s *sgen) Receive(dst []float64) (int, error) {
	ph, rps := s.ph, s.rps
	for i := range dst {
		dst[i] = math.Sin(ph)
		ph += rps
	}
	s.ph = ph
	return len(dst), nil
}

func (s *sgen) SampleRate() freq.T {
	return s.sr
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

	for n := 0; n < 100 && err == nil; n++ {
		err = hlb.Process()
		org, rot := hlb.Org(), hlb.Rotated()
		fmt.Printf("block:\n")
		for i := 0; i < N; i++ {
			fmt.Printf("%d in %f out %f\n", i, org[i], rot[i])
		}
	}
}
