package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"os"
	"testing"
)

func TestCopyEncryptDecrypt(t *testing.T) {
	payload := "Foo not bar"
	// dst = src + 16(iv)
	src := bytes.NewReader([]byte(payload))
	dst := new(bytes.Buffer)
	key := newEncryptionKey()
	_, err := copyEncrypt(key, src, dst)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(payload))
	fmt.Println(len(dst.String()))

	out := new(bytes.Buffer)
	nw, err := copyDecrypt(key, dst, out)
	if err != nil {
		t.Error(err)
	}

	if nw != 16+len(payload) {
		t.Fail()
	}

	if out.String() != payload  {
		t.Errorf("decryption failed!!!")
	}
}

func TestMy1(t *testing.T) {
	// Generate a new 32-byte encryption key
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// Open the source file
	srcFile, err := os.Open("source.txt")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create("encrypted.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer dstFile.Close()

	// Encrypt the source file and write the encrypted data to the destination file
	if _, err := copyEncrypt(key, srcFile, dstFile); err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}

	fmt.Println("File encrypted successfully!")

}

func TestMy2(t *testing.T) {
	// Generate a new 32-byte encryption key
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// Open the source file
	srcFile, err := os.Open("source.txt")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create("encrypted.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer dstFile.Close()

	// Encrypt the source file and write the encrypted data to the destination file
	if _, err := copyEncrypt(key, srcFile, dstFile); err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}

	fmt.Println("File encrypted successfully!")

}
