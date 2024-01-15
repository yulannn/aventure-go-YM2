package controller

import (
	"aventure/templates"
	"net/http"
	"aventure/backend"
	"fmt"
	"encoding/json"
	"os"
	"io/fs"
	"strconv"
	"io/ioutil"
)

func FormSubmission(w http.ResponseWriter, r *http.Request) {

    nomFichier := "perso.json"

    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
        return
    }

	taille, errTaille := strconv.Atoi(r.FormValue("taille"))
    poids, errPoids := strconv.Atoi(r.FormValue("poids"))
    age, errAge := strconv.Atoi(r.FormValue("age"))

    if errTaille != nil || errPoids != nil || errAge != nil {
        http.Error(w, "Erreur lors de la conversion des champs en types appropriés", http.StatusInternalServerError)
        return
    }

	existingIDs, _ := backend.PersoIDs("perso.json")

	var nouveauID int
    for {
        nouveauID = backend.RandomID()
        if !backend.Contains(existingIDs, nouveauID) {
            break
        }
    }

	nouveauIDString := strconv.Itoa(nouveauID)


    form := backend.Personnage{
		ID: 			nouveauIDString,
        Nom:            r.FormValue("nom"),
        Sexe:           r.FormValue("sexe"),
        Motivation: 	r.FormValue("motivation"),
        Personnalite:   r.FormValue("personnalite"),
		Taille:          taille,
        Poids:           poids,
        Age:             age,
    }

    var dataForms backend.Personnages

	file, _ := ioutil.ReadFile("perso.json")

	json.Unmarshal(file, &dataForms)

    dataForms.PersonnagesData = append(dataForms.PersonnagesData, form)

    dataWrite, errWrite := json.Marshal(dataForms)
    if errWrite != nil {
        http.Error(w, fmt.Sprintf("Erreur lors du marshal du fichier : %v", errWrite), http.StatusInternalServerError)
        return
    }

    errWriteFile := os.WriteFile(nomFichier, dataWrite, fs.FileMode(0644))
    if errWriteFile != nil {
        http.Error(w, fmt.Sprintf("Erreur lors de l'écriture du fichier : %v", errWriteFile), http.StatusInternalServerError)
        return
    }

    fmt.Println("Ajouté avec succès")
    http.Redirect(w, r, "http://localhost:8081/profil", http.StatusSeeOther)
}

func Accueil(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}

func Profil(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "profil", nil)
}