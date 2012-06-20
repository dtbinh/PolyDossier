package pages

import (
	"encoding/json"
	"fmt"
)

type Functions struct {
	List []string
}

func ListFunctions(c Credentials) []byte {
	functions := Functions{[]string{"PersInfo", "Bullletin", "Horraire", "ChoixCours", "Financier"}}

	encoded, err := json.Marshal(functions)
	if err != nil {
		fmt.Println("Json Marshaling issues")
	}
	return encoded
}
