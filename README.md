# Simple Go Web Server

This Go package provides a simple web server with basic routing functionality. It serves static files from a directory, handles a form submission, and responds to a `/hello` endpoint with a greeting.

## Features

- **Static File Serving**: Serves static files from the `static` directory.
- **Basic Routing**:
  - `/hello`: Returns a simple "Hello!" message.
  - `/form`: Accepts a POST request with form data and responds with a greeting using the submitted name.

## Requirements

- Go 1.16 or later

## Installation

1. Clone this repository:
   ```bash
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```bash
   cd <project-directory>
   ```

## Usage

1. Ensure you have a `static` directory with `index.html` and `form.html`:

   - **`index.html`**: Can be accessed via `http://localhost:8080/` (root URL).
   - **`form.html`**: Can be accessed via `http://localhost:8080/form.html`.

2. Build and run the server:

   ```bash
   go run main.go
   ```

3. The server will start on port `8080`. You can access it in your browser:
   - `http://localhost:8080/` - Serves `index.html`.
   - `http://localhost:8080/form.html` - Serves `form.html` for form submissions.
   - `http://localhost:8080/hello` - Returns "Hello!" for a GET request.

### Example Form Submission

You can navigate to `http://localhost:8080/form.html` to fill the form

OR

You can submit a form using the following `curl` command:

```bash
curl -X POST -d "name=YourName" http://localhost:8080/form
```

This will return:

```bash
POST request successful
Hello YourName!
```

### Code Overview

`main.go`

- `helloHandler`: Handles requests to `/hello`. It only allows `GET` requests and responds with "Hello!". If the URL path or method is incorrect, it returns a `404` error.
- `formHandler`: Handles requests to `/form`. It processes form data from a `POST` request and responds with a personalized greeting using the submitted name. If form parsing fails, an error message is returned.
- `main`: Sets up the server by registering the handlers for `/hello`, `/form`, and serving static files from the `./static` directory. The server listens on port `8080`.
