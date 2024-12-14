package pi

func LeibnicSum(result chan float64, quit chan bool, startIndex int, stepIndex int) {
	pi := 0.0
	for i := startIndex; ; i += stepIndex {
		select {
		case <-quit:
			result <- pi
			return
		default:
			sum := 4.0 / (2.0*float64(i) + 1.0)
			if i%2 == 0 {
				pi += sum
			} else {
				pi -= sum
			}
		}
	}
}
