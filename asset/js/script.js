// déclaration des variables 
const searchInput =  document.getElementById('searchInput')
const suggestion = document.getElementById('suggestions')


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
                suggestion.appendChild(suggesEL);
            });
        })
        .catch(error => console.error('Erreur lors de la récupération des suggestions:', error));
});