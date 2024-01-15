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


	fmt.Println("Serveur lancé sur le port 8080")
	log.Fatal(http.ListenAndServe(":8081", nil))

}