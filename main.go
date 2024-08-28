package pi

import (
	"errors"
)

func pi_leibniz(n float64) (float64, error) {
	if n == 0 {
		return 0, errors.New("n cant be 0")
	}

	var pi float64 = 0
	var sign float64 = 1
	var i float64 = 1

	for i = 1; i <= (n * 2); i += 2 {
		pi = pi + sign*(4 / i)
		sign = -sign
	}

	return pi, nil
}

func pi_wallis(n float64) (float64, error) {
	if n == 0 {
		return 0, errors.New("n cant be 0")
	}

	var pi float64 = 4.
	var i float64 = 3

	for i = 3; i <= (n + 2); i += 2 {
		pi = pi * ((i - 1) / i) * ((i + 1) / i)
	}

	return pi, nil
}

func pi_nilakantha(n float64) (float64, error) {
	if n == 0 {
		return 0, errors.New("n cant be 0")
	}

	var pi float64 = 3
	var sign float64 = 1
	var i float64 = 2

	for i = 2; i <= (n * 2); i += 2 {
		pi = pi + sign*(4 / (i * (i + 1) * (i + 2)))
		sign = -sign
	}

	return pi, nil
}
