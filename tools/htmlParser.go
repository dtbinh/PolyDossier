package tools

import (
	"exp/html"
	"fmt"
	"io/ioutil"
	"net/http"
	"studash/errors"
	"time"
)

func ParseHTML(w http.ResponseWriter, r *http.Request) []byte {
	t := time.Now()
	resp, err := goResponse(r)
	respCpy, err := goResponse(r)

	if err != nil {
		fmt.Println(r.Method, err)
		return WeirdJSON()
	}

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			fmt.Println("Error token")
			break
		}
	}

	data, err := ioutil.ReadAll(respCpy.Body)
	if err != nil {
		fmt.Println(r.Method, err)
		return WeirdJSON()
	}

	resp.Body.Close()
	respCpy.Body.Close()

	elapsed := time.Since(t)

	if elapsed < 200*time.Millisecond {
		fmt.Println(time.Since(t).String() + "\t" + r.URL.Path)
	} else {
		fmt.Println(time.Since(t).String() + "\t" + r.URL.Path + "\t[Trop Long!!!]")
	}

	return data
}

func WeirdJSON() []byte {
	return []byte("{name:lolsaure}")
	// return pages.KAR
}
