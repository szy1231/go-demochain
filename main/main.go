/*
@Time : 2020/3/3 13:05
@Author : pasiyu
@File : main.go
*/
package main

import "demochain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 1 BTC to Jacky")
	bc.Print()
}