/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-22 08:41:43
 * @LastEditTime: 2019-08-22 08:45:17
 * @LastEditors: Please set LastEditors
 */
package main

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
func main() {

}
