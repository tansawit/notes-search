# Notes Search

A web app to search through my public notes. Built using [ElasticSearch](https://www.elastic.co/), [Go](https://golang.org/), [React](https://reactjs.org/) and [Docker](http://docker.com/). A work in progress.

The app is hosted at [search.tansawit.me](http://search.tansawit.me)

## Introduction

Previously, I store all my notes on a Hugo-based site hosted at [notes.tansawit.me](https://notes.tansawit.me). However, after a few month of using that solution, I ran into two problems.

1. The search functionality of my theme is limited to document title, and does not search the content
2. I found myself spending way too much time thinking about the 'optimal directory' to place each notes

Those two issues, along with an urge to learn more about ElasticSearch and React and to try to implement it, leads me to building this. With this, every word in every document will be indexed and searchable and, without a directory structure, organizing notes is no longer something that needs to be done.

## Project Structure

### Frontend (`public/`)

A react app with search functionality provided by ElasticSearch and the backend API server.

### Backend (`server/`)

This folder contains the Go files for serving the API server, preparing the notes, and populating the ElasticSearch index:

## Running Locally

> **Note** this project is still still very much incomplete. Please see the 'Current Progress' section below for current status and functionality 

You need to have [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/) installed and running on your machine.

Clone the repository and navigate into the notes-search folder in your terminal.

Run `docker-compose up` and wait for the containers to build.

Navigate to `localhost:8080` on your browser. You can now search for terms across all documents by typing into the search bar.


## Current Progress

- Backend path for querying ElasticSearch functional.
- Frontend search functional

## TODO

Frontend

- Clicking on Card shows full note
- Fledge out front-end functionality
- Input validation/sanitization

Backend

- Add highlight to search result returns
- Switch over to fully using official ElasticSearch Go Client

Deployment

- Add SSL Certificate

Content

- Migrate all notes from current directory/website to here

Misc.

- Code Comments/Cleanup

## Future Plans

- Create and API route to add notes into index and make it into another product/service to make it easier to add bits of info.
- Use backend API to create an Alfred workflow for easy searching

## Stupid Future Plans

- Desktop/Mobile App based on the backend API?
