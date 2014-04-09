package main

func main() {
	xs := []float64{1, 2, 3, 4}
	for i, x := range xs {
		xs[i] = x*x            // HL
	}
}
