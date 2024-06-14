package app

import (
	"log"
	"math/rand"
	"time"
)

func getQuestion() []Questions {
	res := make([]Questions, 0, 6)
	for _, q := range Quiz {
		f, s := randomNumbers()
		res = append(res, q[f], q[s])
	}
	log.Println(res)
	return res
}

//func getQuestion() []Questions {
//	res := make([]Questions, 6)
//	f, s := randomNumbers()
//	p := QuizEasy[f]
//	p1 := QuizEasy[s]
//	res[0], res[1] = p, p1
//	f, s = randomNumbers()
//	p = QuizMedium[f]
//	p1 = QuizMedium[s]
//	res[2], res[3] = p, p1
//	f, s = randomNumbers()
//	p = QuizHard[f]
//	p1 = QuizHard[s]
//	res[4], res[5] = p, p1
//	return res
//}

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
