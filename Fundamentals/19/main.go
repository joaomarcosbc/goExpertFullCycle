package main

func main() {
	var testVar interface{} = "James"

	println(testVar.(string))

	res, ok := testVar.(int)

	println(res, ok)
}