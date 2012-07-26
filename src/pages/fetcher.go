package pages

import (
	"encoding/json"
	"fmt"
)

type Functions struct {
	List []string
}

func ListFunctions(c Credentials) []byte {
	functions := Functions{[]string{"PersInfo", "Bulletin", "Horraire", "ChoixCours", "Financier"}}

	encoded, err := json.Marshal(functions)
	if err != nil {
		fmt.Println("Json Marshaling issues")
	}
	return encoded
}

func FetchInfos(c Credentials, funcName string) []byte {
	switch funcName {
	case "PersInfo":
		return []byte("Les informations personnelles de l'étudiant.")
	case "Bulletin":
		// return adapterFactory.Make(funcname).Parse(http.Get("loclaure"))
		return []byte("Le bulletin de l'Étudiant.")
	case "Horraire":
		return []byte("L'horraire de l'étudiant.")
	case "ChoixCours":
		return []byte("Les possibilités de choix de cours offertes.")
	case "Financier":
		return []byte("Toutes les informations financières de l'étudiant.")
	default:
		return []byte("Unrecognized Function")
	}
	return []byte{}
}
