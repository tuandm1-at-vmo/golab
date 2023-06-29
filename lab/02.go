package lab

import (
	"encoding/json"
	"os"
)

func GetAllPrimesUpTo(n uint, source string) ([]uint, error) {
	var primes map[uint]uint
	jsonBytes, err := os.ReadFile(source)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytes, &primes)
	if err != nil {
		return nil, err
	}

	res := make([]uint, 0)
	var i uint = 0
	for {
		p := primes[i]
		if p > n {
			break
		}
		res = append(res, p)
		i++
	}
	return res, nil
}
