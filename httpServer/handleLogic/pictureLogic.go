package handleLogic

import (
	"encoding/json"
	"entryTask/common/log"
	"entryTask/httpServer/common"
	"entryTask/httpServer/config"
	"fmt"
	"io"
	"net/http"
	"os"
)

type uploadFileRet struct {
	Ret int32  `json:"ret"` //业务返回码
	Url string `json:"url"` //文件地址
}

// UploadFile 上传文件
func UploadFile(w http.ResponseWriter, r *http.Request) {

	rsp := &uploadFileRet{Ret: 0}
	_, status := checkSession(r)
	if !status {
		fmt.Println("get session err")
		rsp.Ret = common.InvalidSession
		rsp.Url = config.Config.LoginPage
		msg, _ := json.Marshal(rsp)
		_, err := w.Write(msg)
		if err != nil {
			log.Log.Errorf("err:%s", err)
		}
		return
	}

	picture, handler, err := r.FormFile("picture")
	if err != nil {
		return
	}

	file, err := os.OpenFile(config.Config.PicturePath+handler.Filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("open file")
		return
	}

	defer file.Close()

	_, err = io.Copy(file, picture)
	if err != nil {
		log.Log.Errorf("err:%s", err)
	}
	rsp.Url = config.Config.PictureUrl + handler.Filename
	msg, _ := json.Marshal(rsp)
	_, err = w.Write(msg)
	if err != nil {
		log.Log.Errorf("err:%s", err)
	}
}
