package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//表组测试
func TestWithTable(t *testing.T) {
	inputs := [...]int{2, 3, 4}
	expect := [...]int{4, 9, 16}
	for k, i := range inputs {
		r := Square(i)
		if expect[k] != r {
			t.Errorf("input is %d, the excepted is %d, the actual %d", i, expect[k], r)
		}
	}
}

//断言
func TestWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expect := [...]int{1, 4, 9}
	for k, i := range inputs {
		assert.Equal(t, expect[k], Square(i))
	}
}

//模拟调用
func TestSendJSON(t *testing.T) {
	//创建一个模拟的服务器
	server := mockServer()
	defer server.Close()
	//Get请求发往模拟服务器的地址
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("创建Get失败")
	}
	defer resp.Body.Close()

	log.Println("code:", resp.StatusCode)
	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("body:%s\n", json)
}

func mockServer() *httptest.Server {
	//API调用处理函数
	sendJSON := func(rw http.ResponseWriter, r *http.Request) {
		u := struct {
			Name string
		}{
			Name: "张三",
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(u)
	}
	//适配器转换
	return httptest.NewServer(http.HandlerFunc(sendJSON))
}

//基准测试
func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer() //测试片段开始
	for i := 0; i < b.N; i++ {
		StrAppend(elems...)
	}
	b.StopTimer() //测试片段结束

}

//基准测试
func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StrAppendByBuffer(elems...)
	}
	b.StopTimer()
}
