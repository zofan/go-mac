package mac

import (
	"github.com/zofan/go-fake"
	"math/rand"
	"strings"
)

func GetPrefix(vendor string) string {
	vendor = strings.ToLower(vendor)

	for _, p := range Prefixes {
		if strings.Contains(strings.ToLower(p.Name), vendor) {
			return p.Prefix
		}
	}

	return fake.RandHexString(2) + `:` + fake.RandHexString(2) + `:` + fake.RandHexString(2)
}

func RandVendor(vendor string) string {
	return GetPrefix(vendor) + `:` +
		fake.RandHexString(2) + `:` + fake.RandHexString(2) + `:` + fake.RandHexString(2)
}

func Rand(split string) string {
	return Prefixes[rand.Intn(len(Prefixes)-1)].Prefix + split +
		fake.RandHexString(2) + split + fake.RandHexString(2) + split + fake.RandHexString(2)
}

func RandDash() string {
	return Rand(`-`)
}

func RandColon() string {
	return Rand(`:`)
}
