package controller

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// return 404 not found page
		tmpl, err := template.ParseFiles(
			path.Join("views/pages", "404.html"), // halaman yang ingin ditampilkan
			path.Join("views", "layout.html"),
			path.Join("views/includes", "sidebar.html"),
			path.Join("views/includes", "navbar.html"),
			path.Join("views/includes", "footer.html"),
			path.Join("views/includes", "scripts.html"),
		)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}

	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "dashboard.html"), // halaman yang ingin ditampilkan
		path.Join("views", "layout.html"),
		path.Join("views/includes", "sidebar.html"),
		path.Join("views/includes", "navbar.html"),
		path.Join("views/includes", "footer.html"),
		path.Join("views/includes", "scripts.html"),
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

func Buttons(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "buttons.html"), // halaman yang ingin ditampilkan
		path.Join("views", "layout.html"),
		path.Join("views/includes", "sidebar.html"),
		path.Join("views/includes", "navbar.html"),
		path.Join("views/includes", "footer.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

func Cards(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "cards.html"), // halaman yang ingin ditampilkan
		path.Join("views", "layout.html"),
		path.Join("views/includes", "sidebar.html"),
		path.Join("views/includes", "navbar.html"),
		path.Join("views/includes", "footer.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

/**
list of pages
1. Login
2. Register
3. Forgot Password
4. 404 Page
5. Blank Page
*/
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "login.html"), // halaman yang ingin ditampilkan
		path.Join("views", "pages.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "register.html"), // halaman yang ingin ditampilkan
		path.Join("views", "pages.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "forgot-password.html"), // halaman yang ingin ditampilkan
		path.Join("views", "pages.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

func Blank(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "blank.html"), // halaman yang ingin ditampilkan
		path.Join("views", "layout.html"),
		path.Join("views/includes", "sidebar.html"),
		path.Join("views/includes", "navbar.html"),
		path.Join("views/includes", "footer.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

func Charts(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "charts.html"), // halaman yang ingin ditampilkan
		path.Join("views", "layout.html"),
		path.Join("views/includes", "sidebar.html"),
		path.Join("views/includes", "navbar.html"),
		path.Join("views/includes", "footer.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

func Tables(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("views/pages", "tables.html"), // halaman yang ingin ditampilkan
		path.Join("views", "layout.html"),
		path.Join("views/includes", "sidebar.html"),
		path.Join("views/includes", "navbar.html"),
		path.Join("views/includes", "footer.html"),
		path.Join("views/includes", "scripts.html"),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

/**
1. Colors
2. Borders
3. Animations
4. Other
*/
