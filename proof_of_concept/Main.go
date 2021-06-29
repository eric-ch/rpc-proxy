/*
 Go implementation of RPC-Proxy for Openxt
*/

package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	rules := tempReadConfig("rpc-proxy1.rules")
	fmt.Println(rules)
	bouncer()
}
