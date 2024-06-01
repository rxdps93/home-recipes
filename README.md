# home-recipes
A simple API for my wife to access her recipes. Recipe/cooking websites are cesspools of javascript bloat and advertisements. This is a fun project for me and gives her a comfortable way to access her recipes.

I wanted to learn Go so this gives me an excuse to do so. Down the road I will be flexing my electrical engineering muscles (that I don't actually have yet) and would like to attempt to build her a small e-ink device that is purpose built for accessing these recipes.

The go-sqlite3 dependency should be installed before doing `go run .` otherwise it hangs for a while. I don't particularly like how go handles modules/libraries so far.

As of 5/30/2024: run `go run .` and navigate to localhost:8080 to see the progress

Overview of files:
* main.go - Contains the main function and some miscellaneous helper functions.
* db.go - Contains all database related functions.
* api.go - Contains the endpoint handler functions.
* html.go - Contains functions which generate HTML for the endpoints to return.
