package routeur

import (
	"net/http"
	"aventure/controller"
	"fmt"
	"log"
)


func Initserv() {

	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/accueil", controller.Accueil)
	http.HandleFunc("/profil", controller.Profil)

	http.HandleFunc("/treatment", controller.FormSubmission)
	http.HandleFunc("/supprimer", controller.SupprimerPersonnage)

	http.HandleFunc("/modifier-personnage-action", controller.ModifierPersonnageAction)
	http.HandleFunc("/modifier-personnage-action-treatment", controller.ModifierPersonnageTreatment)





	fmt.Println("Serveur lancé sur le port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

}