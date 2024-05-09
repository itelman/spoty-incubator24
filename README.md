## SPOTY

### Author:

Ilyas Telman (@itelman).

### WARNING!!!

Due to limitations caused by the applied public API (underrecorded, low quality data, and provided invalid responses), this application handles and displays several errors and exceptions.

### Objectives

This project consists of receiving a public API and manipulating the data contained in it, in order to create a website, displaying the information. It creates and runs a server, in which it will be possible to use a web **GUI** (graphical user interface). Additionally, it allows rating on music albums.

- The website applies the following [API](https://www.last.fm/api), and consists of several endpoints:

- The first one, `index/home page`, contains the list of bands and artists in current top charts, allowing users to see more information about them.

- The second one, `/albums`, directs to the list of albums of a specific performer that was requested by the user.

- The third one, `/albuminfo`, displays the information about a specific album of the artist requested by the user, such as their name(s), image, the release date, and songs included in the repertoire.

- And the last one, `/search`, directs to search results on albums from the API, based on a query provided by the user.

- Given all this a user friendly interface was designed which displays the artists' info through data visualizations (blocks, cards, list, pages, graphics, etc).

- This project also focuses on the creation of events/actions and on their visualization.

- The event/action applied is known as a client call to the server (client-server) with features that trigger an action. This action communicates with the server in order to recieve information.

- An event consists in a system that responds to some kind of action triggered by the client, time, or any other factor.

The following HTTP requests were implemented:

1. GET requests to routes `/albums`, `/albuminfo` and `/search`: as a result the server sends HTML responses containing data. The requests accept keywords, such as `?artist=` (artist name), `?album=` (album name) and `?q=` (query for receiving search results).

The website also returns the following HTTP status codes:

- OK (200), if everything went without errors.
- Not Found (404), if a resource is not found; for example a page.
- Method Not Allowed (405), for unsupported requests (GET/POST requests).
- Internal Server Error (500), for server-side errors.

#### Ratings

Users are able to rate albums. The ratings are viewed by users themselves.

## Usage: how to run

- Download the repository to your local machine.
- Open the repository via Terminal. Make sure that the Golang compiler is installed.
- Run the following command:
```console
go run main.go
```
- Run the server on:
```console
http://localhost:8080/
```