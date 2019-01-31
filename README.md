# Faithful Hilbert Transform

This repository implements the "hilbert transform" on a sound data stream in a faithful manner.
The Hilbert transform is a phase shift in the frequency domain of all frequencies by 90 degrees.

The mechanism implemented here is a sliding window of Fourier transforms in which only the middle
part contributes to the output of the transform.  This buffering/windowing, at least according
to our tests, provides much more accurate and also more efficient Hilbert transforms than say 
using an approximation by convolution.

Enjoy.

## Citing Faithful Hilbert Transform

Various citation methods using zenodo DOI:
[![DOI](https://zenodo.org/badge/167553884.svg)](https://zenodo.org/badge/latestdoi/167553884)

BibTeX:

```
@misc{scott_cotton_2019_2553680,
  author       = {Scott Cotton},
  title        = {wsc0/hilbert: Faithful},
  month        = jan,
  year         = 2019,
  doi          = {10.5281/zenodo.2553680},
  url          = {https://doi.org/10.5281/zenodo.2553680}
}
```



