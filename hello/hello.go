
package main

import (
	"os"
	"reflect"
	"strconv"
	"fmt"
	"basic/stringutil"
)

func main() {
	fmt.Println("Hello World")
	fmt.Printf(stringutil.Reverse("!oG, olleH"))
}

func myPrint(args ...interface{}) {
	for _, arg := range args {
		switch v := reflect.ValueOf(arg); v.Kind() {
		case reflect.String:
			os.Stdout.WriteString(v.String())
		case reflect.Int:
			os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))
		}
	}
}
