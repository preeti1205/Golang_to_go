/*
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// reader should convert each byte read into 'A'
func (m MyReader) Read(p []byte) (n int, e error){
	for i:=range p{
		p[i] = 'A'
	}

	return len(p), nil
}

func main() {
	//validate() accepts a Reader type
	reader.Validate(MyReader{})
}
