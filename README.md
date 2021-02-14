# Go Photo Gallery

Photo Gallery app with Golang

MVC
gorilla/mux

## View

View cycle example:  
* A request hits App at <span style="color:#2aa198">"/"</span>.
* Router redirects the request to the <span style="color:#8f3f71">homeHandler</span>() handler.
* <span style="color:#8f3f71">homeHandler</span>() creates ```var``` <span style="color:#268bd2">homeView</span> of type ```*```views.<span style="color:#268bd2">View</span> by initializing it with views.<span style="color:#8f3f71">NewView</span>() function and passing <span style="color:#2aa198">"bootstrap"</span> and <span style="color:#2aa198">"views/home.gohtml"</span> as arguments.
* views.<span style="color:#8f3f71">NewView</span>() takes layout template <span style="color:#2aa198">"bootstrap"</span>, adds to it (yields) <span style="color:#2aa198">"views/home.gohtml"</span> template and other layout template files to it, then parses all the passed and added files, checks for errors and puts returned result to ```var``` <span style="color:#268bd2">homeView</span> in <span style="color:#8f3f71">homeHandler</span>() function.  
<span style="color:#2aa198">"bootstrap.gohtml"</span> file acts as generic scaffold layout that yields given template file:
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
    In the views.<span style="color:#8f3f71">NewView</span>(layout, template) function the layout and the template can be any template file, for eg. with <span style="color:#2aa198">"bootstrap.gohtml"</span> template <span style="color:#2aa198">"home.gohtml"</span> will return Home Page, <span style="color:#2aa198">"contacts.gohtml"</span> will return Contacts Page.
* <span style="color:#268bd2">homeView.Template.</span><span style="color:#8f3f71">ExecuteTemplate</span>(w, homeView.Layout, nil) executes that result by writing response w to the client

