package math

func max(f1, f2 float64) float64 {
	if f1 < f2 {
		return f2
	}
	return f1
}

func maxf(f1, f2 float32) float32 {
	if f1 < f2 {
		return f2
	}
	return f1
}
