package hilbert

import (
	"math"

	"github.com/irifrance/snd"
)

type pad struct {
	snd.Source
	n int
}

func (p *pad) Samples(dst []float64) (int, error) {
	i := 0
	for i < p.n && i < len(dst) {
		dst[i] = 0.0
		i++
	}
	p.n -= i
	n, err := p.Source.Samples(dst[i:])
	if err != nil {
		return 0, err
	}
	return n + i, nil
}

type discard struct {
	snd.Sink
	n int
}

func (d *discard) PutSamples(vs []float64) error {
	if d.n >= len(vs) {
		d.n -= len(vs)
		return nil
	}
	n := d.n
	d.n = 0
	return d.Sink.PutSamples(vs[n:])
}

func princVal(v float64) float64 {
	return math.Mod(v+math.Pi, 2*math.Pi) - math.Pi
}

func rotate(d []complex128) {
	m := len(d) / 2
	for i := 1; i < m; i++ {
		d[i] *= -1i
	}
	for i := m; i < len(d); i++ {
		d[i] *= 1i
	}
}

func unWrap(phs []float64) {
	unWrapFrom(0.0, phs)
}

func unWrapFrom(last float64, phs []float64) {
	var dp, acc float64
	acc = 0.0
	for i, ph := range phs {
		dp = ph - last
		if dp < -math.Pi {
			acc += 2 * math.Pi
		} else if dp > math.Pi {
			acc -= 2 * math.Pi
		}
		last = ph
		phs[i] = ph + acc
	}
}

func wrap(phs []float64) {
	for i, ph := range phs {
		phs[i] = princVal(ph)
	}
}

func diff(vs []float64) {
	diffFrom(0.0, vs)
}

func diffFrom(last float64, vs []float64) {
	for i, v := range vs {
		vs[i] = v - last
		last = v
	}
}

func unDiff(vs []float64) {
	unDiffFrom(0.0, vs)
}

func unDiffFrom(last float64, vs []float64) {
	var t float64
	for i, v := range vs {
		t = v + last
		vs[i] = t
		last = t
	}
}
