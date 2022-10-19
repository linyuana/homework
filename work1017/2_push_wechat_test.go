package work1017

import (
	"bufio"
	"encoding/json"
	//"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestPushWeChat(t *testing.T) {
	data := []struct {
		title    string
		desp     string
		answer_1 string
		answer_2 string
	}{
		{title: "今天", desp: "今天要早点睡觉呀", answer_1: "SUCCESS", answer_2: "ok"},
		{title: "明天", desp: "明天也要早点睡觉呀", answer_1: "SUCCESS", answer_2: "ok"},
		{title: "后天", desp: "后天更要早点睡觉呀", answer_1: "SUCCESS", answer_2: "ok"},
	}
	for _, v := range data {

		pushid, readkey, error := pushWeChat(v.title, v.desp)
		if pushid == "" || readkey == "" {
			t.Fatalf("返回的pushid或readkey有误!")
		}
		if error != v.answer_1 {
			t.Fatalf("expect:[%v] != result[%v]", error, v.answer_1)
		}

		time.Sleep(time.Duration(5) * time.Second)

		errmsg := checkPush(pushid, readkey)
		if errmsg != v.answer_2 {
			t.Fatalf("expect:[%v] != result[%v]", errmsg, v.answer_2)
		}
	}
}

func pushWeChat(title, desp string) (string, string, string) {
	api := "https://sctapi.ftqq.com/"

	// s := bufio.NewScanner(os.Stdin)
	// fmt.Print("请输入您的密钥：")
	// if !s.Scan() {
	// 	return 
	// }

	//从文件读取sendkey
	f, err := os.Open("sendkey.txt")
	if err != nil {
    		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
    	if err != nil {
    		panic(errs)
    	}
	//发送get请求
	resp, err := http.Get(api + string(line) + ".send?title=" + title + "&desp=" + desp)
	if err != nil {
		panic(err)
	}

	result := Transformation(resp)
	// fmt.Println(result)
	//取出result中的data
	data := result["data"]
	// fmt.Println(data)
	//取出data中的pushid, readkey, error
	temp := data.(map[string]interface{})
	pushid := temp["pushid"]
	readkey := temp["readkey"]
	error := temp["error"]
	// fmt.Println(pushid, readkey, error)
	return pushid.(string), readkey.(string), error.(string)
}

func checkPush(pushid, readkey string) string {
	api := "https://sctapi.ftqq.com/"
	resp, err := http.Get(api + "push?id=" + pushid + "&readkey=" + readkey)
	if err != nil {
		panic(err)
	}
	result := Transformation(resp)
	//取出result中的data
	data := result["data"]
	// fmt.Println(data)
	//取出data中的wxstatus
	temp := data.(map[string]interface{})
	wxstatus := temp["wxstatus"]
	// fmt.Println(wxstatus)
	//取出wxstatus中的errmsg
	wxstatus_map := make(map[string]interface{})
	err = json.Unmarshal([]byte(wxstatus.(string)), &wxstatus_map)
	if err != nil {
		panic(err)
	}
	errmsg := wxstatus_map["errmsg"]
	// fmt.Println(errmsg)
	return errmsg.(string)
}

// transformation 将get请求的返回值转换为键值对
func Transformation(response *http.Response) map[string]interface{} {
	var result map[string]interface{}
	body, err := io.ReadAll(response.Body)
	if err == nil {
		json.Unmarshal([]byte(string(body)), &result)
	}
	return result
}
