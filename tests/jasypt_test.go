package test

import (
	"encoding/base64"
	"log"
	"testing"

	"astuart.co/go-jasypt"
)

func TestJasypt(t *testing.T) {
	en := "nzd28bw2X+5o2E85uZE4h0FO2JIPk4Vz"
	bs, err := base64.StdEncoding.DecodeString(en)
	if err != nil {
		t.Fail()
		log.Printf("decode fail: %v", err)
	}
	d := jasypt.Decryptor{
		Algorithm: jasypt.AlgoPBEWithMD5AndDES,
		Password:  "lidc",
	}

	_, err = d.Decrypt(bs)
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	}
}
