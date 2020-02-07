## Backend Server

This folder contains the Go files for serving the API server, preparing the notes, and populating the ElasticSearch index:

### Directories

- `notes/` stores all of the notes exported and converted from my previous site

### Go Files

- `app.go` calls the functions to populate the ElasticSearch index 'library' and run the API server.
- `search.go` are functions to handle the searching and search result handling
- `connection.go` contains functions used when connecting to ElasticSearch.
- `load.go` helps with parsing and inserting the notes into the index.

### Scripts

- `markdown-to-text.sh` for converting the current notes from Markdown to plaintext and storing it in the `notes` folder.
- `wait-for-elasticsearch.sh` runs when the backend container start. It pings the ElasticSearch node and prevents the server from starting until the node is online. This prevents the server throwing an error when connecting to a booting node.


