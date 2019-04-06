package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

func Encrypt(sliceToEncrypt []byte, ceaserCount int) string{
	encryptedSlice := ""
	for i:= 0; i< len(sliceToEncrypt); i++{
		temp := int(sliceToEncrypt[i]) + ceaserCount
		encryptedSlice += string(temp)
		}
	return string(encryptedSlice)
}

func ConcatSlice(sliceToConcat []byte) string {
	newString := ""
	for i := 0; i < len(sliceToConcat); i++ {
		newString += string(sliceToConcat[i]) + "-"
	}
	return newString

}
func EZGreetings(name string) string{

		temp := greetings.PrintGreetings(name)
		return temp

	}


