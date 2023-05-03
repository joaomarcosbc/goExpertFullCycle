package main

const a = "Hello World"
var aplus string

var (
	b bool 		= true // False by default
	c int  		= 10 // 0 by default
	d string 	= "Joao" // Empty by default
	e float64 	= 1.2 // +0.000000e+000 by default
)
 
func main() {
	var funcScope string = "X" // when declared on local scope, must be used
	shortCut := "Can be declared like that"
	shortCut = "Revalue with only ="

	println(funcScope)
	println(shortCut)

}

func x() {
	// can not reach funcScope variable
}
