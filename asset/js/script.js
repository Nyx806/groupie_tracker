    // déclaration des variables 

    /* variable ded recherche des groupe ou des membres */
    const searchInput =  document.getElementById('searchInput')
    const suggestion = document.getElementById('suggestions')

    /* variable pour les liste de suggestion  */
    const list = document.getElementById('list')

    /* variable pour la recherche des localisation */
    const searchInputLoc = document.getElementById('searchInputLoc')
    const suggestionLoc = document.getElementById('suggestionsLoc')

    /* variable pour la recherche des date de creation */
    const searchCreaDate = document.getElementById('searchCreaDate')
    const suggestionsCreaDate = document.getElementById('suggestionsCreaDate')

    /* variable pour la recherche des premier album */
    const searchFirstAlbum = document.getElementById('searchFirstAlbum')
    const suggestionsFirstAlbum = document.getElementById('suggestionsFirstAlbum')
    
    document.querySelector('.toggle-button').addEventListener('click', toggleFilters);

    function toggleFilters() {
        var filterMenu = document.getElementById('filterMenu');
        filterMenu.classList.toggle('show-filters');
        console.log(" click du bouton filtre ");
    }


    // permet de pouvoir selectionner les suggestion et de les coller dans le champ de recherche
    function handleSuggestionClick(suggestion) {
        // Divisez la suggestion en deux parties : le nom de l'artiste et le suffixe "-"
        var parts = suggestion.split(' - ');
        // Récupérez le nom de l'artiste
        var artistName = parts[0];
        // Collez le nom de l'artiste dans le champ de recherche
        document.getElementById('searchInput').value = artistName;
    }


    /* evenement pour la recherche des groupes et des membres  */
    searchInput.addEventListener('input', function() {
        const query = this.value.trim()
        if (query === '') {
            suggestion.innerHTML = ''
            return
        }
            // Effectue une requête AJAX pour obtenir des suggestions basées sur la requête de recherche
            fetch(`/suggest?query=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(data => {
                console.log(" voici la data : ",data)
                
                suggestion.innerHTML = ''
                data.forEach(suggestionText => {
                    const suggesEL = document.createElement('li')
                    suggesEL.textContent = suggestionText;
                    suggesEL.id = "list"
                    suggestion.appendChild(suggesEL);
                });

                // Ajoute un gestionnaire d'événements pour chaque suggestion
                suggestion.querySelectorAll('li').forEach(suggestion => {
                    suggestion.addEventListener('click', function() {
                        handleSuggestionClick(suggestion.textContent)
                        hideSuggestions(); 
                        
                    })
                })          

            })
            .catch(error => console.error('Erreur lors de la récupération des suggestions:', error));

        
                        
    });

    /* evenement pour la recherche des localisation  */
    searchInputLoc.addEventListener('input', function() {
        const query = this.value.trim()
        if (query === '') {
            suggestionLoc.innerHTML = ''
            return
        }
            // Effectue une requête AJAX pour obtenir des suggestions basées sur la requête de recherche
            fetch(`/suggestLoc?query=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(data => {
                console.log(" voici la data : ",data)
                
                suggestionLoc.innerHTML = ''
                data.forEach(suggestionText => {
                    const suggesEL = document.createElement('li')
                    suggesEL.textContent = suggestionText;
                    suggesEL.id = "list"
                    suggestionLoc.appendChild(suggesEL);
                });

                // Ajoute un gestionnaire d'événements pour chaque suggestion
                suggestionLoc.querySelectorAll('li').forEach(suggestionLoc => {
                    suggestionLoc.addEventListener('click', function() {
                        handleSuggestionClick(suggestionLoc.textContent)
                        hideSuggestions(); 
                        /* searchInputLoc.value = suggestionLoc.textContent
                        suggestionLoc.innerHTML = '' */
                    })
                })

            })
            .catch(error => console.error('Erreur lors de la récupération des suggestions:', error));
                        
    }); 


    /* evenement pour la recherche de la date de création  */
    searchCreaDate.addEventListener('input', function() {
        const query = this.value.trim()
        if (query === '') {
            suggestionsCreaDate.innerHTML = ''
            return
        }
            console.log("voici la query : ",query)
            // Effectue une requête AJAX pour obtenir des suggestions basées sur la requête de recherche
            fetch(`/suggestCreaDate?query=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(data => {
                console.log(" voici la data : ",data)
                
                suggestionsCreaDate.innerHTML = ''
                data.forEach(suggestionText => {
                    const suggesEL = document.createElement('li')
                    suggesEL.textContent = suggestionText;
                    suggesEL.id = "list"
                    suggestionsCreaDate.appendChild(suggesEL);
                });

                // Ajoute un gestionnaire d'événements pour chaque suggestion
                suggestionsCreaDate.querySelectorAll('li').forEach(suggestionsCreaDate => {
                    suggestionsCreaDate.addEventListener('click', function() {

                        handleSuggestionClick(suggestionsCreaDate.textContent)
                        hideSuggestions();

                        /* searchCreaDate.value = suggestionsCreaDate.textContent
                        suggestionsCreaDate.innerHTML = '' */
                    })
                })

            })
            .catch(error => console.error('Erreur lors de la récupération des suggestions:', error));
                        
    });

    /* evenement pour la recherche du premier album  */

    searchFirstAlbum.addEventListener('input', function() {
        const query = this.value.trim()
        if (query === '') {
            suggestionsFirstAlbum.innerHTML = ''
            return
        }
            console.log("voici la query : ",query)
            // Effectue une requête AJAX pour obtenir des suggestions basées sur la requête de recherche
            fetch(`/suggestFirstAlbum?query=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(data => {
                console.log(" voici la data : ",data)
                
                suggestionsFirstAlbum.innerHTML = ''
                data.forEach(suggestionText => {
                    const suggesEL = document.createElement('li')
                    suggesEL.textContent = suggestionText;
                    suggesEL.id = "list"
                    suggestionsFirstAlbum.appendChild(suggesEL);
                });

                // Ajoute un gestionnaire d'événements pour chaque suggestion
                suggestionsFirstAlbum.querySelectorAll('li').forEach(suggestionsFirstAlbum => {
                    suggestionsFirstAlbum.addEventListener('click', function() {

                        handleSuggestionClick(suggestionsFirstAlbum.textContent)
                        hideSuggestions();

                        searchFirstAlbum.value = suggestionsFirstAlbum.textContent
                        suggestionsFirstAlbum.innerHTML = ''
                        
                    })
                })

            })
            .catch(error => console.error('Erreur lors de la récupération des suggestions:', error));
                        
    });

    // Ajoutez une fonction pour masquer les suggestions
    function hideSuggestions() {
        suggestion.innerHTML = '';
    }

    // Ajoutez un événement click sur le document pour masquer les suggestions lorsque l'on clique à l'extérieur de la liste
    document.addEventListener('click', function(event) {
        const isSuggestion = event.target.closest('#list');
        if (!isSuggestion) {
            hideSuggestions();
        }
    });








