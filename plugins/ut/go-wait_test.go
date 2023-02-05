package ut

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

var SecretKey = "-----BEGIN CERTIFICATE-----\n" + "MIICmTCCAYECBgGEhIQcMjANBgkqhkiG9w0BAQsFADAQMQ4wDAYDVQQDDAVEQUIyQjAeFw0yMjExMTcwNzM0MjFaFw0zMjExMTcwNzM2MDFaMBAxDjAMBgNVBAMMBURBQjJCMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu7o72Md6SGrKwGqTbM2E0esWYuiuANWyg5YIeef0l5nLp41vMQP28a9DiP0wdYkEbTs9FhZkZbtuIzLWtp3l1VXfTJ7TDh6UeR6aWJG+3MU5rJXKDlzdGg4f45ST5y99x7bam5qoRRhJsK1R42acpGcaaDH8Dpqtqu0wYTTuWcqZRK12uXcQMJybps2T34xWVwJi/mXc/Fl388LaBWVZv6ISuxvt7ui1WCo4OSUCHPVf8NiJ+Dj00CpYdg9q4SNd4SgpYApTDf3ndpusY70RHh+grMyNx169O5v6cJF+fwTlskyCrs9CQX8K0vrz1CmXVqBuYCcCqWxFCFU/lUYrgQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBjdLVq2d7bRGxTGUcA7MQLfVK1LzZ8Wuo+Mu4ueV8XiW4OtyafBcUFrSkAsaLYsZGxgpOj+jchIII+4fIIFHIHqEypfmfC377oV8WREiJXHbdaVla4gOg9PBUvV8qi1FpENm7eOnzVHoTaDmAGaKFMfn/KBZ4J6yLHPc7XcZIPdIk5tuk7UoU4A+DwTD7IkiKSCFAySSL77zWwdCbIZSQMoEHF4lWFDNKQkFzAgShkYt9AAPntYNMKPJZYyq22W3DKE1Kp46wcnzXjIqz8cS95sYfE1GZS2uvwC7ODv+m1ZR1uKGzyKzzq7of5x+IOTa7tXm+ab2Fmqcx+/GoD1qka" + "\n-----END CERTIFICATE-----"

func TestJwt(t *testing.T) {
	n := time.Now()
	tokenStr := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfQnp2QnNfRm1wRmY1OTI5d09NcTVva3RxNXE3ZTJfdDdENWxubHIyVlFvIn0.eyJleHAiOjE2NzU1NzIwNzMsImlhdCI6MTY3NTU3MDI3MywianRpIjoiNTIyZmIzMWEtYmMzMi00OGVjLThhY2ItYjM5Yjc1ODliNmRlIiwiaXNzIjoiaHR0cHM6Ly9rZXljbG9hay1zaXQucGxhdGZvcm0ubnAuZG9jdG9yYW55d2hlcmUuY29tL3JlYWxtcy9EQUIyQiIsInN1YiI6Ijk2NmRiNDFjLWNkZDktNDc3OS04N2RhLWQ1NTcwMjEyY2U1OSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImIyYi5pbnN1cmVyLnBvcnRhbCIsInNlc3Npb25fc3RhdGUiOiJiMzAxM2ZiNi0yNzJmLTQwYzEtODViZS02ODVkZDI5NGUyZmIiLCJhY3IiOiIxIiwicmVzb3VyY2VfYWNjZXNzIjp7ImIyYi5pbnN1cmVyLnBvcnRhbCI6eyJyb2xlcyI6WyJyb2xlLmluc3VyZXJhZG1pbiJdfX0sInNjb3BlIjoiZW1haWwgYjJiLmluc3VyZXJ1c2Vyc2NvcGVzIHByb2ZpbGUiLCJzaWQiOiJiMzAxM2ZiNi0yNzJmLTQwYzEtODViZS02ODVkZDI5NGUyZmIiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5hbWUiOiJJbnN1cmVyIDAxIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiaW5zdXJlcjAxQGdtYWlsLmNvbSIsImdpdmVuX25hbWUiOiJJbnN1cmVyIDAxIiwiZmFtaWx5X25hbWUiOiIiLCJlbWFpbCI6Imluc3VyZXIwMUBnbWFpbC5jb20ifQ.UP_lPpcDUBK-bZ8rg0mbuLcE4XnqQhHAhP8eM4XMEqibgnER9apdf-Vgv7Tud5RHpxWGWpunzCK9JxsvPAVPWDkfXEoEkVjQcR9hGwaaTo001J3XXecq_x1yvsQmXOnViVsR0EJvnGHtc7DSRQ8kDXGdaaGQUO7ajJOlGLefSs4yLanoa95zAXI_s5IOlsg9fiwOc-GeN0tWrJIQaIcavX28-M4T9r4IfVJ0AkSPmNbYAczUXEVO-gdW_P-FXJs4M0W6mYsJFAHIkFRdezrVYeh1_DOiJmvJEB-OoNEZYwvYMhXtsBLRLqbzyFGvEO78AgM1SVV9Jubv4bII0AhKpA"
	key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
	if er != nil {
		fmt.Println(er)
		return
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	
	if token.Claims.Valid() == nil {

	}
	val, _ := json.MarshalIndent(token, "", "  ")
	fmt.Println(string(val))
	fmt.Printf("Duration %v", time.Now().Sub(n))
}
