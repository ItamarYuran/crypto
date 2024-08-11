package crypto

import "fmt"

const HEXTABLE = "0123456789abcdef"

func Xor(st1, st2 string) (string, error) {
	if len(st1) != len(st2) {
		return "", fmt.Errorf("not of the same length")
	}
	byte1, _ := HexToByte(st1)
	byte2, _ := HexToByte(st2)
	byteret := make([]byte, len(byte1))
	for i := range byte1 {
		byteret[i] = byte1[i] ^ byte2[i]
	}
	return byteToHex(byteret), nil

}
func byteToHex(arr []byte) string {
	s := ""
	for i := 0; i < len(arr); i++ {

		high := arr[i] >> 4
		low := arr[i] & 15
		s += string(HEXTABLE[high])
		s += string(HEXTABLE[low])

	}
	return s
}
