package crypto

import "math"

func Hammin(st1, st2 string) int {
	sum := 0
	byt1 := []byte(st1)
	byt2 := []byte(st2)

	for i, j := range byt1 {
		if i > len(byt2) {
			sum += int(j)
			continue
		}
		sum = sum + int(math.Abs(float64(byt1[i]-byt2[i])))
	}
	return sum

}
