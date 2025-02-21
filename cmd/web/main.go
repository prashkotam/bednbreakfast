package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/prashkotam/bednbreakfast/internal/config"
	"github.com/prashkotam/bednbreakfast/internal/handlers"
	"github.com/prashkotam/bednbreakfast/internal/render"
)

const portNumber = ":8080"

var App config.Appconfig

var session *scs.SessionManager

func main() {

	App.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = App.InProduction

	App.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create Template cache")
	}

	App.TemplateCache = tc
	App.UseCache = false

	repo := handlers.NewRepo(&App)
	handlers.NewHandler(repo)
	render.NewTemplate(&App)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port no %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&App),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
