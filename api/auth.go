package api

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/gochat/models"
)

// GET /login
// 登录页面
func Login(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "login")
}

// GET /signup
// 注册页面
func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "signup")
}

// POST /signup
// 注册新用户
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Errorf("parse form error:%v", err)
	}
	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		log.Errorf("create use error:%v",err)
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// 通过邮箱和密码字段对用户进行认证
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := models.UserByEmail( request.PostFormValue("email"))
	if err != nil {
	    log.Errorf("find user by email error:%v",err)
	}
	if user.Password == models.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Errorf("create session error:%v",err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

// GET /logout
// 用户退出
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
	    log.Errorf("get cookie error:%v",err)
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
