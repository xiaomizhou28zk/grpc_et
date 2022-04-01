package main

import (
	"entryTask/common/log"
	"entryTask/httpServer/common"
	"entryTask/httpServer/config"
	"entryTask/httpServer/handleLogic"
	"net/http"
)

// registerHandler 注册处理方法
func registerHandler() {
	http.HandleFunc("/", handleLogic.LoginPage)
	http.HandleFunc("/userLogin", handleLogic.UserLogin)
	http.HandleFunc("/userLogout", handleLogic.UserLogout)
	http.HandleFunc("/getUserInfo", handleLogic.GetUserInfo)
	http.HandleFunc("/updateUserInfo", handleLogic.UpdateUserInfo)
	http.HandleFunc("/uploadFile", handleLogic.UploadFile)

	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html"))))
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("picture"))))
}

func main() {
	err := log.Init()
	if err != nil {
		return
	}
	log.Log.Infof("server start...")

	err = config.Init()
	if err != nil {
		log.Log.Errorf("load config err:%s", err)
		return
	}

	common.InitPool()

	registerHandler()

	//err = http.ListenAndServeTLS(":8080", "./config/server.crt", "./config/server.key", nil)
	if err != nil {
		log.Log.Errorf("err:%s", err)
	}

	http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Log.Errorf("err:%s", err)
	}
}
