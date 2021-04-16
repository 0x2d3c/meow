package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)

const (
	ONE   = "1"
	ZERO  = "0"
	TILDE = "~"
)

var (
	cat            Cat
	country, end   string
	decode, encode string
)

type Cat struct{}

func (m Cat) meow() string {
	switch country {
	case "cn":
		end = "呜"
		return "喵"
	case "us":
		end = "mew"
		return "meow"
	default:
		end = "呜"
		return "喵"
	}
}

func main() {
	flag.StringVar(&country, "c", "", "which country")
	flag.StringVar(&encode, "e", "", "encode message")
	flag.StringVar(&decode, "d", "", "decode message")
	flag.Parse()

	if country == "" && encode == "" && decode == "" {
		fmt.Println(`Usage of meow:
  -c string
        which country
  -d string
        decode message
  -e string
        encode message
  example:
	meow.exe -e https://github.com
	喵喵~喵~~~呜喵喵喵~喵~~呜喵喵喵~喵~~呜喵喵喵~~~~呜喵喵喵~~喵喵呜喵喵喵~喵~呜喵~喵喵喵喵呜喵~喵喵喵喵呜喵喵~~喵喵喵呜喵喵~喵~~喵呜喵喵喵~喵~~呜喵喵~喵~~~呜喵喵喵~喵~喵呜喵喵~~~喵~呜喵~喵喵喵~呜喵喵~~~喵喵呜喵喵~喵喵喵喵呜喵喵~喵喵~喵呜`)
	}

	var msg string
	switch {
	case encode != "":
		msg = enc(encode)
	case decode != "":
		msg = dec(decode)

	}
	fmt.Println(msg)
}

// enc encryption needs to be encrypted as a meow string
func enc(dst string) string {
	var buf bytes.Buffer
	for _, c := range dst {
		strBin := fmt.Sprintf("%b", c)
		strBin = strings.ReplaceAll(strBin, ONE, cat.meow())
		strBin = strings.ReplaceAll(strBin, ZERO, TILDE)
		buf.WriteString(strBin + end)
	}
	return buf.String()
}

// dec what did the cat say
func dec(msg string) string {
	msg = strings.ReplaceAll(msg, cat.meow(), ONE)
	msg = strings.ReplaceAll(msg, TILDE, ZERO)
	ms := strings.Split(msg, end)

	var buf bytes.Buffer
	for _, m := range ms {
		buf.WriteRune(strBin2Rune(m))
	}
	return buf.String()
}

// strBin2Rune binary string to utf8 character
func strBin2Rune(s string) rune {
	var n rune
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		n += rune((int(s[l-i-1]) & 0xf) << uint8(i))
	}
	return n
}
