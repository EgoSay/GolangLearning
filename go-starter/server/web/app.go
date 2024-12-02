package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-starter/db"
	"go-starter/web/routers"
	"log"
	"net/http"
)

type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
}

func NewApp(d db.DB, cors bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}
	techHandler := app.GetTechnologies
	if !cors {
		techHandler = disableCors(techHandler)
	}
	app.handlers["/apis/technologies"] = techHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	//初始化
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	//基础配置
	// InitGin(r)

	//注册路由
	routers.NewInner().Register(r)
	routers.NewAdmin().Register(r)

	//啟動服務
	return http.ListenAndServe("0.0.0.0:9900", r)
}

func (a *App) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
