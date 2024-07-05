# GoLang Login/Signup Web Application with htmx

This project is a simple web application built using GoLang that demonstrates user authentication (login/signup) functionalities. It utilizes htmx for AJAX-like behavior without writing much JavaScript. User passwords are securely encrypted before storing them in the database.

## Features

- **User Authentication**: Allows users to sign up with a username and password, and login with those credentials.
- **Encryption**: User passwords are encrypted using bcrypt before storing them in the database for security.
- **Session Management**: Utilizes HTTP cookies for session management to keep users authenticated between requests.
- **Dynamic UI Updates**: Uses htmx to handle form submissions asynchronously, providing a seamless user experience without full page reloads.

## Technologies Used

- **Go (Golang)**: Backend server language.
- **HTML**: Frontend templates for rendering web pages.
- **CSS**: Styling the frontend.
- **htmx**: For handling AJAX-like behavior directly in HTML.

## Installation

1. Clone the repository:

   ```bash
   git clone (https://github.com/FileGoneIsBack/login-signup-page.git)
   ```
   
2. cd login
3. Install dependencies (assuming you have Go installed)
   ```bash
   go mod tidy
   ```
4. Set up your database configuration in config.go.
5. Build and run the application
   ```bash
   go build & ./login
    ```
## Usage

- Navigate to http://localhost:8080 in your web browser.
- Sign up with a new username and password.
- Log in with your credentials.
- Explore the dashboard or any other features you have implemented.

### Notes:

- **Database**: Ensure you have a database set up and configured. Update the `config.go` file with your database credentials and settings.
- **Security**: Always handle passwords securely, as demonstrated with bcrypt encryption in this project.
