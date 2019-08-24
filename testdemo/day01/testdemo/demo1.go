/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-19 19:16:08
 * @LastEditTime: 2019-08-20 17:16:26
 * @LastEditors: Please set LastEditors
 */
package testdemo

const (
	frenchHelloPrefix  = "用Franch进行问好:"
	englishHelloPrefix = "用English进行问好:"
	SpanishHelloPrefix = "用Spanish进行问好:"
	common             = "Hello,"
	English            = "English"
	Franch             = "Franch"
	Spanish            = "Spanish"
)

func Hello(name string, language string) string {
	prefix := ""
	switch language {
	case Franch:
		prefix = frenchHelloPrefix
	case English:
		prefix = englishHelloPrefix
	case Spanish:
		prefix = SpanishHelloPrefix
	default:
		prefix = ""
	}
	return prefix + common + name
}

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
func Sum(nums ...int64) int64 {
	sum := int64(0)
	for _, v := range nums {
		sum += v
	}
	return sum
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
