package GestionBDD

import (
	"database/sql"
	"fmt"
	"log"
)

// Projet is the struct that holds the project information.
func GetAllProjet(db *sql.DB) ([]Projet, error) {
	rows, err := db.Query("SELECT * FROM Projet") // Executes a query that returns rows, in this case a SELECT.
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Projets []Projet // Declare the variable Projets as a slice of Projet.
	for rows.Next() {    // Iterate over the rows.
		var Projet Projet                                                                                                                                  // Declare the variable Projet as a Projet.
		if err := rows.Scan(&Projet.Id_Projet, &Projet.Nom_Projet, &Projet.Description, &Projet.Date_Debut, &Projet.Date_Fin, &Projet.Durée); err != nil { // Copies the columns in the row into the values pointed at by dest.
			return nil, err
		}
		Projets = append(Projets, Projet) // Append the project to the Projets slice.
	}

	return Projets, nil // Return the Projets slice.
}

// Gets the project information by the project ID.
func GetProjetById(db *sql.DB, id int) (Projet, error) {
	var projet Projet
	err := db.QueryRow("SELECT * FROM Projet WHERE Id_Projet = ?", id).Scan(&projet.Id_Projet, &projet.Nom_Projet, &projet.Description, &projet.Date_Debut, &projet.Date_Fin, &projet.Durée)
	if err != nil {
		log.Println(err)
	}
	return projet, nil
}

// Gets the project name by the project ID.
func GetProjectNameById(db *sql.DB, id int) (string, error) {
	var projetName string
	err := db.QueryRow("SELECT Nom_Projet FROM Projet WHERE Id_Projet = ?", id).Scan(&projetName)
	if err != nil {
		log.Println(err)
	}
	return projetName, nil
}

// Gets the project description by the project ID.
func GetProjectDescriptionById(db *sql.DB, id int) (string, error) {
	var projetDescription string
	err := db.QueryRow("SELECT Description FROM Projet WHERE Id_Projet = ?", id).Scan(&projetDescription)
	if err != nil {
		log.Println(err)
	}
	return projetDescription, nil
}

// Gets the project start date by the project ID.
func GetProjectStartDateById(db *sql.DB, id int) (string, error) {
	var projetStartDate string
	err := db.QueryRow("SELECT Date_Debut FROM Projet WHERE Id_Projet = ?", id).Scan(&projetStartDate)
	if err != nil {
		log.Println(err)
	}
	return projetStartDate, nil
}

// Gets the project end date by the project ID.
func GetProjectEndDateById(db *sql.DB, id int) (string, error) {
	var projetEndDate string
	err := db.QueryRow("SELECT Date_Fin FROM Projet WHERE Id_Projet = ?", id).Scan(&projetEndDate)
	if err != nil {
		log.Println(err)
	}
	return projetEndDate, nil
}

// Gets the project span by the project ID.
func GetProjectSpanById(db *sql.DB, id int) (string, error) {
	var projetSpan string
	err := db.QueryRow("SELECT Duree FROM Projet WHERE Id_Projet = ?", id).Scan(&projetSpan)
	if err != nil {
		log.Println(err)
	}
	return projetSpan, nil
}

// Formation is the struct that holds the formation information.
func GetAllFormations(db *sql.DB) ([]Formation, error) {
	rows, err := db.Query("SELECT * FROM Formations") // Executes a query that returns rows, in this case a SELECT.
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure that a function call is performed later in a program's execution.
	var Formations []Formation
	for rows.Next() {
		var Formation Formation
		if err := rows.Scan(&Formation.Id_Formation, &Formation.Nom_Formation, &Formation.Description, &Formation.Date_Debut, &Formation.Date_Fin, &Formation.Durée); err != nil { // Copies the columns in the row into the values pointed at by dest.
			return nil, err
		}
		Formations = append(Formations, Formation) // Append the formation to the Formations slice.
	}

	return Formations, nil
}

// Gets the formation information by the formation ID.
func GetFormationsById(db *sql.DB, id int) (Formation, error) {
	var formation Formation
	err := db.QueryRow("SELECT * FROM Formations WHERE Id_Formation = ?", id).Scan(&formation.Id_Formation, &formation.Nom_Formation, &formation.Description, &formation.Date_Debut, &formation.Date_Fin, &formation.Durée)
	if err != nil {
		if err == sql.ErrNoRows { // If there are no rows.
			return Formation{}, fmt.Errorf("aucune information trouvée pour l'ID %d", id) // Return the error message.
		}
		return Formation{}, err // Return the error message.
	}
	return formation, nil // Return the formation.
}

// Gets the formation name by the formation ID.
func GetFormationsNameById(db *sql.DB, id int) (string, error) {
	var formationName string                                                                                   // Declare the variable formationName as a string.
	err := db.QueryRow("SELECT Nom_Formation FROM Formations WHERE Id_Formation = ?", id).Scan(&formationName) // Scan copies the columns in the row into the values pointed at by dest.
	if err != nil {
		log.Println(err) // Print the error message.
	}
	return formationName, nil // Return the formationName.
}

// Gets the formation description by the formation ID.
func GetFormationsDescriptionById(db *sql.DB, id int) (string, error) {
	var formationDescription string
	err := db.QueryRow("SELECT Description FROM Formations WHERE Id_Formation = ?", id).Scan(&formationDescription)
	if err != nil {
		log.Println(err)
	}
	return formationDescription, nil
}

// Gets the formation start date by the formation ID.
func GetFormationsStartDateById(db *sql.DB, id int) (string, error) {
	var formationStartDate string
	err := db.QueryRow("SELECT Date_Debut FROM Formations WHERE Id_Formation = ?", id).Scan(&formationStartDate)
	if err != nil {
		log.Println(err)
	}
	return formationStartDate, nil
}

// Gets the formation end date by the formation ID.
func GetFormationsEndDateById(db *sql.DB, id int) (string, error) {
	var formationEndDate string
	err := db.QueryRow("SELECT Date_Fin FROM Formations WHERE Id_Formation = ?", id).Scan(&formationEndDate)
	if err != nil {
		log.Println(err)
	}
	return formationEndDate, nil
}

// Gets the formation span by the formation ID.
func GetFormationsSpanById(db *sql.DB, id int) (string, error) {
	var formationSpan string
	err := db.QueryRow("SELECT Duree FROM Formations WHERE Id_Formation = ?", id).Scan(&formationSpan)
	if err != nil {
		log.Println(err)
	}
	return formationSpan, nil
}

// Gets the login information.
func Getlogins(db *sql.DB) (string, string, error) {
	var user, pass string
	err := db.QueryRow("SELECT username, password FROM Logins").Scan(&user, &pass)
	if err != nil {
		log.Println(err)
	}
	return user, pass, nil
}
