package main

import (
	"create-packages/myModule1"      // full module name
	mod2 "create-packages/myModule2" // alias module name
)

func main() {
	// calling module1 function
	myModule1.PrintStuff()

	//calling module2 function
	mod2.PrintStuff()

}
