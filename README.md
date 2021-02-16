# Go Photo Gallery

Photo Gallery app with Golang

MVC
gorilla/mux

### View

Simple View cycle example:  
* A request hits App at ```"/"```.
* Router redirects the request to the ```homeHandler()``` handler.
* ```homeHandler()``` creates ```var homeView``` of type ```*views.View``` by initializing it with ```views.NewView()``` function and passing ```"bootstrap"``` and ```"views/home.gohtml"``` as arguments.
* ```views.NewView()``` takes layout template name ```"bootstrap"```, adds to it (yields) ```"views/home.gohtml"``` template and other layout template files, then parses all the passed and added files, checks for errors and puts returned result back to ```var homeView``` in ```homeHandler()``` function.  
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
    In the ```views.NewView(layoutName, template)``` function the layoutName and the template can be any template file, for eg. with ```"bootstrap.gohtml"``` template ```"home.gohtml"``` will return Home Page, with ```"bootstrap.gohtml"``` template ```"contacts.gohtml"``` will return Contacts Page.
* ```homeView.Render(w, nil)``` calls ```Template.ExecuteTemplate(w, homeView.Layout, interface{})```, which executes template, then ```homeView.Render()``` writes response w to the client.

### Controller

Simple Controller example with View:
* A request hits App at ```"/signup"```.
* Router redirects the request to the ```controllers.NewUsers().New``` handler.
* ```NewUsers()``` creates new page from  given templates with signup form by initializing it with ```views.NewView()``` function and passing ```"bootstrap"``` and ```"views/users/new.gohtml"``` as arguments.
* ```views.NewView()``` takes layout template name ```"bootstrap"```, adds to it (yields) ```"views/users/new.gohtml"``` template and other layout template files, then parses all the passed and added files, checks for errors and puts returned result to ```Users``` struct field ```NewView```, which is of type ```views.View```.
* Next ```New``` executes this template with ```NewView.Render()``` method. 
