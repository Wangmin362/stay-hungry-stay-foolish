package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"unsafe"
)

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func mySha256(k string) [32]byte {
	h := sha256.New()
	h.Write(string2bytes(k))
	hashSlice := h.Sum(nil)
	var hash [32]byte
	for i := 0; i < 32; i++ {
		hash[i] = hashSlice[i]
	}
	return hash
}

func TestSha256(t *testing.T) {
	// 待计算摘要的数据
	data := "Hello, SHA-256!"

	// 创建一个SHA-256的哈希对象
	hasher := sha256.New()

	// 写入数据到哈希对象
	hasher.Write([]byte(data))

	// 计算摘要
	hash := hasher.Sum(nil)

	bytes := mySha256(data)
	fmt.Println(bytes)

	// 将摘要转换为十六进制字符串
	hashString := hex.EncodeToString(hash)

	// 输出摘要结果
	fmt.Println("SHA-256 Hash:", hashString)
}

func TestHugeNum(t *testing.T) {
	// 创建一个大整数
	var hugeNum big.Int

	// 设置大整数的值为 2^65
	hugeNum.Exp(big.NewInt(2), big.NewInt(65), nil)

	// 打印大整数的字符串表示
	fmt.Println("Huge number:", &hugeNum)
}

func TestRightMove(t *testing.T) {

	var num int64 = 1234567890
	shifted := num >> 7
	fmt.Printf("Original number: %d\n", num)
	fmt.Printf("Number after shifting right by 7 bits: %d\n, %T", shifted, shifted)

}
