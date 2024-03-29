package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type InfoArtists struct {
	ID           int         `json:"id"`
	Image        string      `json:"image"`
	Name         string      `json:"name"`
	Members      []string    `json:"members"`
	CreationDate int         `json:"creationDate"`
	FirstAlbum   string      `json:"firstAlbum"`
	Locations    interface{} `json:"locations"`
	ConcertDates string      `json:"concertDates"`
	Relations    string      `json:"relations"`
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

	http.HandleFunc("/search", searchHandle)

	http.HandleFunc("/suggest", suggestHandle)

	http.HandleFunc("/suggestLoc", suggestLocation)

	http.HandleFunc("/suggestFirstAlbum", suggestFirstAlbum)

	http.HandleFunc("/suggestDate", suggestDate)

	http.HandleFunc("/filter", filter)

	// Lance le serveur
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Serveur lancé avec succès  sur le port 8080")
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

	/* fmt.Println("Data Artist:", dataArtist) */ // Affichez les données de l'API debug

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

func findArtistId(apiartist string, id []int) (*InfoArtists, error) {
	artistList, err := takeArtistes(apiartist)
	if err != nil {
		return nil, err
	}

	// Recherche de l'artiste par ID
	for _, artist := range artistList {
		var count int
		if artist.ID == id[count] {
			return &artist, nil
		}
		count++
	}

	// Retourner une erreur si l'ID n'est pas trouvé
	return nil, fmt.Errorf("Artiste avec l'ID %d non trouvé", id)
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

func searchHandle(w http.ResponseWriter, r *http.Request) {

	// Récupérez les données de l'API
	API, err := takeJSON()
	if err != nil {
		log.Print("Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}
	// Recherchez les artistes dont le nom correspond à la requête
	DataArtist, err := takeArtistes(API.InfoArtists)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}
	// Récupérez la requête de recherche depuis les paramètres de l'URL
	query := r.URL.Query().Get("search")

	var result []InfoArtists

	queryWithoutSpaces := strings.ReplaceAll(query, " ", "")

	// Effectuez la recherche
	for _, artist := range DataArtist {

		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			result = append(result, artist)
			/* log.Printf("Artiste trouvé : %v\n", artist) */
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ReplaceAll(strings.ToLower(member), " ", ""), strings.ToLower(queryWithoutSpaces)) {
				result = append(result, artist)
				log.Printf("menbre trouvé : %v\n", member)
				log.Printf("données de result : %v\n", result)
			}
		}

	}

	// Affichez les résultats dans le modèle HTML
	tmpl, err := template.ParseFiles(templateHtml)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du modèle HTML", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, result)
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du modèle", http.StatusInternalServerError)
		return
	}

	return
}

func suggestHandle(w http.ResponseWriter, r *http.Request) {
	API, err := takeJSON()
	if err != nil {
		log.Print("Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	query := r.URL.Query().Get("query")

	var suggestions []string

	// Récupérez les données de l'API
	DataArtist, err := takeArtistes(API.InfoArtists)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	// Effectuez la recherche
	for _, artist := range DataArtist {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			suggestions = append(suggestions, artist.Name+" - groupe")
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
				suggestions = append(suggestions, member+" - membre")
			}
		}

	}

	if len(suggestions) == 0 {
		suggestions = append(suggestions, "Aucun résultat trouvé")
	}

	json.NewEncoder(w).Encode(suggestions)
}

