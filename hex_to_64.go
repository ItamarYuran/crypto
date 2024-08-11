package crypto

import "fmt"

const TABLE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func HexDigitToInt(digit byte) (int, error) {
	switch {
	case digit >= '0' && digit <= '9':
		return int(digit - '0'), nil
	case digit >= 'a' && digit <= 'f':
		return int(digit - 'a' + 10), nil
	case digit >= 'A' && digit <= 'F':
		return int(digit - 'A' + 10), nil
	default:
		return 0, fmt.Errorf("invalid hex digit: %c", digit)
	}
}

func HexToByte(hex string) ([]byte, error) {
	if len(hex)%2 != 0 {
		return nil, fmt.Errorf("errrror")
	}
	bytes := make([]byte, len(hex)/2)
	for i := 0; i < len(hex); i = i + 2 {
		high, err := HexDigitToInt(hex[i])
		if err != nil {
			return nil, fmt.Errorf("wont work")
		}
		low, err := HexDigitToInt(hex[i+1])
		if err != nil {
			return nil, fmt.Errorf("wont work")
		}
		bytes[i/2] = byte(high<<4 | low)

	}
	return bytes, nil
}

func byteTo64(arr []byte) (string, error) {
	if len(arr)%3 != 0 {
		fmt.Println(len(arr) % 3)
		return "", fmt.Errorf("nope body")
	}
	s := ""
	for i := 0; i < len(arr); i = i + 3 {
		high := arr[i]
		mid := arr[i+1]
		low := arr[i+2]
		c1 := (high >> 2) & 63
		c2 := (high&3)<<4 | (mid >> 4 & 15)
		c3 := ((mid & 15) << 2) | (low >> 6 & 3)
		c4 := low & 63
		s += string(TABLE[c1])
		s += string(TABLE[c2])
		s += string(TABLE[c3])
		s += string(TABLE[c4])

	}
	return s, nil

}
func HexTo64(str string) (string, error) {
	bytes, err := HexToByte(str)
	if err != nil {
		return "", fmt.Errorf(" here nope")
	}
	sixtyfour, err := byteTo64(bytes)
	if err != nil {
		return "", err
	}
	return sixtyfour, nil

}
