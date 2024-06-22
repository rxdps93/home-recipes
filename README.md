# home-recipes
A simple API for my wife to access her recipes. Recipe/cooking websites are cesspools of javascript bloat and advertisements. This is a fun project for me and gives her a comfortable way to access her recipes.

I wanted to learn Go so this gives me an excuse to do so. Down the road I will be flexing my electrical engineering muscles (that I don't actually have yet) and would like to attempt to build her a small e-ink device that is purpose built for accessing these recipes.

The go-sqlite3 dependency should be installed before doing `go run .` otherwise it hangs for a while. I don't particularly like how go handles modules/libraries so far.

As of 5/30/2024: run `go run .` and navigate to localhost:8080 to see the progress
As of 6/21/2024: quit the application with CTRL-C (SIGINT). this triggers the db to wipe test data.

Overview of project structure:
* `main.go` - Contains the main function and some miscellaneous helper functions.
    * may move webserver/mux config to api code
* `recipes.db` - The sqlite3 database file
* `css/` - Contains stylesheet(s). Currently only 1 but eventually will change.
    * `style.css` - Stylesheet - will eventually be split up for maintainability
* `internal/` - Contains internal/private code for the application
    * `api/` - Where all api/endpoint code lives
        * `api.go` - Contains handler functions for endpoints
    * `db/` - Where all database io code lives
        * `db.go` - A large file containing all database related functions and project structs. May get split up.
    * `pages/` - Where all html generation code lives
        * `all_recipes.go` - Landing page for viewing all recipes
        * `all_tags.go` - Landing page for viewing all tags
        * `home.go` - Primary landing/home page for the application
        * `recipe_by_tag.go` - View all recipes with the specified tag
        * `recipe_detail.go` - View a specified recipe
        * `test.go` - Page to test new styling/structural changes before adding to real pages
        * `utils.go` - Contains all shared/helper functions
