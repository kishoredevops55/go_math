package geometry

import "errors"

func CubeVoulme(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("0 length edge is not allowed")
	}
}
