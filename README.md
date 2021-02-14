# Go Photo Gallery

Photo Gallery app with Golang

MVC
gorilla/mux

## View

View cycle example:  
* A request hits App at ```"/"```.
* Router redirects the request to the ```homeHandler()``` handler.
* ```homeHandler()``` creates ```var homeView``` of type ```*views.View``` by initializing it with ```views.NewView()``` function and passing ```"bootstrap"``` and ```"views/home.gohtml"``` as arguments.
* ```views.NewView()``` takes layout template ```"bootstrap"```, adds to it (yields) ```"views/home.gohtml"``` template and other layout template files to it, then parses all the passed and added files, checks for errors and puts returned result to ```var homeView``` in ```homeHandler()``` function.  
```"bootstrap.gohtml"``` file acts as generic scaffold layout that yields given template file:
    ```HTML
    bootstrap.gohtml

    <body>
        <div class="container-fluid">
            {{ template "yield" . }}
        </div>
    </body>
    ```
    ```HTML
    home.gohtml

    {{ define "yield" }}
        <h1>Stuff</h1>
    {{ end }}
    ```
    In the ```views.NewView(layout, template)``` function the layout and the template can be any template file, for eg. with ```"bootstrap.gohtml"``` template ```"home.gohtml"``` will return Home Page, ```"contacts.gohtml"``` will return Contacts Page.
* ```homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)``` executes that result by writing response w to the client

