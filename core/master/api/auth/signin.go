package authentication

import (
	"log"
	"login/core/database"
	"login/core/models/functions"
	"login/core/sessions"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func Signin(w http.ResponseWriter, r *http.Request) {
    // Parse form data
    err := r.ParseForm()
    if err != nil {
        toast := functions.Toastr{
            Icon:  "error",
            Title: "Error",
            Text:  "All Fields Must Be Valid",
        }
        w.Write([]byte(functions.Toast(toast)))
        http.Error(w, "All Fields Must Be Valid", http.StatusBadRequest)
        http.Redirect(w, r, "/auth/", http.StatusSeeOther)
        return
    }

    username := r.Form.Get("username")
    password := r.Form.Get("password")
    log.Printf("username: %s", username)
    // Authenticate user
    user, err := database.Container.AuthenticateUser(username, password)
    if err != nil {
        if err == database.ErrUserNotFound || err == database.ErrInvalidPassword {
            toast := functions.Toastr{
                Icon:  "error",
                Title: "Error",
                Text:  "Invalid username or password",
            }
            w.Write([]byte(functions.Toast(toast)))
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
            http.Redirect(w, r, "/auth/", http.StatusSeeOther)
            return
        }
        http.Error(w, "Failed to authenticate user", http.StatusInternalServerError)
        return
    }

    // Set up session or token for authenticated user if needed
    sessionToken := uuid.NewString()
    expiresAt := time.Now().Add(30 * time.Minute)
    if _, remember := r.Form["remember-me"]; remember {
        expiresAt = time.Now().Add(48 * time.Hour)
    }

    // Store session in session store
    sessions.Sessions[sessionToken] = sessions.Session{
        User:   user,
        Expiry: expiresAt,
    }

    // Set session token cookie
    cookie := &http.Cookie{
        Name:    "session-token",
        Value:   sessionToken,
        Expires: expiresAt,
        Path:    "/",
    }
    http.SetCookie(w, cookie)
    http.Redirect(w, r, "/dash/", http.StatusSeeOther)
}
