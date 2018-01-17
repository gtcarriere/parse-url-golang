// parse-url.go
// A simple URL parser written in Go.
//
// Copyright (C) 2018  Gabriel Carriere
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

// Imports required packages
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type URL struct {
	Protocol  string
	Host      string
	Path      string
	Variables string
}

func unescape(s string, mode string) (string, error) {
	n := 0
	t := make([]byte, len(s)-2*n)
	j := 0
	for i := 0; i < len(s); {
		t[j] = s[i]
		j++
		i++
	}
	return string(t), nil
}

func getprotocol(rawurl string) (protocol, path string, err error) {
	for i := 0; i < len(rawurl); i++ {
		c := rawurl[i]
		switch {
		case c == ':':
			return rawurl[:i], rawurl[i+1:], nil
		}
	}
	return "", rawurl, nil
}

func split(s string, c string) (string, string) {
	i := strings.Index(s, c)
	if i < 0 {
		return s, ""
	}
	return s[:i], s[i:]
}

func Parse(rawurl string) (*URL, error) {
	u, frag := split(rawurl, "#")
	url, err := parse(u, false)

	if err != nil {
	}
	if frag == "" {
	}
	return url, nil
}

func parse(rawurl string, viaRequest bool) (*URL, error) {
	var rest string
	var err error

	url := new(URL)

	// split off leading http, mailto, or whatever, if it exists
	if url.Protocol, rest, err = getprotocol(rawurl); err != nil {
	}
	url.Protocol = strings.ToLower(url.Protocol)
	// splits between path and variables
	rest, url.Variables = split(rest, "?")

	//splits between host and path
	var splitHP string
	splitHP, rest = split(rest[2:], "/")
	url.Host = splitHP

	if err := url.path(rest); err != nil {
	}
	return url, nil
}

func (u *URL) path(p string) error {
	path, err := unescape(p, "")
	if err != nil {
		return err
	}
	u.Path = path
	return nil
}

func main() {

	// Code to get input from user (URL to parse)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the URL to parse: ")
	s, _ := reader.ReadString('\n')
	//s := "http://example.com/path/to/something?var1=1&&var2=2"

	// Parse the URL and ensure there are no errors.
	url, err := Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("Parsed URL:")
	fmt.Println("\n")
	fmt.Println("Protocol:   " + url.Protocol)
	fmt.Println("Host:       " + url.Host)
	fmt.Println("Path:       " + url.Path)
	fmt.Println("Variables:  " + url.Variables)
}
