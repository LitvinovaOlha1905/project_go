package main

import (
	"examples/headFirst/src/greeting"
	"examples/headFirst/src/greeting/deutsch"
	"github.com/headfirstgo/greeting/dansk"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.Hallo()
	deutsch.GutenTag()
	dansk.GodMorgen()
}
