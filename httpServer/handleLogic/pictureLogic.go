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

	fmt.Printf("UploadFile [+]")

	rsp := &uploadFileRet{Ret: 0}
	_, status := checkSession(r)
	if !status {
		log.Log.Errorf("check session err")
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

	fmt.Printf("file name:%s   file size:%d  file path:%s\n", handler.Filename, handler.Size, config.Config.PicturePath+handler.Filename)

	file, err := os.OpenFile(config.Config.PicturePath+handler.Filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("UploadFile err:%s", err)
		log.Log.Errorf("open file err:%s", err)
		return
	}

	defer file.Close()

	_, err = io.Copy(file, picture)
	if err != nil {
		fmt.Printf("UploadFile 1 err:%s", err)
		log.Log.Errorf("err:%s", err)
	}
	rsp.Url = config.Config.PictureUrl + handler.Filename
	msg, _ := json.Marshal(rsp)
	_, err = w.Write(msg)
	if err != nil {
		fmt.Printf("UploadFile 2 err:%s", err)
		log.Log.Errorf("err:%s", err)
	}
	fmt.Printf("UploadFile [-]")
}
