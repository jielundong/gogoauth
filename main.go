package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	c := flag.Bool("create", false, "create a new google auth with name, issuer and secret if provided, otherwise, create a new secret automatically")
	issuer := flag.String("issuer", "", "issuer")
	name := flag.String("name", "", "user name")
	secret := flag.String("s", "", "secret string for google auth")
	png := flag.Bool("png", false, "create qrcode image qr.png")
	flag.Parse()
	if *c {
		create(*secret, *name, *issuer, *png)
		return
	}
	mac := hmac.New(sha1.New, []byte(*secret))
	binary.Write(mac, binary.BigEndian, time.Now().Unix()/30)
	H := mac.Sum(nil)
	O := H[19] & 0x0f
	u := binary.BigEndian.Uint32(H[O : O+4])
	code := (u & 0x7fffffff) % 1000000
	fmt.Printf("%06d", code)
}

func create(secret, name, issuer string, png bool) {

	if secret == "" {
		rand.Seed(time.Now().Unix())
		token := make([]byte, 8)
		rand.Read(token)
		secret = hex.EncodeToString(token)
	}
	secret = base32.StdEncoding.EncodeToString([]byte(secret))
	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		panic(err)
	}
	URL.Path += "/" + issuer + ":" + name
	params := url.Values{}
	params.Add("secret", secret)
	params.Add("issuer", issuer)

	URL.RawQuery = params.Encode()
	fmt.Print(URL.String())
	if png {
		qrcode.WriteFile(URL.String(), qrcode.Medium, 256, "qr.png")
	}
}
