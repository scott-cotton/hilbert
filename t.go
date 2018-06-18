package hilbert

import (
	"github.com/irifrance/audio/fft"
	"github.com/irifrance/snd"
)

// Type T implements a Processor for the Hilbert
// transform.
type T struct {
	ft    *fft.T
	ftBuf []complex128
	inBuf []float64
	oBuf  []float64
	init  bool
	n     int
	ii    int
	src   snd.Source
	snk   snd.Sink
}

// New generates a new Hilbert transformer, placing
// a rotated version of src in dst.  New processes
// n elements at a time with a quality factor of q.
//
// Normally, n*q samples should be atleast 1 period of the slowest
// frequency for which the transform is (relatively) accurate.
// The larger the ratio q/n, also the fewer edge effects
// in the output.  However, in our experience so long as the
// lowest frequency constraint is respected, we have found q>=2
// to be pretty good with diminishing returns.  q=1 sucks.
//
// T is fairly accurate phase-wise, but suffers from some
// amplitude fluctuation, amplified at lower frequencies.
//
func New(dst snd.Sink, src snd.Source, n, q int) *T {
	if q < 1 {
		panic("invalid q")
	}
	m := n * q
	res := &T{}
	res.ft = fft.New(m, false)
	res.ftBuf = res.ft.Win(nil)
	res.inBuf = make([]float64, m)
	res.oBuf = make([]float64, n)
	res.n = n
	if q%2 == 1 {
		res.src = src
	} else {
		res.src = &pad{Source: src, n: n / 2}
	}
	res.snk = &discard{Sink: dst, n: m / 2}
	res.ii = (m - n) / 2
	return res
}

// N() implements snd.Processor, giving the
// number of input samples mapped to M().
func (t *T) N() int {
	return t.n
}

// M() implements snd.Processor.
func (t *T) M() int {
	return t.n
}

// Org returns the last processed original signal, a
// slice of size N().
func (t *T) Org() []float64 {
	return t.inBuf[t.ii : t.ii+t.n]
}

// Rotated gives the rotated signal, a slice of size M().
func (t *T) Rotated() []float64 {
	return t.oBuf
}

// Process processes one block of samples.
func (t *T) Process() error {
	copy(t.inBuf, t.inBuf[t.n:])
	newStart := len(t.inBuf) - t.n
	n, e := t.src.Samples(t.inBuf[newStart:])
	if e != nil {
		return e
	}
	for i := newStart + n; i < len(t.inBuf); i++ {
		t.inBuf[i] = 0.0
	}
	for i, v := range t.inBuf {
		t.ftBuf[i] = complex(v, 0.0)
	}
	t.ft.Do(t.ftBuf)
	rotate(t.ftBuf)
	t.ft.Inv().Do(t.ftBuf)
	ii := t.ii
	for i := 0; i < t.n; i++ {
		t.oBuf[i] = real(t.ftBuf[ii+i])
	}
	return t.snk.PutSamples(t.oBuf)
}
