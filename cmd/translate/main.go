package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/wzshiming/translate"
)

var (
	from  = "auto"
	to    string
	out   string
	files []string
)

func init() {
	flag.StringVar(&from, "from", from, "from lang")
	flag.StringVarP(&to, "to", "t", to, "to lang")
	flag.StringVarP(&out, "out", "o", out, "out")
	flag.StringSliceVarP(&files, "file", "f", files, "file")
	flag.Usage = func() {
		w := os.Stderr
		fmt.Fprintf(w, "Translate:\n")
		for i := 0; ; i++ {
			v, ok := translate.GoogleCodeMap[translate.GoogleCode(i)]
			if !ok {
				break
			}
			fmt.Fprintf(w, "   %s\n", v[1])
		}
		fmt.Fprintf(w, "Examples:\n")
		fmt.Fprintf(w, "    %s [Options] {text}\n", os.Args[0])
		fmt.Fprintf(w, "    %s -to zh-CN hello\n", os.Args[0])
		fmt.Fprintf(w, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	if to == "" {
		flag.Usage()
		return
	}

	text := []byte{}
	args := flag.Args()
	if len(args) != 0 {
		text = []byte(strings.Join(args, " "))
		text = append(text, '\n')
	}
	for _, file := range files {
		body, err := openFile(file)
		if err != nil {
			log.Println(err)
			continue
		}
		text = append(text, body...)
		text = append(text, '\n')
	}
	text = bytes.TrimSpace(text)
	if len(text) == 0 {
		return
	}

	ret, err := translate.GoogleTranslate(string(text), from, to)
	if err != nil {
		log.Println(err)
		return
	}
	if out == "" {
		fmt.Println(ret)
	} else {
		err = ioutil.WriteFile(out, []byte(ret), 0644)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func openFile(file string) ([]byte, error) {
	if file == "-" {
		return ioutil.ReadAll(os.Stdin)
	}
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