func suggestLocation(w http.ResponseWriter, r *http.Request) {
	API, err := takeJSON()
	if err != nil {
		log.Print("Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	query := r.URL.Query().Get("query")

	var suggestionLoc []string
	locations := make(map[string]bool) // Utiliser une map pour stocker les localisations uniques

	// Récupérez les données de l'API
	DataLoc, err := takeLocation(API.InfoLocations)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	// Effectuez la recherche
	for _, loc := range DataLoc.Index {
		for _, location := range loc.Locations {
			locationLower := strings.ToLower(location)
			if strings.Contains(locationLower, strings.ToLower(query)) && !locations[locationLower] {
				suggestionLoc = append(suggestionLoc, location)
				locations[locationLower] = true
			}
		}
	}

	if len(suggestionLoc) == 0 {
		suggestionLoc = append(suggestionLoc, "Aucun résultat trouvé")
	}

	json.NewEncoder(w).Encode(suggestionLoc)
}

func suggestDate(w http.ResponseWriter, r *http.Request) {
	API, err := takeJSON()
	if err != nil {
		log.Print("Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	query := r.URL.Query().Get("query")

	var suggestionDate []string
	dates := make(map[string]bool) // Utiliser une map pour stocker les dates uniques

	// Récupérez les données de l'API
	DataArtist, err := takeArtistes(API.InfoArtists)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	// Effectuez la recherche
	for _, artist := range DataArtist {

		datesLower := strings.ToLower(strconv.Itoa(artist.CreationDate))
		if strings.Contains(datesLower, strings.ToLower(query)) && !dates[datesLower] {
			suggestionDate = append(suggestionDate, strconv.Itoa(artist.CreationDate))
			dates[datesLower] = true
		}
	}

	if len(suggestionDate) == 0 {
		suggestionDate = append(suggestionDate, "Aucun résultat trouvé")
	}

	json.NewEncoder(w).Encode(suggestionDate)
}

func suggestFirstAlbum(w http.ResponseWriter, r *http.Request) {
	API, err := takeJSON()
	if err != nil {
		log.Print("Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	query := r.URL.Query().Get("query")

	var suggestionFirstAlbum []string

	// Récupérez les données de l'API
	DataArtist, err := takeArtistes(API.InfoArtists)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	// Effectuez la recherche
	for _, artist := range DataArtist {
		if strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(query)) {
			suggestionFirstAlbum = append(suggestionFirstAlbum, artist.FirstAlbum)
			fmt.Println()
		}
	}

	if len(suggestionFirstAlbum) == 0 {
		suggestionFirstAlbum = append(suggestionFirstAlbum, "Aucun résultat trouvé")
	}

	json.NewEncoder(w).Encode(suggestionFirstAlbum)
}

/* partie filtre */

func filter(w http.ResponseWriter, r *http.Request) {

	API, err := takeJSON()
	if err != nil {
		log.Print("Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	// Récupérez les données de l'API
	DataArtist, err := takeArtistes(API.InfoArtists)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données depuis l'API", http.StatusInternalServerError)
	}

	// Récupérez les données de l'API

	// Récupérez les paramètres de l'URL

	check := r.URL.Query().Get("check")
	filterPlage := r.URL.Query().Get("filterPlage")
	location := r.URL.Query().Get("location")
	creaDate := r.URL.Query().Get("creaDate")
	firstAlbum := r.URL.Query().Get("FirstAlbum")

	//

	filterPlageInt, err := strconv.Atoi(filterPlage)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion de la plage de dates", http.StatusInternalServerError)
		return
	}

	var result []InfoArtists

	// Effectuez la recherche

	for _, artist := range DataArtist {

		var numMembers int

		numMembers = len(artist.Members)

		if check != "" {
			if check == "1" && numMembers == 1 {
				result = append(result, artist)
				fmt.Println("check1 result", artist)
			} else if check == "2" && numMembers == 2 {
				result = append(result, artist)
				fmt.Println("check2 result", artist)
			} else if check == "3" && numMembers == 3 {
				result = append(result, artist)
				fmt.Println("check3 result", artist)
			} else if check == "4" && numMembers == 4 {
				result = append(result, artist)
				fmt.Println("check4 result", artist)
			} else if check == "5" && numMembers == 5 {
				result = append(result, artist)
				fmt.Println("check5 result", artist)
			} else if check == "6" && numMembers == 6 {
				result = append(result, artist)
				fmt.Println("check6 result", artist)
			} else if check == "7" && numMembers == 7 {
				result = append(result, artist)
				fmt.Println("check7 result", artist)
			}
		} else {
			result = allFilter(artist, result, artist.CreationDate, filterPlageInt, location, firstAlbum, creaDate, filterPlageInt, API)
		}

	}

	// Affichez les résultats dans le modèle HTML
	tmpl, err := template.ParseFiles(templateHtml)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du modèle HTML", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, result)
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du modèle", http.StatusInternalServerError)
		return
	}

	return
}

func filtreDate(artist InfoArtists, result []InfoArtists, date int, plage int) []InfoArtists {
	if plage != 1958 {
		if date <= plage {
			result = append(result, artist)
			return result
		}
	}
	return result
}

func filterLocation(artist InfoArtists, result []InfoArtists, location string, API *GroupieTracker) []InfoArtists {

	fmt.Println("location : ", location)

	if location != "" {
		// Récupérez les données de l'API
		dataLocations, err := takeLocation(API.InfoLocations)
		if err != nil {
			return nil
		}

		var artistIDs []int

		fmt.Println("Locations trouvées : ", dataLocations)

		// Effectuez la recherche
		for _, locations := range dataLocations.Index {
			for _, loc := range locations.Locations {
				if strings.Contains(strings.ToLower(loc), strings.ToLower(location)) {
					artistIDs = append(artistIDs, locations.ID)
					fmt.Println("Locations trouvées : ", artistIDs)
				}
			}
		}

		var artistInfoLoc []InfoArtists

		for _, id := range artistIDs {
			artistInfo, err := findArtistId(API.InfoArtists, []int{id})
			if err != nil {
				return nil
			}

			artistInfoLoc = append(artistInfoLoc, *artistInfo)
		}

		return artistInfoLoc
	}
	return result
}

func filterFirstAlbum(artist InfoArtists, result []InfoArtists, firstAlbum string) []InfoArtists {
	if firstAlbum != "" {
		if strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(firstAlbum)) {
			result = append(result, artist)
			return result
		}
	}
	return result
}

func filterCreaDate(artist InfoArtists, result []InfoArtists, date string) []InfoArtists {
	if date != "" {
		if date == strconv.Itoa(artist.CreationDate) {
			result = append(result, artist)
			return result
		}
	}
	return result
}

func allFilter(artist InfoArtists, result []InfoArtists, date int, plage int, location string, firstAlbum string, creaDate string, filterPlageInt int, API *GroupieTracker) []InfoArtists {

	fmt.Println("filtre date")
	result = filtreDate(artist, result, artist.CreationDate, filterPlageInt)

	fmt.Println("filtre location")
	result = filterLocation(artist, result, location, API)
	fmt.Println("result location : ")
	fmt.Println(result)

	fmt.Println("filtre first album")
	result = filterFirstAlbum(artist, result, firstAlbum)

	fmt.Println("filtre creaDate")
	result = filterCreaDate(artist, result, creaDate)
	return result
}
