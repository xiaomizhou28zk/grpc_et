package main

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"entryTask/tcpServer/Dao"
	"fmt"
	"strconv"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	/*
		c, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer c.Close()
		if _, err := c.Do("AUTH", "fc4234450eac470d8604fbdff4a34121"); err != nil {
			c.Close()
			fmt.Println(err)
			return
		}
	*/

	p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mmp := make(map[int]int)

	for i := 0; i < len(p); i++ {
		mmp[p[i]] = p[i+1]
		i++
	}

	fmt.Println(mmp)
}

var xc *xorm.Engine

type scott struct {
	P int
}

const (
	politicsModel = "political_flag_and_signs:videoqc"
	fireModel     = "violent_behaviors_m_fire:videoqc"
)

func judgeResult(result bool, safetyViolation string, labellingViolation string) (bool, string, string) {

	hitPoliticsModel := strings.Contains(labellingViolation, politicsModel)
	hitFireModel := strings.Contains(safetyViolation, fireModel)

	//safety没有命中点火,机审结果不变，violation不变
	if !hitFireModel {
		return result, safetyViolation, labellingViolation
	}

	//点火violation文案不单独展示，需要去除
	newSafetyViolation := rearrangeViolation(safetyViolation, fireModel)

	//命中点火，未命中涉政，机审结果改为true
	if !hitPoliticsModel {
		return true, newSafetyViolation, labellingViolation
	}

	//点火与涉政同时命中,点火跟涉政文案同时被去除，并在safety添加hateful文案
	newLabellingViolation := rearrangeViolation(labellingViolation, politicsModel)
	newSafetyViolation = fmt.Sprintf("%s#hateful_content_m_fire_politics:videoqc", newSafetyViolation)

	return false, newSafetyViolation, newLabellingViolation
}

func rearrangeViolation(src string, sep string) string {
	uniqueMap := make(map[string]struct{})
	newViolation := ""
	for _, elem := range strings.Split(src, "#") {
		if elem == "" {
			continue
		}
		for _, item := range strings.Split(elem, ";") {
			if item == "" {
				continue
			}
			if item == sep {
				continue
			}
			if _, ok := uniqueMap[item]; ok {
				continue
			}
			uniqueMap[item] = struct{}{}
			newViolation = fmt.Sprintf("%s#%s", newViolation, item)
		}
	}
	return newViolation
}

func getTaskDetails() {
	//tasks := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	//tasks := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//tasks := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	tasks := make([]uint64, 0)

	ch := make(chan uint64)

	newTask := make([]uint64, 0)

	getTask := func(t []uint64, wg *sync.WaitGroup) {
		defer wg.Done()

		for _, item := range t {
			ch <- item
		}
	}

	//分批获取，防止一次包体过大
	go func() {
		wg := new(sync.WaitGroup)
		start := 0
		end := 0
		for range tasks {
			end++
			if (end)%20 == 0 {
				wg.Add(1)
				go getTask(tasks[start:end], wg)
				start = end
			}
		}

		if start < end {
			wg.Add(1)
			go getTask(tasks[start:end], wg)
		}

		wg.Wait()
		close(ch)
	}()

	for task := range ch {
		newTask = append(newTask, task)
	}

	fmt.Printf("test:%v\n", newTask)
}

func test11() {
	result := false
	label := "#political_flag_and_signs:videoqc;political_flag_and_signs:videoqc#political_flag_and_signs:videoqc#other"
	safety := "#violent_behaviors_m_fire:videoqc"

	a, b, c := judgeResult(result, safety, label)

	fmt.Printf("r:%t   s:%s   l:%s\na:%t   b:%s   c:%s\n", result, safety, label, a, b, c)
}

func test12() {
	result := false
	label := "#political_flag_and_signs:videoqc;political_flag_and_signs:videoqc#political_flag_and_signs:videoqc#other"
	safety := "#oothee"

	a, b, c := judgeResult(result, safety, label)

	fmt.Printf("r:%t   s:%s   l:%s\na:%t   b:%s   c:%s\n", result, safety, label, a, b, c)
}

func test13() {
	result := false
	label := "#other000000"
	safety := "#violent_behaviors_m_fire:videoqc"

	a, b, c := judgeResult(result, safety, label)

	fmt.Printf("r:%t   s:%s   l:%s\na:%t   b:%s   c:%s\n", result, safety, label, a, b, c)
}

