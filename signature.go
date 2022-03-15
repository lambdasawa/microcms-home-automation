package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func verifyMicroCMSSignature(signature string, body []byte) error {
	actualMAC, err := hex.DecodeString(signature)
	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, []byte(os.Getenv("MICROCMS_WEBHOOK_SECRET")))
	mac.Write(body)
	expectedMAC := mac.Sum(nil)
	if !hmac.Equal(actualMAC, expectedMAC) {
		return err
	}

	return nil
}
