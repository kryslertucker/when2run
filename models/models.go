package models

import (
	"database/sql"
	"html/template"
	"time"

	"github.com/gorilla/sessions"
)

const (
	minDefaultTemperature float64 = 0
	maxDefaultTemperature float64 = 110
)

type (

	// Controller interface has a Dispatcher
	Controller interface {
		Dispatcher()
	}

	//Server struct
	Server struct {
		Sess *sessions.CookieStore
		Tpl  *template.Template
		Db   *sql.DB
	}

	//User type
	User struct {
		Email    string
		Password string
		Settings Settings
	}

	//Settings type
	Settings struct {
		MinTemp float64
		MaxTemp float64
	}

	//ResponseMain ...
	responseMain struct {
		Temp    float64 `json:"temp"`
		TempMin float64 `json:"temp_min"`
		TempMax float64 `json:"temp_max"`
	}

	//ResponseElem ...
	responseElem struct {
		Ts          int          `json:"dt"`
		TempValues  responseMain `json:"main"`
		TsFormatted time.Time
		GoRun       bool
	}

	// Responsetype ...
	Responsetype struct {
		List []responseElem `json:"list"`
	}

	// TplData ...
	TplData struct {
		Title     string
		TabActive string
		Data      Responsetype
	}

	// App ...
	App struct {
		User        User
		UserLogged  bool
		CurrentView string
		Data        interface{}
		MsgError    string
	}
)

//New creates new settings with defaults temperatures
func (r *Settings) New() Settings {
	return Settings{
		MinTemp: minDefaultTemperature,
		MaxTemp: maxDefaultTemperature,
	}
}

func (u User) isEmpty() bool {
	if u.Email == "" && u.Password == "" {
		return true
	}
	return false
}