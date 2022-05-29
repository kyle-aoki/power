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
	as := os.Args[1:]
	if len(as) < 1 {
		panic("not enough arguments")
	}
	slt, err := ioutil.ReadFile("/Users/kyle/.power")
	if err != nil {
		panic(err)
	}
	uns := as[0]
	sby := []byte(slt)
	bun := []byte(uns)
	bys := append(bun, sby...)
	shf := sha512.New()
	_, err = shf.Write(bys)
	if err != nil {
		panic(err)
	}
	hsh := shf.Sum(nil)
	bsf := base64.URLEncoding.EncodeToString(hsh)[:20]
	nou := strings.Replace(bsf, "_", "", -1)
	corr := "aA1!" + nou
	sh := os.Getenv("SHELL") //fetch default shell
	pbc := fmt.Sprintf("echo -n \"%s\" | pbcopy", corr)
	err = exec.Command(sh, "-c", pbc).Run()
	if err != nil {
		panic(err)
	}
}
