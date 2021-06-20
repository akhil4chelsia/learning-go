package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "Hello world!"
	//String length
	fmt.Println(len(s))

	//Iterate over each characters
	for _, ch := range s {
		fmt.Println(string(ch))
	}

	//Use operators with string
	fmt.Println("dog" < "cat")
	fmt.Println("dog" == "Dog")
	fmt.Println("dog" == "dog")
	fmt.Println("dog" > "cat")

	//string comparison, return 1, -1, 0
	fmt.Println(strings.Compare("Dog", "dog"))
	fmt.Println(strings.Compare("Dog", "Dog"))
	fmt.Println(strings.Compare("Dog", "Cat"))

	//String equalfold - ignore case, returns boolean
	fmt.Println(strings.EqualFold("Hello", "hello"))

	//case conversion
	s2 := "HeLLo WoRlD"
	fmt.Println(strings.ToLower(s2), ",", strings.ToUpper(s2))

	//String searching
	s3 := "the quick brown fox, jumps over the lazy dog."
	fmt.Println(strings.Contains(s3, "jump"))    //return boolean
	fmt.Println(strings.ContainsAny(s3, "abcd")) // boolean
	fmt.Println(strings.Index(s3, "jump"))       // returns index

	//Prefix and suffix
	file := "overview.pdf"
	fmt.Println(strings.HasSuffix(file, ".pdf"))
	fmt.Println(strings.HasPrefix(file, "over"))

	//Count substring
	fmt.Println(strings.Count(s3, "the"))

	//String manipulation
	// Split and Join
	splitted := strings.Split(s3, " ")
	for index, word := range splitted {
		fmt.Print(index, ":", word, ",")
	}
	fmt.Println()
	fmt.Println(strings.Join(splitted, "-")) // joining

	//Fields - similar to split but operates only on white spaces
	fmt.Println(strings.Fields(s3)) //returns array splitted by white space

	//FieldsFunc - is a customizable version of Fields that uses a call back to decide where to split
	result := strings.FieldsFunc(s3, func(r rune) bool {
		return unicode.IsPunct(r) //Split the string on punctuation
	})
	for _, each := range result {
		fmt.Println(strings.TrimSpace(each))
	}

	//replace
	replacer := strings.NewReplacer(
		".", "|",
		",", "|")
	fmt.Println(replacer.Replace(s3))

	//Map function used to transform
	transform := func(r rune) rune {
		return r
	}
	r := strings.Map(transform, s3)
	fmt.Println(r)

	//String builder
	var sb strings.Builder
	sb.WriteString("this is string 1\n")
	sb.WriteString("this is string 2\n")
	sb.WriteString("this is string 3\n")
	fmt.Println(sb.String())
	fmt.Println(sb.Cap())
	sb.Grow(1024)
	fmt.Println(sb.Cap())
	for i := 0; i < 10; i++ {
		fmt.Fprint(&sb, "String %d", i)
	}
	fmt.Println(sb.String())
	sb.Reset() // clear and reuse
	fmt.Println(sb.Cap())
}
