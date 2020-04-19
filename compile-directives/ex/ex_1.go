package ex

/*
  执行 GOOS=linux GOARCH=386 go tool compile -S ex_1.go > ex_1.S
*/
func appendStr(word string) string {
	return "new " + word
}

//go:noinline
func appendStr2(word string) string {
	return "new " + word
}

func test() {
	appendStr("11")
	appendStr2("22")
}
