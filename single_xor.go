package crypto

import (
	"strings"
)

func TryAll(str string) [][]byte {
	var strings []string
	var bytes [][]byte

	for i := 0; i < 256; i++ {
		low := byte(i) & 15
		high := (i & 240) >> 4
		s := string(HEXTABLE[low])
		t := string(HEXTABLE[high])
		byt := s + t
		stringo := ""

		for p := 0; p < len(str); p = p + 2 {
			stringo += byt

		}
		strings = append(strings, stringo)
	}

	for i, st := range strings {
		strings[i], _ = Xor(str, st)
		b, _ := HexToByte(strings[i])
		bytes = append(bytes, b)
	}
	return bytes
}

func Distribution(strs []string) []string {
	m := make(map[string]int)
	m["a"] = 7
	m["e"] = 6
	m["i"] = 4
	m["o"] = 4
	m[" "] = 3

	var toRet []string

	for _, str := range strs {
		b := true
		length := len(str)
		for key, val := range m {

			val = int(float64(val) * float64(length) / 100)

			appearence := strings.Count(str, key)

			if appearence < val {
				b = false
			}
		}
		if b {
			toRet = append(toRet, str)
		}
	}
	return toRet

}

func SingleXor(cipher string) []string {

	bytes := TryAll(cipher)
	var strs []string
	for _, b := range bytes {
		strs = append(strs, strings.ToLower(string(b)))

	}
	new := Distribution(strs)

	return new

}
