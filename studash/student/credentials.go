package student

// Les informations de connexion au dossier étudiant.
type Credentials struct {
	// Le matricule
	Id string
	// Mot de passe
	Pass string
	// Date de naissance en format (AAMMJJ)
	Birth string
}

// func ParseCredentials(r *http.Request) Credentials {

// }
