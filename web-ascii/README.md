# ASCII-Art Web Service

A web service written in Go that transforms text into graphic ASCII banners. It features a web interface, custom error handling, and a clean modular structure.

## Features
* **Three Banner Styles:** `standard`, `shadow`, and `thinkertoy`.
* **Web Interface:** Easy-to-use form to input text and select styles.
* **Error Handling:** Custom `error.html` page for 404, 405, and 500 errors.
* **Modular Code:** Separated into `main.go`, `handlers.go`, and `ascii.go` for better organization.

## Project Structure

```text
ascii-art-stylize/
├── main.go          # Server startup and routes
├── handlers.go      # Request handling logic and HTML rendering
├── ascii.go         # ASCII generation logic
├── banners/         # Font files (.txt)
├── templates/       # HTML templates (index.html, error.html)
└── static/          # CSS styles
```
## Usage

Clone the repository and install dependencies:
```bash
git clone https://01.tomorrow-school.ai/git/akuncher/ascii-art-web-stylize.git
cd ascii-art-web-stylize
```


### Run the server:
```bash
go run .
```

### Open your browser and go to:
```text
http://localhost:8080
```


* Enter your text on the main page and choose a font.

* Click Generate to see the ASCII art output.