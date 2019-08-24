/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-21 10:50:02
 * @LastEditTime: 2019-08-21 10:52:10
 * @LastEditors: Please set LastEditors
 */
package day03

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