func test14() {
	result := true
	label := ""
	safety := ""

	a, b, c := judgeResult(result, safety, label)

	fmt.Printf("r:%t   s:%s   l:%s\na:%t   b:%s   c:%s\n", result, safety, label, a, b, c)
}

type Node struct {
	P1 string `json:"p_1"`
	P2 string `json:"p_2"`
}

type Node2 struct {
	P1 string `json:"p_1"`
	P2 string `json:"p_2"`
	P3 string `json:"p_3"`
}

type test struct {
	Q1 uint32 `json:"q_1"`
}

type test2 struct {
	A string `json:"a"`
	B uint32 `json:"b"`
}

type N1 struct {
	p uint64
}

type N2 struct {
	p1 *N1
	p2 uint64
}

/*
func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func test1111() {
	reader := ExampleReadFrameAsJpeg("./sample_data/in1.mp4", 5)
	img, err := imaging.Decode(reader)
	if err != nil {
		fmt.Println(err)
	}
	err = imaging.Save(img, "./sample_data/out1.jpeg")
	if err != nil {
		fmt.Println(err)
	}
}
*/

func main2222() {
	//file := "https://safety-pms.sv.test.shopee.co.id/api/v1/policy_management/download_file?file_key=12345.png"

	//s, _ := imgtype.Get(file)
	//fmt.Println(s)

	ppp := "AbbbNNNN123123"
	fmt.Println(strings.ToLower(ppp))

	/*
		a1 := []string{"111", "2220000000", "3333333"}

		a2 := []string{"222", "333", "111"}

		sort.SliceStable(a1, func(i, j int) bool {
			return a1[i] > a1[j]
		})
		sort.SliceStable(a2, func(i, j int) bool {
			return a2[i] > a2[j]
		})

		fmt.Println(a1)
		fmt.Println(a2)

		jj := &test{}

		err := json.Unmarshal([]byte(""), jj)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(jj.Q1)
		}


				t1 := &test{
					Q1: 111,
				}

				t2 := &test{
					Q1: 111,
				}

				tt1, _ := json.Marshal(t1)
				tt2, _ := json.Marshal(t2)

				if string(tt1) != string(tt2) {
					fmt.Println("diff")
				} else {
					fmt.Println("not diff")
				}

				str := `1231231231
			123123123123123
			12312312`

				fmt.Println(str)


					g := ""

					p, err := json.Marshal(&g)
					if err != nil {
						fmt.Println(err)
						return
					}

					fmt.Println("p:", string(p))

					var g1 string
					err = json.Unmarshal(p, &g1)
					if err != nil {
						fmt.Println(err)
						return
					}

					fmt.Println("g1:", g1)

					if len(g1) == 0 {
						fmt.Println("=====")
					}

						p := make([]*Node, 0)
						p = append(p, &Node{"1", "2"})

						ps, err := json.Marshal(&p)
						if err != nil {
							fmt.Printf("err:%s", err)
							return
						}

						println("test:", ps)

						var p2 []*Node

						err = json.Unmarshal(ps, &p2)
						if err != nil {
							fmt.Printf("err:%s", err)
							return
						}

						println("suc:", p2[0].P2)

							a := "&"

							arr := strings.Split(a, "&")

							fmt.Println(arr)


								t := time.Now()
								newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
								fmt.Println(newTime.UnixMilli())

								li := []int{1, 2, 3, 4, 5}

								fmt.Println(li[0:3])

	*/

	/*
		n1 := &Node2{"1", "2", "3"}

		nstr, err := json.Marshal(n1)
		if err != nil {
			fmt.Printf("err:%s", err)
			return
		}
		fmt.Println(string(nstr))

		n2 := &Node{}
		err = json.Unmarshal(nstr, n2)

		if err != nil {
			fmt.Printf("err:%s", err)
			return
		}
		fmt.Printf("hahah %s", n2.P1)
	*/
	//test14()

	//var mp map[string]int32

	//fmt.Printf("test:%d\n", mp["test"])

	//getTaskDetails()

	/*
		mp := make(map[uint64][]*int)

		for i := uint64(0); i < 10; i++ {
			p := 10
			mp[i] = append(mp[i], &p)
		}

		fmt.Println("===:", mp)

	*/

	/*
		a := uint64(time.Now().UnixNano() / int64(time.Millisecond))

		fmt.Println("a:", a)

		arr := strings.Split("", ";")

		fmt.Println("arr:", arr)

	*/
	//p := uint8(50)

	//p = ^p
	//p = p | (uint32(1) << 1)
	//fmt.Println(p)

	//test33()
	//test1()
	//now := time.Now()
	//s := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location()).String()
	//fmt.Println("s:", s)
}

