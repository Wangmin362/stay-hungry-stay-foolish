package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// 生成随机的盐
func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}
	return salt, nil
}

// 加密函数
func encrypt(plaintext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 解密函数
func decrypt(ciphertext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func main() {
	plaintext := []byte("Hello, AES256!")
	key := []byte("your-key-here456o4564654se455455") // 32 bytes for AES256

	// 生成随机的盐
	salt, err := generateSalt(aes.BlockSize)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// 将盐和明文拼接起来
	plaintextWithSalt := append(salt, plaintext...)

	// 加密
	ciphertext, err := encrypt(plaintextWithSalt, key)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}

	fmt.Println("Encrypted:", ciphertext)

	// 解密
	decrypted, err := decrypt([]byte(ciphertext), key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	fmt.Println("Decrypted:", decrypted)
}
