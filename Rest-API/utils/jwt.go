package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

var secret_key = "85336e5c0da97b2da919f6760743283ccb6f3867498f33fa98c052be07326607"
var expiration_time int64 = 3600 * 24 * 5

func GerarJWT(username string) string {
	now := strconv.Itoa(int(time.Now().Unix()))
	exp := strconv.Itoa(int(time.Now().Unix() + expiration_time))

	header := map[string]string{
		"typ": "JWT",
		"alg": "HS256",
	}
	claims := map[string]string{
		"username": username,
		"iat":      now,
		"exp":      exp,
		"iss":      "devmatheusguerra",
	}
	return gerarToken(header, claims)
}

func gerarToken(header, claims map[string]string) string {
	// Header to JSON
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return ""
	}

	// Claims to JSON
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return ""
	}

	// Header to Base64
	headerBase64 := base64.StdEncoding.EncodeToString(headerJSON)
	// Claims to Base64
	claimsBase64 := base64.StdEncoding.EncodeToString(claimsJSON)
	// Concatenate
	token := headerBase64 + "." + claimsBase64
	// Hash token
	hash := hmac.New(sha256.New, []byte(secret_key))
	hash.Write([]byte(token))
	// Hash to Base64
	hashBase64 := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	// Concatenate
	token = token + "." + hashBase64

	return token
}

func Validar(token string) bool {
	// Split token
	splitToken := strings.Split(token, ".")
	// Hash token
	hash := hmac.New(sha256.New, []byte(secret_key))
	hash.Write([]byte(splitToken[0] + "." + splitToken[1]))
	// Hash to Base64
	hashBase64 := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	// Compare
	if hashBase64 != splitToken[2] {
		return false
	}
	return true
}

func RenovarJWT(token string) string {
	if !Validar(token) {
		return ""
	}
	username := obterUsuarioDoToken(token)
	return GerarJWT(username)

}

func obterUsuarioDoToken(token string) string {
	splitToken := strings.Split(token, ".")
	claimsJSON, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return ""
	}
	var claims map[string]string
	err = json.Unmarshal(claimsJSON, &claims)
	if err != nil {
		return ""
	}
	return claims["username"]
}

func VerificarSeTokenEstaExpirado(token string) bool {
	splitToken := strings.Split(token, ".")
	claimsJSON, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false
	}
	var claims map[string]string
	err = json.Unmarshal(claimsJSON, &claims)
	if err != nil {
		return false
	}
	exp := claims["exp"]
	if exp == "" {
		return false
	}
	expInt, err := strconv.Atoi(exp)
	if err != nil {
		return false
	}
	if time.Now().Unix() > int64(expInt) {
		return true
	}
	return false
}
