package adapters

import (
	"net/http"
	"studash/errors"
	"studash/tools"
)

const kPolyHost = "https://www4.polymtl.ca"

type DefaultAdapter struct {
	parser *tools.HTMLParser
}

// ----------- BELOW IS NOT REALY USED ------------ //

func GoResponse(r *http.Request) (*http.Response, error) {
	defaultClient := &http.Client{}

	// Si on arrive sur la page.
	if r.URL.Path == "/" {
		// this line is used as a debug, it's only needed on tuesdays 
		// evening because of poly server's "maintaining"
		// return defaultClient.Get("http://www.iana.org/domains/example/")

		// the regular call  that sends a request to poly's servers
		return defaultClient.Get(kPolyHost + "/poly/poly.html")
	}

	// In case we don't just want to access to the root of the "dossier etudiant"
	switch r.Method {
	case "GET":
		return defaultClient.Get(kPolyHost + r.URL.Path)
	case "POST":
		r.ParseForm()
		return defaultClient.PostForm(kPolyHost+r.URL.Path, r.Form)
	default:
		// If you get here... WTF are you trying to do ?
		// crash-test ?
		return nil, &errors.RequestError{Action: r.URL.Path, Method: r.Method, Problem: errors.ErrMethod.Error()}
	}
	return nil, nil
}
