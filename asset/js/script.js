// déclaration des variables 
const searchInput =  document.getElementById('searchInput')
const suggestion = document.getElementById('suggestions')

const list = document.getElementById('list')

const searchInputLoc = document.getElementById('searchInputLoc')
const suggestionLoc = document.getElementById('suggestionsLoc')


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
                    searchInput.value = suggestion.textContent
                    suggestion.innerHTML = ''
                })
            })

        })
        .catch(error => console.error('Erreur lors de la récupération des suggestions:', error));
                     
});

/* searchInputLoc.addEventListener('input', function() {
    const query = this.value.trim()
    if (query === '') {
        suggestionLoc.innerHTML = ''
        return
    }
        // Effectue une requête AJAX pour obtenir des suggestions basées sur la requête de recherche
        fetch(`/suggest?query=${encodeURIComponent(query)}`)
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
                    searchInputLoc.value = suggestionLoc.textContent
                    suggestionLoc.innerHTML = ''
                })
            })

        })
        .catch(error => console.error('Erreur lors de la récupération des suggestions:', error));
                     
}); */



