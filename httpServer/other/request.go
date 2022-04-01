package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type userLoginRsp struct {
	Ret     int32  `json:"ret"`     //业务返回码
	Uid     string `json:"uid"`     //用户ID
	Nick    string `json:"nick"`    //用户昵称
	Picture string `json:"picture"` //用户头像
	Url     string `json:"url"`     //跳转链接
}

func reqUserLogin(uid string, client *http.Client) {
	url2 := "http://101.42.251.239:8080/userLogin"
	data := url.Values{}

	data.Set("uid", uid)
	data.Set("pwd", "12345")

	req, err := http.NewRequest("POST", url2, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	//req.Header.Add("Content-Type", "application/json;charset=UTF8")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println("fei 200")
		return
	}

	var res userLoginRsp
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		fmt.Println(err)
	}
	if res.Ret != 0 {
		fmt.Printf("=========err========:ret:%d    uid:%s\n", res.Ret, uid)
	}

}

func reqWrap(uid string, wg *sync.WaitGroup) {
	defer wg.Done()
	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: 30 * time.Second, Transport: tr}
	for i := 0; i < 1; i++ {
		//time.Sleep(time.Millisecond * 500)
		//fmt.Println("uid:", uid)
		reqUserLogin(uid, client)
	}
	client.CloseIdleConnections()
}

func test1() {

	t1 := time.Now().UnixMilli()

	wg := new(sync.WaitGroup)
	for i := 1; i < 2000; i++ {
		uid := strconv.Itoa(i)
		wg.Add(1)
		go reqWrap(uid, wg)
	}

	wg.Wait()
	fmt.Println("=======total:", time.Now().UnixMilli()-t1)
}

func test2() {

	t1 := time.Now().UnixMilli()
	wg := new(sync.WaitGroup)
	for i := 1; i < 200; i++ {

		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(10000000)
		uid := strconv.Itoa(randomNum)
		wg.Add(1)
		go reqWrap(uid, wg)
	}

	wg.Wait()
	fmt.Println("=======total:", time.Now().UnixMilli()-t1)

}

func main() {
	test1()
	select {}
}
