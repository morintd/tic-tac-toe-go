package main

import (
	controllerpkg "tic-tac-toe/controller"
	uipkg "tic-tac-toe/ui"
)

func main() {
	controller := controllerpkg.NewController(uipkg.ConsoleUI{})
	controller.Start()
}
