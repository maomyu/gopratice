/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-11 17:08:21
 * @LastEditTime: 2019-08-11 17:19:43
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 发送Get请求
func GetRequest() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	GetRequest()
}
