# Go Photo Gallery

Photo Gallery app with Golang

MVC
gorilla/mux

## View

View cycle example:  
* A request hits App at "/"
* Router redirects the request to the *homeHandler* handler
* *homeHandler* creates var *homeView* of type *views.View by initializing it with views.NewView function and passing "views/home.gohtml" as argument. This home.gohtml file is the main template layout of the requested page.
* views.NewView() takes main template file and adds other layout files to it, then parses all the passed files, checks for errors and puts returned result to var *homeView* in *homeHandler* function
* *homeView* executes that result by writing response w to the client

