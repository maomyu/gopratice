/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-19 19:16:27
 * @LastEditTime: 2019-08-21 10:10:19
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"testing"

	"github.com/yuwe1/gopratice/testdemo/day01/testdemo"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := testdemo.Hello("mowuya", "")
		want := "Hello,mowuya"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := testdemo.Hello("", "")
		want := "Hello,world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := testdemo.Hello("Elodie", "Spanish")
		want := "Hola,Elodie"
		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

// 基准测试
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testdemo.Repeat("a")
	}
}

// 测试数组
func TestSum(t *testing.T) {
	t.Run("test five num", func(t *testing.T) {
		got := testdemo.Sum(1, 2, 3, 4, 5)
		want := int64(15)
		if want != got {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestPerimeter(t *testing.T) {
	rectangle := testdemo.Rectangle{10.0, 10.0}
	got := testdemo.Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
