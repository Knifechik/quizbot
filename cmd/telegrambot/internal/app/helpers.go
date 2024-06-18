package app

import (
	"math/rand"
	"time"
)

func getQuestion() []Questions {
	res := make([]Questions, 0, 6)
	for _, q := range Quiz {
		f, s := randomNumbers()
		res = append(res, q[f], q[s])
	}
	return res
}

func randomNumbers() (int, int) {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		first := rd.Intn(5)
		second := rd.Intn(5)
		if first == second {
			continue
		} else {
			return first, second
		}
	}
}
