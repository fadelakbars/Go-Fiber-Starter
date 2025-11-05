package helpers

import (
	"math/rand"
)

const voucherCodeLength = 6
const numericCharset = "0123456789"

// GenerateVoucherCode generates a 6-digit numeric voucher code.
func GenerateVoucherCode() string {
	code := make([]byte, voucherCodeLength)
	for i := range code {
		code[i] = numericCharset[rand.Intn(len(numericCharset))]
	}
	return string(code)
}
