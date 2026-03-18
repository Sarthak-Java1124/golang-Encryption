package main

import (
	"fmt"
	"os"

	"github.com/Sarthak-Java1124/golang-Encryption/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run an encrypt to encrypt , a decrypt to decrypt")
		os.Exit(1)

	}

}

func printHelp() {

}

func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info , run go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("\n Encrypting ....")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n File successfully protected")

}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info , run go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	password, _ := term.ReadPassword(0)

	fmt.Println("\n Decrypting ....")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n File successfully protected")
}
func getPassword() []byte {
	fmt.Print("Enter Password")
	password, _ := term.ReadPassword(0)
	fmt.Print("\n Confirm Password : ")
	password2, _ := term.ReadPassword(0)

	if !validatePassword(password, password2) {
		fmt.Print("\n Passwords do not match. Please try again\n")
		return getPassword()
	}
	return password

}

func validatePassword() bool {

}
func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
