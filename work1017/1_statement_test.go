package work1017

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestStatement(t *testing.T) {

	result, err := statement()
	if err != nil {
		t.Fatal(err)
	}
	if result == "" {
		t.Fatalf("一言语录的result不能为空！")
	}

}

// Statement 获取一言语录
func statement() (string, error) {
	const api = "https://api.ixiaowai.cn/ylapi/index.php"
	resp, err := http.Get(api)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf(resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return string(b), err
}
