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


func ModifierPersonnageTreatment(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")


    data, err := os.ReadFile("./perso.json")
    if err != nil {
        http.Error(w, "Erreur lors de la lecture du fichier", http.StatusInternalServerError)
        return
    }

    var personnages backend.Personnages
    err = json.Unmarshal(data, &personnages)
    if err != nil {
        http.Error(w, "Erreur lors du décodage du JSON", http.StatusInternalServerError)
        return
    }
    var indexPerso int
    for index, p := range personnages.PersonnagesData {
        if p.ID == id {
            indexPerso = index
            break
        }
    }


    taille, errTaille := strconv.Atoi(r.FormValue("taille"))
    poids, errPoids := strconv.Atoi(r.FormValue("poids"))
    age, errAge := strconv.Atoi(r.FormValue("age"))

    if errTaille != nil || errPoids != nil || errAge != nil {
        http.Error(w, "Erreur lors de la conversion des champs en types appropriés", http.StatusInternalServerError)
        return
    }

    personnages.PersonnagesData[indexPerso].Nom = r.FormValue("nom")
    personnages.PersonnagesData[indexPerso].Sexe = r.FormValue("sexe")
    personnages.PersonnagesData[indexPerso].Motivation = r.FormValue("motivation")
    personnages.PersonnagesData[indexPerso].Personnalite = r.FormValue("personnalite")
    personnages.PersonnagesData[indexPerso].Taille = taille
    personnages.PersonnagesData[indexPerso].Poids = poids
    personnages.PersonnagesData[indexPerso].Age = age

    dataWrite, err := json.Marshal(personnages)
    if err != nil {
        http.Error(w, "Erreur lors de la sérialisation du JSON", http.StatusInternalServerError)
        return
    }

    err = os.WriteFile("./perso.json", dataWrite, 0644)
    if err != nil {
        http.Error(w, "Erreur lors de l'écriture du fichier", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/profil", http.StatusSeeOther)

}






func ModifierPersonnageAction(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")


    data, err := os.ReadFile("./perso.json")
    if err != nil {
        http.Error(w, "Erreur lors de la lecture du fichier", http.StatusInternalServerError)
        return
    }

    var personnages backend.Personnages
    err = json.Unmarshal(data, &personnages)
    if err != nil {
        http.Error(w, "Erreur lors du décodage du JSON", http.StatusInternalServerError)
        return
    }
    var indexPerso int
    for index, p := range personnages.PersonnagesData {
        if p.ID == id {
            indexPerso = index
            break
        }
    }


    templates.Temp.ExecuteTemplate(w, "modifier-personnage", personnages.PersonnagesData[indexPerso])


}


func SupprimerPersonnage(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
        return
    }

    idASupprimer := r.FormValue("id")

    var personnages backend.Personnages
    data, err := ioutil.ReadFile("./perso.json")
    if err != nil {
        http.Error(w, "Erreur lors de la lecture du fichier", http.StatusInternalServerError)
        return
    }

    err = json.Unmarshal(data, &personnages)
    if err != nil {
        http.Error(w, "Erreur lors du décodage du JSON", http.StatusInternalServerError)
        return
    }

    for i, personnage := range personnages.PersonnagesData {
        if personnage.ID == idASupprimer {
            personnages.PersonnagesData = append(personnages.PersonnagesData[:i], personnages.PersonnagesData[i+1:]...)
            break
        }
    }

    dataWrite, err := json.Marshal(personnages)
    if err != nil {
        http.Error(w, "Erreur lors de la sérialisation du JSON", http.StatusInternalServerError)
        return
    }

    err = ioutil.WriteFile("./perso.json", dataWrite, 0644)
    if err != nil {
        http.Error(w, "Erreur lors de l'écriture du fichier", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/profil", http.StatusSeeOther)
}



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
    file, erreurFile := ioutil.ReadFile("./perso.json")
    if erreurFile != nil {
        fmt.Println(erreurFile)
    }
    var ListePersonnages backend.Personnages

    json.Unmarshal(file, &ListePersonnages)
	templates.Temp.ExecuteTemplate(w, "profil", ListePersonnages)
}