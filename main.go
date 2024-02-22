package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type InfoArtists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type InfoLocations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type DatesInfo struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type InfoRelations struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type GroupieTracker struct {
	InfoArtists   string `json:"artists"`
	InfoLocations string `json:"locations"`
	DatesInfo     string `json:"dates"`
	InfoRelations string `json:"relation"`
}

func takeJSON() (*GroupieTracker, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		log.Printf("Erreur lors de la requête GET : %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	var dataAPI GroupieTracker
	err = json.NewDecoder(resp.Body).Decode(&dataAPI)
	if err != nil {
		log.Printf("Erreur lors du décodage JSON : %v\n", err)
		return nil, err
	}
	return &dataAPI, nil
}

func main() {

	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))

	http.HandleFunc("/", handler)

	// Lance le serveur
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var templateHtml = "index.html"

func handler(w http.ResponseWriter, r *http.Request) {
	dataAPI, err := takeJSON()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
		return
	}

	dataArtist, err := takeArtistes(dataAPI.InfoArtists)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(templateHtml)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du modèle HTML", http.StatusInternalServerError)
		return
	}

	// Exécuter le modèle avec les données de l'API
	err = tmpl.Execute(w, dataArtist)
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du modèle", http.StatusInternalServerError)
		return
	}

}

func takeArtistes(url string) ([]InfoArtists, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erreur lors de la requête GET : %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	var InfoArtists []InfoArtists
	err = json.NewDecoder(resp.Body).Decode(&InfoArtists)
	if err != nil {
		log.Printf("Erreur lors du décodage JSON : %v\n", err)
		return nil, err
	}
	return InfoArtists, nil
}

func takeDates(url string) (*DatesInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erreur lors de la requête GET : %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	var datesInfo DatesInfo
	err = json.NewDecoder(resp.Body).Decode(&datesInfo)
	if err != nil {
		log.Printf("Erreur lors du décodage JSON : %v\n", err)
		return nil, err
	}
	return &datesInfo, nil
}

func takeRelation(url string) (*InfoRelations, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erreur lors de la requête GET : %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	var InfoRelations InfoRelations
	err = json.NewDecoder(resp.Body).Decode(&InfoRelations)
	if err != nil {
		log.Printf("Erreur lors du décodage JSON : %v\n", err)
		return nil, err
	}
	return &InfoRelations, nil
}

func takeLocation(url string) (*InfoLocations, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erreur lors de la requête GET : %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	var InfoLocations InfoLocations
	err = json.NewDecoder(resp.Body).Decode(&InfoLocations)
	if err != nil {
		log.Printf("Erreur lors du décodage JSON : %v\n", err)
		return nil, err
	}
	return &InfoLocations, nil
}

func artistId(apiartist string, index []int) {
	artistList, err := takeArtistes(apiartist)
	if err != nil {
		log.Printf("Erreur : %v\n", err)
		return
	}

	for _, i := range index {
		artist := artistList[i]
		fmt.Println("Voici les informations sur l'artiste :", artist.Name)
		fmt.Println("ID :", artist.ID)
		fmt.Println("Image :", artist.Image)
		fmt.Println("Name :", artist.Name)
	}
}

func all(dataAPI *GroupieTracker) {
	artistList, err := takeArtistes(dataAPI.InfoArtists)
	if err != nil {
		log.Printf("Erreur : %v\n", err)
		return
	}

	locationsList, err := takeLocation(dataAPI.InfoLocations)
	if err != nil {
		log.Printf("Erreur : %v\n", err)
		return
	}

	concertDatesList, err := takeDates(dataAPI.DatesInfo)
	if err != nil {
		log.Printf("Erreur : %v\n", err)
		return
	}

	relationList, err := takeRelation(dataAPI.InfoRelations)
	if err != nil {
		log.Printf("Erreur : %v\n", err)
		return
	}

	for _, artiste := range artistList {
		fmt.Printf("Id: %d\n", artiste.ID)
		fmt.Printf("Image: %s\n", artiste.Image)
		fmt.Printf("Name: %s\n", artiste.Name)
		fmt.Printf("Members: %s\n", artiste.Members)
		fmt.Printf("CreationDate: %d\n", artiste.CreationDate)
		fmt.Printf("FirstAlbum: %s\n", artiste.FirstAlbum)

		// Vérifier la validité de l'index pour locationsList.Index
		if artiste.ID < len(locationsList.Index) {
			fmt.Printf("Locations: %s\n", locationsList.Index[artiste.ID].Locations)
		} else {
			fmt.Println("Locations: N/A")
		}

		// Vérifier la validité de l'index pour concertDatesList.Index
		if artiste.ID < len(concertDatesList.Index) {
			fmt.Printf("ConcertDates: %s\n", concertDatesList.Index[artiste.ID].Dates)
		} else {
			fmt.Println("ConcertDates: N/A")
		}

		// Vérifier la validité de l'index pour relationList.Index
		if artiste.ID < len(relationList.Index) {
			fmt.Printf("Relations: %s\n", relationList.Index[artiste.ID].DatesLocations)
		} else {
			fmt.Println("Relations: N/A")
		}

		fmt.Println("")
	}
}
