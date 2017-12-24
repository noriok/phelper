package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type KeyVal struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func runPeco(pairs []KeyVal) string {
	cmd := exec.Command("peco")
	stdin, _ := cmd.StdinPipe()
	for _, pair := range pairs {
		io.WriteString(stdin, pair.Key+"\n")
	}
	stdin.Close()

	out, _ := cmd.Output()
	key := strings.TrimSpace(string(out))
	for _, pair := range pairs {
		if key == pair.Key {
			return pair.Val
		}
	}
	return ""
}

func readJSON() []KeyVal {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	text := string(bytes)
	fmt.Println(text)

	var pairs []KeyVal
	if err := json.Unmarshal(bytes, &pairs); err != nil {
		log.Fatal(err)
	}
	return pairs
}

func main() {
	pairs := readJSON()
	val := runPeco(pairs)
	fmt.Print(val)
}
