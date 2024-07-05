package authentication

import (
	"log"
	"login/core/database"
	"login/core/models/functions"
	"login/core/models/server"
	"net/http"
)

func init() {
    Route.NewSub(server.NewRoute("/signup", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            // Parse form data
            err := r.ParseForm()
            if err != nil {
                toast := functions.Toastr{
                    Icon:  "error", 
                    Title: "Error",
                    Text:  "Passwords do not match",
                }
                w.Write([]byte(functions.Toast(toast)))
                http.Error(w, "Passwords do not match", http.StatusBadRequest)
                return
            }

            username := r.Form.Get("username")
            password := r.Form.Get("password")
            password2 := r.Form.Get("password2")
            email := r.Form.Get("email")
            log.Printf("username: %s password: %s password2: %s email: %s", username, password, password2, email)
            // Validate password match
            if password != password2 {
                toast := functions.Toastr{
                    Icon:  "error", 
                    Title: "Error",
                    Text:  "Passwords do not match",
                }
                w.Write([]byte(functions.Toast(toast)))
                http.Error(w, "Passwords do not match", http.StatusBadRequest)
                return
            }

            // Create new user
            newUser := &database.User{
                Username: username,
                Password: password,
                Email:    email,
            }

            // Call function to create new user in the database
            err = database.Container.NewUser(newUser)
            if err != nil {
                toast := functions.Toastr{
                    Icon:  "error",
                    Title: "Error",
                    Text:  "Failed to create user",
                }
                w.Write([]byte(functions.Toast(toast)))
                http.Error(w, "Failed to create user", http.StatusInternalServerError)
                return
            }
            // Success message using Toastr
            toast := functions.Toastr{
                Icon:  "success",
                Title: "Success",
                Text:  "User created successfully, please loggin!",
            }
            w.Write([]byte(functions.Toast(toast)))
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }))
}
