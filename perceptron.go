package golangML

import "errors"

//single-layer perceptron
//return error, w []float64, bias float64
func Perceptron(x ...[]float64) (error, []float64, float64) {
	len_x := len(x)
	num_x := len(x[0])
	learning_rate := 0.05
	bias := 1.
	bias_w := 0.3
	var w []float64
	var flag bool
	t := 0
	for i := 1; i < len_x; i++ {
		w = append(w, 1/float64(len_x-1))
	}
	for {
		flag = false
		t++
		for i := 0; i < num_x; i++ {
			sum, y := 0., 0.
			for index, xi := range x {
				if index == len_x-1 {
					sum += bias * bias_w
					if sum <= 0 {
						y = 0.
					} else {
						y = 1.
					}
					if xi[i] != y {
						flag = true
						for idx, xj := range x {
							if idx == len_x-1 {
								continue
							}
							w[idx] = w[idx] + learning_rate*xj[i]*(xi[i]-y)
						}
						bias_w = bias_w + learning_rate*bias*(xi[i]-y)
					}
				} else {
					sum += xi[i] * w[index]
				}
			}
		}
		if !flag {
			return nil, w, bias * bias_w
		}
		if t > 5000 {
			return errors.New("Can't find w"), nil, 0
		}
	}
}