type UserInfo struct {
	ID       int64  `xorm:"id"`
	Uid      string `xorm:"uid"`
	Nick     string `xorm:"nick"`
	Picture  string `xorm:"picture"`
	Password string `xorm:"password"`
}

func test33() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "123456789",
		"127.0.0.1", 3306, "user")
	var err error
	xc, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("1err:", err)
		return
	}
	u := UserInfo{Uid: "11111122", Nick: "test1222", Picture: "333", Password: "0000"}

	_, err = xc.Context(context.Background()).Insert(u)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("=====ok======")
}

func test1() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "123456789",
		"127.0.0.1", 3306, "user")
	var err error
	xc, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("1err:", err)
		return
	}

	//u := &Dao.UserInfo{}

	us := make([]*Dao.UserInfo, 0)
	/*

		err = xc.Table("user_info").Context(context.Background()).Where("nick=?", "kk2kAAA").Find(&us)
		if err != nil {
			fmt.Println("2err:", err)
			return
		}
	*/

	se := xc.Table("user_info").Context(context.Background()).Where("id=?", 1)

	se.Where("nick like ?", "%kk2kAA%")

	//se.Or("uid=?", "111")
	//se.And("uid like ?", "%1%")

	err = se.Find(&us)

	if err != nil {
		fmt.Println("this is a err:", err)
		return
	}

	up := map[string]int{"1": 1, "2": 2}

	fmt.Println("map:", up)

	if len(us) == 0 {
		fmt.Println("no data")
		return
	}

	for _, elem := range us {
		fmt.Println("suc:   ", *elem)
	}

}

var db *sql.DB

func BKDRHash(str string) uint64 {
	seed := uint64(131) // 31 131 1313 13131 131313 etc..
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint64(str[i])
	}
	return hash & 0x7FFFFFFF
}

func InitDB() (err error) {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/user"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("db err:", err)
		return err
	}
	return nil
}

func insertData(sqlStr string) error {

	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("update userinfo err:", err)
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get RowsAffected err", err)
		return err
	}

	fmt.Println("insert n:", n)
	return nil
}

func encryptedByMD5(s string) string {
	h := md5.New()
	//uid不会发生变化
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func mai2n() {

	mmp := make(map[uint64][]int)

	for i := 1; i < 10000001; i++ {
		index := BKDRHash(strconv.Itoa(i)) % 100
		mmp[index] = append(mmp[index], i)
	}

	InitDB()

	//k 表名索引  v uid列表
	for k, v := range mmp {
		execstring := fmt.Sprintf("insert into user_info_%d(uid, nick, picture, password) values", k)
		fmt.Println("insert :", execstring)

		data := ""
		count := 0
		for _, elem := range v {
			pwd := encryptedByMD5(fmt.Sprintf("%d12345", elem))
			data += fmt.Sprintf("(\"%d\",\"%d\",\"https://127.0.0.1:8080/picture/111.png\", \"%s\")", elem, elem, pwd)

			count++

			if count%1000 == 0 {
				insertData(execstring + data)
				data = ""
				continue
			}

			if count == len(v) {
				break
			}

			data += ","
		}
		insertData(execstring + data)

	}
}

func main4() {
	fmt.Println("index:", BKDRHash("99961")%100)
}

func main44() {
	mai2n()
}

func main2() {

	mmp := make(map[uint64][]int)

	for i := 1; i < 10000001; i++ {
		index := BKDRHash(strconv.Itoa(i)) % 100
		if i == 99964 {
			fmt.Printf("index:%d  ", index)
		}
		mmp[index] = append(mmp[index], i)
	}

	li := mmp[51]

	for _, elem := range li {
		if elem == 99964 {
			fmt.Println("find")
		}
	}

	//k 表名索引  v uid列表
	for _, v := range mmp {

		//fmt.Printf("k:%d  len:%d\n", k, len(v))

		count := 0
		for range v {

			count++
			if count%1000 == 0 {
				//fmt.Println("ix:", ix+1)
				//插入
				continue
			}

			if count == len(v) {
				//fmt.Println("====")
			}
		}
		if count == 100000 {
			//fmt.Println("haha", count)
		}
	}
}
