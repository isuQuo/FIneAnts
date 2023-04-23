# Go Web Minimum-Viable-Product (MVP)
This is a template repository for my web applications built with Go.

# Project Structure

This repository contains a Go-based web application. Below is an overview of the directory structure and the purpose of each file:

<pre>
.
├── cmd
│   └── web
│       ├── context.go      # Contains custom context types for passing data between middlewares and handlers
│       ├── funcs.go        # Contains template functions for use in HTML templates
│       ├── handlers.go     # Contains HTTP handlers for handling incoming requests
│       ├── helpers.go      # Contains helper functions for the web application
│       ├── main.go         # Main entry point for the web application
│       ├── middleware.go   # Contains middleware functions for use in the web application
│       ├── routes.go       # Contains route definitions for the web application
│       └── templates.go    # Contains the setup for loading and parsing HTML templates
├── internal
├── pkg
└── ui
    ├── efs.go              # Contains the embedded file system setup for serving static files and templates
    ├── html
    │   ├── layouts
    │   │   └── main.html   # Main layout for the web application
    │   ├── pages
    │   │   └── index.html  # Template for the homepage of the web application
    │   └── partials
    │       └── nav.html    # Template for the navigation bar
    └── static              # Contains static files such as CSS, JavaScript, and images
</pre>

# cmd

This directory contains the entry point for the web application.

## web

The `web` folder contains the main.go file, which is the main entry point of the web application. It initializes and starts the server, configures routes, and sets up any required middleware.

To run the application, navigate to the `web` directory and execute the following command:

```bash
go run main.go
```

This will start the web server and listen for incoming requests. You can access the application by navigating to http://localhost:8080 in your web browser.

Feel free to modify the main.go file to add new routes, middleware, or adjust the server configuration.

# internal
The internal directory contains the application-specific code that is not intended to be reused by other projects. This code is organized into packages based on functionality.

Here you'll find packages related to:

- Database access and manipulation
- Business logic
- Data structures and models
- When adding new functionality to the application, create a new package within this directory and organize your code accordingly. This helps maintain a clean and modular architecture for the web application.

Feel free to use subdirectories within packages to further organize your code if necessary.

# pkg

The pkg directory contains reusable code and libraries that can be imported and used by other projects. This includes utility functions, custom middleware, and other common functionality that could be beneficial to multiple projects.

When creating a new package within this directory, ensure that it is:

- Well-documented
- Contains unit tests
- Has a clear and concise API
- This will make it easier for others to use and understand the purpose of the package.

Remember to use semantic versioning when versioning the packages in this directory, to make it easier for others to manage dependencies in their projects.

# UI (User Interface)

This directory contains the assets and templates required for rendering the user interface of the web application.

## Structure

The UI directory is organized into the following subdirectories:

- `efs.go`: Contains the embedded file system setup for serving static files and templates.
- `html`: Contains the HTML templates for the web application.
  - `layouts`: Contains the layout templates that provide the structure for other templates.
    - `main.html`: The primary layout for the web application, containing the common structure for all pages.
  - `pages`: Contains individual page templates.
    - `index.html`: The template for the homepage of the web application.
  - `partials`: Contains partial templates that can be included in other templates.
    - `nav.html`: The template for the navigation bar.
- `static`: Contains static files such as CSS, JavaScript, and images.

## Usage

When adding a new page to the web application, create a new template in the `html/pages` directory. You can then use the `{{ template "layout" }}` directive in your new template to inherit the structure from the main layout. To include a partial template, use the `{{ template "partial" }}` directive.

For example, to create a new "About" page:

1. Create a new file `html/pages/about.html`.
2. In `about.html`, include the main layout: `{{ template "main" . }}`
3. Add your custom content within the `{{ define "content" }}` and `{{ end }}` directives:
```
{{ define "content" }}
  <h1>About Us</h1>
  <p>Welcome to our About page...</p>
{{ end }}
```
4. Add any required static assets (CSS, JavaScript, images) to the static directory.

Remember to update the application's routing and navigation to include the new page.

Feel free to modify the existing templates, add new templates, or adjust the structure of this directory to suit your project's needs.

# TODO
- [x] Define folder structure
- [x] Define default/boilerplate files/code
- [ ] Document what files and folders are used for