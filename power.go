package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// test
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	args := os.Args[1:]
	if len(args) < 1 {
		panic("not enough arguments")
	}
	var print bool = false
	if len(args) == 2 {
		if args[1] == "-p" {
			print = true
		}
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	powerFile := fmt.Sprintf("%s/%s", homedir, ".power")
	brine, err := ioutil.ReadFile(powerFile)
	if err != nil {
		panic(err)
	}
	arg := args[0]
	brineBytes := []byte(brine)
	argBytes := []byte(arg)
	argBrineBytes := append(argBytes, brineBytes...)
	shah := sha512.New()
	_, err = shah.Write(argBrineBytes)
	if err != nil {
		panic(err)
	}
	brown := shah.Sum(nil)
	bbrown := base64.URLEncoding.EncodeToString(brown)[:20]
	noUnder := strings.Replace(bbrown, "_", "a", -1)
	noDashe := strings.Replace(noUnder, "-", "b", -1)
	corr := "aA1!" + noDashe
	if print {
		fmt.Print(corr)
		os.Exit(0)
	}
	bashe := os.Getenv("SHELL")
	past := fmt.Sprintf("echo -n \"%s\" | pbcopy", corr)
	err = exec.Command(bashe, "-c", past).Run()
	if err != nil {
		panic(err)
	}
}
