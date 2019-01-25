# Faithful Hilbert Transform

This repository implements the "hilbert transform" on a sound data stream in a faithful manner.
The Hilbert transform is a phase shift in the frequency domain of all frequencies by 90%.

The mechanism implemented here is a sliding window of Fourier transforms in which only the middle
part contributes to the output of the transform.  This buffering/windowing, at least according
to our tests, provides much more accurate and also more efficient Hilbert transforms than say 
use an approximation by convolution.

Enjoy.
