package ui

import (
	"fmt"
	"image"
	"image/color"
	"log/slog"
	"os"
	"strings"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kvalv/gameboy"
)

const CYCLES_TO_RUN = 70000 // should be approx 70k

type Game struct {
	offset int
	cpu    *gameboy.CPU

	input       *Input
	debugui     debugui.DebugUI
	displayVRAM *DisplayVRAM

	debugger       Debugger
	cyclesPerFrame int

	// reference to the screen screen
	screen *Screen
	init   bool
}

func NewGame() *Game {
	cpu := gameboy.CPU{
		Mem: gameboy.NewMemory().Write(gameboy.BootROM),
	}
	cpu.WithLog(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey || a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})))
	cpu.Mem.CursorAt(0x0104)
	cpu.Mem.Write(gameboy.BootLogo)
	cpu.WithHook(func(cpu *gameboy.CPU, loc int, instr gameboy.Instruction, log *slog.Logger) {
		if loc == 0x28 { // after ld a,(de)
			log.Info("reading bit",
				"A", hexstr(cpu.A),
				"DE", hexstr(cpu.DE()),
				"instr", hexstr(instr.Code()),
			)
		}
		if loc == 0x00a3 {
			log.Info("writing to vram?",
				"HL", hexstr(cpu.HL()),
				"A", hexstr(cpu.A),
				"instr", hexstr(instr.Code()),
			)
		}
	})
	game := &Game{
		displayVRAM:    NewDisplayVRAM(cpu.Mem),
		input:          NewInput(),
		cyclesPerFrame: 1,
		screen:         NewScreen(&cpu),
		cpu:            &cpu,
	}

	// "Double up" all the bits of the graphics data
	// game.BreakPointAt(0x0095) // logo to vram routine
	// game.BreakPointAt(0x0055) // logo loaded into vram, start scrolling
	// game.BreakPointAt(0x0055) // scroll logo

	return game
}

var _ ebiten.Game = (*Game)(nil)

// Update implements ebiten.Game.
func (g *Game) Update() error {
	cpu := g.cpu
	g.input.Update()
	// mem := cpu.Mem

	if g.input.KeyQ {
		return ebiten.Termination
	}
	if g.input.KeyN {
		g.debugger.Stepped = true
		fmt.Printf("key n pressed\n")
	}

	_, err := g.debugui.Update(func(ctx *debugui.Context) error {
		ctx.Window("Info", image.Rect(250, 10, 580, 490), func(layout debugui.ContainerLayout) {
			ctx.Header("info", true, func() {
				ctx.SetGridLayout([]int{-2, -1}, nil)

				ctx.Text("Next instruction")
				ctx.Text(nextInstr(cpu))

				ctx.Text("TPS")
				ctx.Text(fmt.Sprintf("%0.2f", ebiten.ActualTPS()))
			})
			ctx.Header("PPU Registers", false, func() {
				ctx.SetGridLayout([]int{-2, -1}, nil)
				ctx.Text("SCY")
				ctx.Text(fmt.Sprintf("%#2x", g.cpu.Mem.SCY()))

				ctx.Text("LY")
				ctx.Text(fmt.Sprintf("%#2x", g.cpu.Mem.LY()))
			})
			ctx.Header("Debugger", false, func() {
				ctx.SetGridLayout([]int{-2, -1}, nil)
				ctx.Checkbox(&g.debugger.Enabled, "Debug enabled").On(func() {
					fmt.Printf("Enabled")
				})
			})

		})
		return nil
	})

	if err != nil {
		return err
	}

	// are we in debug mode?
	if dbg := g.debugger; dbg.Enabled {
		if dbg.BreakPoint > 0 {
			// run until we reach this step
			for cpu.PC != dbg.BreakPoint {
				cpu.Step()
			}
			g.debugger.BreakPoint = 0
			return nil
		}
		if !dbg.Stepped {
			return nil
		}
		n0 := g.cpu.Cycles
		for g.cpu.Cycles-n0 < CYCLES_TO_RUN {
			g.cpu.Step()
		}
		g.debugger.Stepped = false
		return nil
	}

	// otherwise just run regularly...
	if !g.init {
		g.init = true
		// for g.cpu.InstrCount < 1000000 {
		for !g.cpu.Mem.VRAM().HasData() {
			g.cpu.Step()
		}
		fmt.Printf("data is now set\n")
	} else {
		n0 := g.cpu.Cycles
		for g.cpu.Cycles-n0 < CYCLES_TO_RUN {
			g.cpu.Step()
		}
	}

	g.offset++
	return nil
}

func (g *Game) BreakPointAt(loc uint16) *Game {
	g.debugger = Debugger{
		Enabled:    true,
		BreakPoint: loc,
	}
	return g
}

// Draw implements ebiten.Game.
func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA64{0xff, 0x00, 0x00, 0xaa})

	// // lcd image
	lcdImage := ebiten.NewImage(g.screen.Size())
	g.screen.Draw(lcdImage)
	drawRelative(screen, lcdImage, 0.8, 0.5)

	// vram
	vramImage := ebiten.NewImage(g.displayVRAM.Size())
	g.displayVRAM.Draw(vramImage)
	drawRelative(screen, vramImage, 0.01, 0.5)

	// debug stuff 1
	var b strings.Builder
	g.cpu.Dump(&b)
	s := strings.ReplaceAll(b.String(), "DE", "  DE")
	s = strings.ReplaceAll(s, "HL", "  HL")
	ebitenutil.DebugPrint(screen, s)

	// debug stuff 2
	g.debugui.Draw(screen)
}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 900, 600
}

func drawRelative(container, img *ebiten.Image, dx, dy float64) {
	op := &ebiten.DrawImageOptions{}
	W, H := container.Bounds().Dx(), container.Bounds().Dy()
	w, h := img.Bounds().Dx(), img.Bounds().Dy()

	x := float64(W-w) * dx
	y := float64(H-h) * dy
	op.GeoM.Translate(x, y)
	container.DrawImage(img, op)
}
func hexstr[V uint8 | uint16 | int | int8](v V, n ...int) string {
	if len(n) > 0 {
		// if n is provided, use it as the width
		return fmt.Sprintf("%0*x", n[0], v)
	}
	return fmt.Sprintf("%#x", v)
}

// creates a human-readable (assembly-ish) representation of the next instruction
func nextInstr(cpu *gameboy.CPU) string {
	s := cpu.CurrentInstr().String()
	if strings.HasSuffix(s, "n8") {
		n, _ := cpu.Mem.Access(cpu.PC + 1)
		// ADD A,0x04
		return fmt.Sprintf("%s%s", strings.TrimSuffix(s, "n8"), hexstr(n))
	}
	if strings.HasSuffix(s, "n16") {
		n := cpu.Mem.AccessU16(cpu.PC + 1)
		// ADD A,0x04
		return fmt.Sprintf("%s%s", strings.TrimSuffix(s, "n16"), hexstr(n))
	}
	if strings.Contains(s, "a8") {
		n, _ := cpu.Mem.Access(cpu.PC + 1)
		return strings.ReplaceAll(s, "a8", hexstr(n))
	}
	return s
}
