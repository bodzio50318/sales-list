package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"

	"github/bodzio50318/saleslist/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ApiServer struct {
	listenAddress string
	store         storage.Storage
}
type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
}

func (s *ApiServer) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = newTemplate()

	e.GET("/item", s.handleGetItems)
	e.POST("/item", s.handlePostItems)

	log.Println("Starting a server on port: ", s.listenAddress)
	e.Logger.Fatal(e.Start(s.listenAddress))
}

func NewApiServer(listenAddress string, store storage.Storage) *ApiServer {
	return &ApiServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
