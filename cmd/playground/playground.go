package playground

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Привет!"))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Это страница /api."))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc(`/api/`, apiPage)
	mux.HandleFunc(`/`, mainPage)

	err = http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан ENV KEY")
	}

	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)

	return cipherText, nil
}

func (enc *Encrypter) Decrypt(encryptedStr []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
