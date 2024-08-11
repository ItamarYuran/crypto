package main

import (
	"crypto"
	"fmt"
)

func main() {

	///////// SINGLE XOR
	// cipher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	// lst := crypto.SingleXor(cipher)
	// for _, str := range lst {
	// 	fmt.Println(str)
	// }

	////// MULTY XOR
	// word := crypto.GetTxt()
	// words := strings.Split(word, "\n")
	// var list []string
	// for _, r := range words {
	// 	ls := crypto.SingleXor(r)
	// 	list = append(list, ls...)
	// }

	// for _, r := range list {
	// 	fmt.Println(r)
	// }

	///////Encrypt
	// f := crypto.Encrypt(crypto.Stanza, crypto.Key)
	// fmt.Println(f)
	// fmt.Println(f == crypto.Encryption)

	one := "this is a test"
	two := "wokka wokka!!!"
	in := crypto.Hammin(one, two)
	fmt.Println(in)
}
