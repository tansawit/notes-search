const vm = new Vue({
    el: '#vue-instance',
    data() {
        return {
            baseUrl: 'http://localhost:3000', // API url
            searchTerm: 'Hello World', // Default search term
            searchDebounce: null, // Timeout for search bar debounce
            searchResults: [], // Displayed search results
            numHits: null, // Total search results found
            searchOffset: 0, // Search result pagination offset

            selectedParagraph: null, // Selected paragraph object
            bookOffset: 0, // Offset for book paragraphs being displayed
            paragraphs: [] // Paragraphs being displayed in book preview window
        }
    },
    async created() {
        this.searchResults = await this.search() // Search for default term
    },
    methods: {
        /** Debounce search input by 100 ms */
        onSearchInput() {
            clearTimeout(this.searchDebounce)
            this.searchDebounce = setTimeout(async () => {
                this.searchOffset = 0
                this.searchResults = await this.search()
            }, 100)
        },
        /** Call API to search for inputted term */
        async search() {
            const response = await axios.get(`${this.baseUrl}/search`, { params: { term: this.searchTerm, offset: this.searchOffset } })
            this.numHits = response.data.hits.total
            return response.data.hits.hits
        },
        /** Get next page of search results */
        async nextResultsPage() {
            if (this.numHits > 10) {
                this.searchOffset += 10
                if (this.searchOffset + 10 > this.numHits) { this.searchOffset = this.numHits - 10 }
                this.searchResults = await this.search()
                document.documentElement.scrollTop = 0
            }
        },
        /** Get previous page of search results */
        async prevResultsPage() {
            this.searchOffset -= 10
            if (this.searchOffset < 0) { this.searchOffset = 0 }
            this.searchResults = await this.search()
            document.documentElement.scrollTop = 0
        }
    }
})