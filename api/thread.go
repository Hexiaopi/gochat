package api

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/gochat/models"
)

// GET /threads/new
// 创建群组页面
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "auth.navbar", "new.thread")
	}
}

// POST /thread/create
// 执行群组创建逻辑
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			log.Errorf("parse form error:%v", err)
		}
		user, err := sess.User()
		if err != nil {
			log.Errorf("get user from session error:%v", err)
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread( topic); err != nil {
			log.Errorf("create thread error:%v", err)
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// 通过 ID 渲染指定群组页面
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID( uuid)
	if err != nil {
		errorMessage(writer, request, "该群组不存在，无法获取")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(writer, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
