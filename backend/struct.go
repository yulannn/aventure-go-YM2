package backend


type Personnage struct {
	ID			string   `json:"id"`
	Nom             string `json:"nom"`
	Sexe            string `json:"sexe"`
	Motivation      string `json:"motivation"`
	Personnalite    string `json:"personnalite"`
	Taille          int    `json:"taille"`
	Poids           int    `json:"poids"`
	Age             int    `json:"age"`
}

type Personnages struct {
	PersonnagesData []Personnage `json:"personnages"`
}