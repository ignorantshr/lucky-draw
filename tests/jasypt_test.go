package test

import (
	"encoding/base64"
	"log"
	"testing"

	"astuart.co/go-jasypt"
)

func TestJasypt(t *testing.T) {
	en := "1t5nz/OjJvJKV44K6nYgPfOD0iCVYaVs"
	// en := jfNpQI9nKsUkBMkIxX0qz5Ft9T5ACtKnUgUeJBCFuxK3ofh24PbuNlnxIOr0P7Jeay81gCY3hIUTLvF5xlgVp9sAktdAjOaL
	bs, err := base64.StdEncoding.DecodeString(en)
	if err != nil {
		t.Logf("decode fail: %v", err)
	}
	d := jasypt.Decryptor{
		Algorithm: jasypt.AlgoPBEWithMD5AndDES,
		Password:  "lidc",
		// Password:  "FOO_BAR",
	}

	out, err := d.Decrypt(bs)
	if err != nil {
		log.Fatalln(err)
	}

	text := "lenovolidc"
	// text := "asdfasdfasdasdasuesrfqweafasdnlv,sdklfjasdklfsjadfklsajfksdw"
	if text != string(out) {
		t.Logf("Decrypt fail: %v", err)
	}
	log.Printf("Decrypt text: %s", out)

}
