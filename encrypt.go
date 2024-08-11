package crypto

import (
	"encoding/hex"
	"fmt"
)

const Stanza = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
const Encryption = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
const Key = "ICE"

func Encrypt(str, key string) string {
	bytes := []byte(str)
	keyBytes := []byte(key)
	fmt.Println(len(bytes), len(keyBytes))

	for i, b := range bytes {
		bytes[i] = b ^ key[i%3]
	}
	return byteToHex(bytes)

}

func ByteToHex(b []byte) string {
	return hex.EncodeToString(b)
}
