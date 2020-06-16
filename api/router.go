package api

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/hexiaopi/gochat/config/database"
	"github.com/hexiaopi/gochat/models"
)

func RegisterRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(Logger)

	assets := http.FileServer(http.Dir("public"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	// profiling
	router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)

	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/login", Login).Methods("GET")
	router.HandleFunc("/logout", Logout).Methods("GET")
	router.HandleFunc("/signup", Signup).Methods("GET")
	router.HandleFunc("/signup", SignupAccount).Methods("POST")
	router.HandleFunc("/authenticate", Authenticate).Methods("POST")
	router.HandleFunc("/thread/new", NewThread).Methods("GET")
	router.HandleFunc("/thread/create", CreateThread).Methods("POST")
	router.HandleFunc("/thread/read", ReadThread).Methods("GET")
	router.HandleFunc("/thread/post", PostThread).Methods("POST")
	router.HandleFunc("/err", Err).Methods("GET")
	return router
}

// Checks if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(database.DB); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 生成 HTML 模板
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("layout").Funcs(funcMap)
	templates := template.Must(t.ParseFiles(files...))
	err := templates.ExecuteTemplate(writer, "layout", data)
	if err != nil {
		log.Println(err)
	}
}

// 异常处理统一重定向到错误页面
func errorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// 日期格式化
func formatDate(t time.Time) string {
	datetime := "2006-01-02 15:04:05"
	return t.Format(datetime)
}
