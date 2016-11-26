// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	. "github.com/knq/emoji"
)

const (
	gemojiURL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"
)

var (
	flagOut = flag.String("o", "gemoji_data.go", "out")
)

func main() {
	var err error

	flag.Parse()

	// generate data
	buf, err := generate()
	if err != nil {
		log.Fatal(err)
	}

	// open out
	file, err := os.Create(*flagOut)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// write
	_, err = file.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
}

var replacer = strings.NewReplacer(
	"emoji.Gemoji", "Gemoji",
	"emoji.Emoji", "\n",
	"}}", "},\n}",
	"{Emoji:", "{",
	", Description:", ", ",
	", Category:", ", ",
	", Aliases:", ", ",
	", Tags:", ", ",
	", UnicodeVersion:", ", ",
	", IOSVersion:", ", ",
)

func generate() ([]byte, error) {
	var err error

	// load gemoji data
	res, err := http.Get(gemojiURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// read all
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal
	var data Gemoji
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// add header and format
	str := fmt.Sprintf(hdr, gemojiURL, data)
	str = replacer.Replace(str)
	return format.Source([]byte(str))
}

const hdr = `
package emoji

// GemojiData is the original set of Gemoji data.
//
// see: %s
var GemojiData = %#v
`
