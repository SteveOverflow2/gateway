package rest

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func VerifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			tokenString := getTokenString(r)
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("Wrong authorization header format, missing: Bearer "))
				if err != nil {
					return
				}
				return
			}

			accessToken := &jwt.StandardClaims{}
			token, err := jwt.ParseWithClaims(tokenString, accessToken, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, errors.New("Unauthorized")
				}
				pub, err := convertToRsaPublicKey("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt2pzou+xKaZrocBNe3ZGuyJRveQE5PvLRovGw6NMFG0u1d26qx8cIYzSC667JxjuuEzgtST3XvDQSxX+YYBdrd0lYxrkbCzmlPzDbVCUzhOVjRAReg+k/W8Y8KdTZmviTZQPw5Bpgi2mF29nL7ilFQunnnfahhtRtstqAePqaxpVhNin5Tz8f6Z4Q+3gNhNFIvL6ZyugfONFbWA4zCX+y5vksUBnShWHbZKWZRNnNpQ8vcNITlnwpNCFbC/zsFn+BoaOpC5hu9Yi90WiB+MWANFAlXaZFLloMh93t8FuGe1rp7uBydVEDfzTRLbKzaTSSrpXmCk8rVK2ilicd2+a6QIDAQAB")
				if err != nil {
					return nil, err
				}
				return pub, nil
			})
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				_, err2 := w.Write([]byte("Error parsing the JWT"))
				if err2 != nil {
					return
				}
				return
			}
			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("Validation error"))
				if err != nil {
					return
				}
				return
			}
			fmt.Println("AUTH SUCCES")
			r.Header.Set("Authorization", tokenString)
			endpointHandler(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("No Authorization token"))
			if err != nil {
				return
			}
		}
	})
}

func getTokenString(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return ""
	}
	return tokenString[7:]
}

func convertToRsaPublicKey(secret string) (interface{}, error) {
	line_length := 64
	var lines []string
	for i := 0; i <= len(secret); i += line_length {
		if (i + line_length) < len(secret) {
			lines = append(lines, (secret[i:i+line_length] + "\n"))
		} else {
			lines = append(lines, (secret[i:] + "\n"))
			break
		}
	}
	key := "-----BEGIN PUBLIC KEY-----\n" + "" + strings.Join(lines, "") + "-----END PUBLIC KEY-----"
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}
	return pub, nil
}
