package student

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Les informations de connexion au dossier étudiant.
type Credentials struct {
	// Le matricule
	Id string
	// Mot de passe
	Pass string
	// Date de naissance en format (AAMMJJ)
	Birth string
}

func BuildCredentials(r *http.Request) Credentials {
	rawJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("[ERROR] : Build Credentials ReadAll request", err)
	}
	defer r.Body.Close()

	var dictionary map[string]string
	err = json.Unmarshal(rawJson, &dictionary)
	if err != nil {
		log.Print("[ERROR] : Build Credentials Json UnMarshall", err)
	}

	return Credentials{dictionary["username"], dictionary["password"], dictionary["dateOfBirth"]}
}

func (c Credentials) UrlForm() url.Values {
	tmp := url.Values{}
	tmp.Add("code", c.Id)
	tmp.Add("nip", c.Pass)
	tmp.Add("naissance", c.Birth)
	return tmp
}

// func ParseCredentials(r *http.Request) Credentials {

// }
