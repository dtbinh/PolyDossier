package tools

import (
	"exp/html"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const kPolyHost = "https://www4.polymtl.ca"

func goResponse(r *http.Request) (*http.Response, error) {
	defaultClient := &http.Client{}

	// Si on arrive sur la page.
	if r.URL.Path == "/" {
		return defaultClient.Get("http://www.iana.org/domains/example/")
		// return defaultClient.Get(kPolyHost + "/poly/poly.html")
	}

	// switch r.Method {
	// case "GET":
	// return defaultClient.Get(kPolyHost + r.URL.Path)
	// case "POST":
	// r.ParseForm()
	// return defaultClient.PostForm(kPolyHost+r.URL.Path, r.Form)
	// default:
	// return nil, &errors.RequestError{Method: r.Method, Ierror: errors.ErrMethod}
	// }
	return nil, nil
}

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
