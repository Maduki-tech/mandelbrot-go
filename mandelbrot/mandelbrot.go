package mandelbrot

import "math/cmplx"

func Mandelbrot(c complex128, maxIter int) int {
	var z complex128 = 0
	for n := 0; n < maxIter; n++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return n
		}
	}
	return maxIter
}
