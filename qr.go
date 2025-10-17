package main

import (
	"os"
	"github.com/mdp/qrterminal/v3"
)

func main() {
	//-------------------
	// Temp print QR of compressed key on screen 
	qrterminal.Generate("", qrterminal.L, os.Stdout)
	// !! delete key!
	// ctr-b :clear-history in tmux 
	// reset in termimal

}
