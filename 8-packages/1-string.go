package main

import (
	"fmt" 
	"strings"
)

func main() {
	var sentence string = "hello world"

	var containHello bool = strings.Contains(sentence, "hello")  // check contains word
	println(containHello)

	var countHello int = strings.Count(sentence, "hello")  // number of "hello"
	println(countHello)

	// check for suffix & prefix for some filename
	var filepath string = "~/Downloads/somedir/file.txt"

	var isFromDownload bool = strings.HasPrefix(filepath, "~/Downloads")
	var isTxtFile bool = strings.HasSuffix(filepath, ".txt")
	fmt.Printf("is from down? - %s\n", isFromDownload)
	fmt.Printf("is .txt file? - %s\n", isTxtFile)

	// join strings 
	var tokens []string = []string{"hello", "world", "!"}
	var joinedSentence string = strings.Join(tokens, " ")
	println(joinedSentence)

	// replace strings
	var sentence2 string = "hello peter"
	var replacedSentence string = strings.Replace(sentence2, "peter", "john", -1)  // replace all peter -> john
	println(replacedSentence)

	// split strings
	var someS3Uri string = "s3://my-bucket/my-key"
	someS3Uri = strings.Replace(someS3Uri, "s3://", "", -1)
	println(someS3Uri)

	var splittedUri []string = strings.Split(someS3Uri, "/")
	println(splittedUri[0], splittedUri[1])

	// convert string to binary
	var stringData string = "hello world"
	var arrBytes []byte = []byte(stringData)
	fmt.Println(arrBytes)

	// convert bytes back to string
	var convertedSting string = string(arrBytes)
	fmt.Println(convertedSting)
}


