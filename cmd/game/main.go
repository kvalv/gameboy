package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvalv/gameboy/ui"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	g := ui.NewGame()

	// ebiten.SetWindowSize(200, 200)
	ebiten.SetWindowTitle("Game Boy")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
