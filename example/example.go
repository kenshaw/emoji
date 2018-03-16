// example/example.go
package main

import (
	"log"

	"github.com/brankas/emoji"
)

func main() {
	a := emoji.FromEmoticon(":-)")
	log.Printf(":-) %+v", a)

	b := emoji.FromAlias("slightly_smiling_face")
	log.Printf(":-) %+v", b)

	s := emoji.ReplaceEmoticonsWithAliases(":-) :D >:(")
	log.Printf("s: %s", s)

	n := emoji.ReplaceEmoticonsWithCodes(":-) :D >:(")
	log.Printf("n: %s", n)
}
