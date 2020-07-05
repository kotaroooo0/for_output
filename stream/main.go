package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	docker := strings.NewReader("docker")
	teams := strings.NewReader("teams")
	zero := strings.NewReader("0")

	k := io.NewSectionReader(docker, 3, 1)
	o := io.NewSectionReader(docker, 1, 1)
	t := io.LimitReader(teams, 1)
	a := io.NewSectionReader(teams, 2, 1)
	r := io.NewSectionReader(docker, 5, 1)

	forOOOO := io.NewSectionReader(docker, 1, 1)
	pr, pw := io.Pipe()
	writer := io.MultiWriter(pw, pw, pw, pw)
	go io.Copy(writer, forOOOO)
	defer pw.Close()

	stream := io.MultiReader(k, o, t, a, r, io.LimitReader(pr, 4), zero)
	io.Copy(os.Stdout, stream)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var u User
		data, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(data, &u)
		// ...
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var u User
		err = json.NewDecoder(r.Body).Decode(&u)
		// ...
	})
}

func sample() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}
