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