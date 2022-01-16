package section17

import q "github.com/stevedesilva/dsa-golang.git/exercises/section16"

func Weave(q1, q2 q.Queue) (q.Queue, error) {
	result := q.Queue{}

	for len(q1.Values) > 0 || len(q2.Values) > 0 {
		if len(q1.Values) > 0 {
			v, _ := q1.Remove()
			result.Add(v)
		}

		if len(q2.Values) > 0 {
			v, _ := q2.Remove()
			result.Add(v)
		}
	}

	return result, nil
}
