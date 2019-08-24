/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-17 16:44:05
 * @LastEditTime: 2019-08-17 17:08:41
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type LendBook struct {
	UserID     string    `json:"userID"`
	Bookcode   string    `json:"bookcode"`
	Lenddate   time.Time `json:"lenddate"`
	Returndate time.Time `json:"returndate"`
	Status     int       `json:"status"`
	Bookisbn   string    `json:"bookisbn"`
}

func main() {
	db, _ := gorm.Open("mysql", "root:micomysql@tcp(192.168.10.252:3306)/user?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.CreateTable(&LendBook{})

}
