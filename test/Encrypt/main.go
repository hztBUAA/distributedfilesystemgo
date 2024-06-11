package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"github.com/anthdm/foreverstore"
)

func main() {
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
	if _, err := foreverstore.CopyEncrypt(key, srcFile, dstFile); err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}

	fmt.Println("File encrypted successfully!")
}
