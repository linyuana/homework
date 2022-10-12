package work1010

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	data := []struct {
		val    string
		answer string
	}{
		{val: "abcd", answer: "dcba"},
		{val: "5123", answer: "3215"},
		{val: "我的世界", answer: "界世的我"},
	}

	for _, v := range data {
		s := []byte(v.val)
		reverseString(s[:])
		if string(s) != v.answer {
			t.Fatalf("expect:[%s] != result[%s]", string(s), v.answer)
		}
	}
}

func reverseString(s []byte) {
	str := []rune(string(s))
	i, j := 0, len(str)-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	strbyte := []byte(string(str))
	for i = 0; i < len(s); i++ {
		s[i] = strbyte[i]
	}
}
