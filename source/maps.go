/*


Exercise: Maps

Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	table:= make(map[string]int)
	var values []string = strings.Fields(s)

	for _,val:= range values {
		_, ok:= table[val]
		if ok==true {
			table[val] = table[val]+1
		}else {
			table[val]=int(1)
		}
	}
	return table
}

func main() {
	wc.Test(WordCount)
}
