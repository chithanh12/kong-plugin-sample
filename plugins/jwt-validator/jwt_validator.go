package jwt_validator

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

var (
	SecretKey                   = "-----BEGIN CERTIFICATE-----\n" + "MIICmTCCAYECBgGEhIQcMjANBgkqhkiG9w0BAQsFADAQMQ4wDAYDVQQDDAVEQUIyQjAeFw0yMjExMTcwNzM0MjFaFw0zMjExMTcwNzM2MDFaMBAxDjAMBgNVBAMMBURBQjJCMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu7o72Md6SGrKwGqTbM2E0esWYuiuANWyg5YIeef0l5nLp41vMQP28a9DiP0wdYkEbTs9FhZkZbtuIzLWtp3l1VXfTJ7TDh6UeR6aWJG+3MU5rJXKDlzdGg4f45ST5y99x7bam5qoRRhJsK1R42acpGcaaDH8Dpqtqu0wYTTuWcqZRK12uXcQMJybps2T34xWVwJi/mXc/Fl388LaBWVZv6ISuxvt7ui1WCo4OSUCHPVf8NiJ+Dj00CpYdg9q4SNd4SgpYApTDf3ndpusY70RHh+grMyNx169O5v6cJF+fwTlskyCrs9CQX8K0vrz1CmXVqBuYCcCqWxFCFU/lUYrgQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBjdLVq2d7bRGxTGUcA7MQLfVK1LzZ8Wuo+Mu4ueV8XiW4OtyafBcUFrSkAsaLYsZGxgpOj+jchIII+4fIIFHIHqEypfmfC377oV8WREiJXHbdaVla4gOg9PBUvV8qi1FpENm7eOnzVHoTaDmAGaKFMfn/KBZ4J6yLHPc7XcZIPdIk5tuk7UoU4A+DwTD7IkiKSCFAySSL77zWwdCbIZSQMoEHF4lWFDNKQkFzAgShkYt9AAPntYNMKPJZYyq22W3DKE1Kp46wcnzXjIqz8cS95sYfE1GZS2uvwC7ODv+m1ZR1uKGzyKzzq7of5x+IOTa7tXm+ab2Fmqcx+/GoD1qka" + "\n-----END CERTIFICATE-----"
	RSAPublicKey *rsa.PublicKey = nil
)

func VerifyToken(tokenStr string) (map[string]Permision, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-256-bit-secret"), nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return getPermission(token)
}

func getPermission(token *jwt.Token) (map[string]Permision, error) {
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token")
	}
	rs := map[string]Permision{}

	auth, k := claims["authorization"]
	if !k {
		return rs, nil
	}
	authStr, _ := json.Marshal(auth)
	var authObj Authorization
	err := json.Unmarshal(authStr, &authObj)
	if err != nil {
		return nil, err
	}
	for _, r := range authObj.Permissions {
		rs[r.RsName] = r
	}
	return rs, nil
}

type (
	Permision struct {
		RsName string   `json:"rsname"`
		RsID   string   `json:"rsid"`
		Scopes []string `json:"scopes"`
	}

	Authorization struct {
		Permissions []Permision `json:"permissions"`
	}
)

func init() {
	key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
	if er != nil {
		fmt.Println(er)
		return
	}

	RSAPublicKey = key
}
