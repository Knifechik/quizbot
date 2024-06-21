package app

import (
	"math/rand"
	"time"
)

func getQuestion() []Questions {
	res := make([]Questions, 0, 6)
	for _, q := range Quiz {
		f, s := randomNumbers(q)
		res = append(res, q[f], q[s])
	}
	return res
}

func randomNumbers(q map[int]Questions) (int, int) {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	numb := len(q)
	for {
		first := rd.Intn(numb)
		second := rd.Intn(numb)
		if first == second {
			continue
		} else {
			return first, second
		}
	}
}
