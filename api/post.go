package api

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/gochat/models"
)

// POST /thread/post
// 在指定群组下创建新主题
func PostThread(writer http.ResponseWriter, request *http.Request) {
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
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := models.ThreadByUUID( uuid)
		if err != nil {
			errorMessage(writer, request, "该群组不存在，无法获取")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			log.Errorf("create post error", err)
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
