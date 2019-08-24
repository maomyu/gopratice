/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-21 10:50:31
 * @LastEditTime: 2019-08-21 10:52:50
 * @LastEditors: Please set LastEditors
 */
package day03

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
