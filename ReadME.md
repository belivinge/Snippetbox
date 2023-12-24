Book "Let's Go!" by Alex Edwards

** Snippetox **

![Web Visualization](https://github.com/belivinge/Snippetbox/blob/master/ui/static/img/Screenshot%20from%202023-12-07%2006-58-33.png)

```

You’ll learn how to:

 - Setup a project repository which follows the Go conventions.
 - Start a web server and listen for incoming HTTP requests.
 - Route requests to different handlers based on the request path.
 - Send different HTTP responses, headers and status codes to users.
 - Fetch and validate untrusted user input from URL query string parameters.
 - Structure your project in a sensible and scalable way.
 - Render HTML pages and use template inheritance to keep your
 - markup DRY and free of boilerplate.
 - Serve static files like images, CSS and JavaScript from your application.
```

A standard Post-Redirect-Get pattern
   1. The user is shown the blank form when they request GET to /sneep/create
   2. When user completes the form it is submitted to the server via POST request to sneep/create
   3. The form data is validated by the creator handler. If there is any failures, the form will be re-diplayed with the highlighted fields. If it passes, the new snippet will be added to the database and then the user will be redirected to "/sneep/:id".
