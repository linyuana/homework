package work1017

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestDogDiary(t *testing.T) {
	result, err := dogDiary()
	if err != nil {
		t.Fatal(err)
	}
	if result == "" {
		t.Fatalf("舔狗日记的result不能为空!")
	}
}

// DogDiary 获取舔狗日记
func dogDiary() (string, error) {
	const api = "https://api.ixiaowai.cn/tgrj/index.php"
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
