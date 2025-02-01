# Mandelbrot written in go

## Tools

- Ebiten
    - 2D Rendering
- image
- math/cmplx
- time

## Math needed for Mandelbrot

### Complex numbers

- Complex numbers have a real and a imaginary part.
- $c = a + bi$

### Iteration Formula:

- For each point $c$ you start with $z = 0$ and repeatedly apply the formular

$$
z = z^2 + c
$$

#### Iteration example

1. $z_0 = 0$
2. $z_1 = 0^2 + c = c$
3. $z_2 = z_1^2 + c$
# mandelbrot-go
