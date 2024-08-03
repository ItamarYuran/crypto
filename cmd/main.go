package main

import (
	"fmt"

	crypto "github.com/ItamarYuran/crypto"
)

func main() {
	var z = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var w = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	t, _ := crypto.HexTo64(z)
	fmt.Println(t)
	fmt.Println(t == w)

	have1 := "1c0111001f010100061a024b53535009181c"
	have2 := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"

	xor, _ := crypto.Xor(have1, have2)
	fmt.Println(xor)
	fmt.Println(xor == want)

}
