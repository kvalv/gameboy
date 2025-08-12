package gameboy

import (
	"fmt"
	"log/slog"
	"os"
	"slices"
	"testing"
)

func TestBootLoader(t *testing.T) {
	// TODO: LD  ...

	t.Run("binary", func(t *testing.T) {
		cpu := CPUHelper{
			t:   t,
			CPU: &CPU{
				// limit: 100_000,
			}}
		mem := NewMemory().Write(BootROM)
		mem.SetLY()
		if mem.LY() == 0 {
			panic("0")
		}

		visited := make(map[int]int)

		cpu.WithHook(func(cpu *CPU, loc int, instr Instruction, log *slog.Logger) {
			visited[loc]++

			// if loc == 0x00a3 {
			// 	log.Info("write to HL", "HL", hexstr(cpu.HL()))
			// }
			if cpu.Mem.VRAM().HasData() {
				// cpu.err = ErrNoMoreInstructions
				// log.Info("vram has data yay")
			}
			if cpu.PC == 0x0095 {
				log.Info("Graphics routine is called")
			}
			if loc == 0xe6 || cpu.InstrCount > 5000000 {
				cpu.err = ErrNoMoreInstructions
			}
		})
		if err := Run(cpu.CPU, mem, logger()); err != nil && err != ErrNoMoreInstructions {
			t.Fatal(err)
		}
		t.Logf("CPU state after boot: %+v", cpu)
		t.Logf("instruction count: %d", cpu.InstrCount)
		t.Logf("has data: %t", cpu.Mem.VRAM().HasData())
		cpu.Mem.Dump(os.Stderr)
		// h := CPUHelper{CPU: &cpu, t: t}
		// h.DumpMem(0, 100)
		// fmt.Printf("data; \n%s\n", hex.Dump(cpu.mem.VRAM().TileData1))
		// fmt.Printf("data; \n%s\n", hex.Dump(cpu.mem.VRAM().TileData2))

		type elem struct {
			code, count int
		}
		var parts []elem

		for code, count := range visited {
			parts = append(parts, elem{
				code:  code,
				count: count,
			})
		}

		slices.SortFunc(parts, func(a, b elem) int {
			return a.code - b.code
		})
		// cpu.DumpMem(0x800, 0xb000)

		for _, el := range parts {
			fmt.Printf("visited %#4x %d times\n", el.code, el.count)
		}

		t.Fail()

	})

	// t.Run("full boot", func(t *testing.T) {
	// 	cpu := CPU{}
	// 	mem := NewMemory().Write(GetBootCode())
	// 	if err := Run(&cpu, mem, logger(true)); err != nil && err != ErrNoMoreInstructions {
	// 		t.Fatal(err)
	// 	}
	// })

	// t.Run("individual blocks", func(t *testing.T) {
	// 	for i, blk := range GetBootCode() {
	// 		cpu := CPU{}
	// 		blk.Offset = 0
	// 		mem := NewMemory().Write(blk)
	// 		t.Log("Running block", i, "at offset", blk.Offset, "with data", blk.Data)
	// 		if err := Run(&cpu, mem, logger(true)); err != nil && err != ErrNoMoreInstructions {
	// 			t.Errorf("block %d: %v", i, err)
	// 		}
	// 	}
	// })

}

// func TestViewTileset(t *testing.T) {

// 	width, height := 64, 64
// 	img := image.NewRGBA(image.Rectangle{
// 		Min: image.Point{0, 0},
// 		Max: image.Point{width, height},
// 	})

// 	ViewManyTiles(DataLogo[:16], img)
// 	// ViewTile(DataLogo[:16], img, 0, 0)

// 	f, _ := os.Create("/tmp/image.png")
// 	png.Encode(f, img)
// }

// func TestViewTile(t *testing.T) {

// 	width, height := 64, 8
// 	img := image.NewRGBA(image.Rectangle{
// 		Min: image.Point{0, 0},
// 		Max: image.Point{width, height},
// 	})

// 	data := []byte{
// 		0x3C, 0x7E, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x7E, 0x5E, 0x7E, 0x0A, 0x7C, 0x56, 0x38, 0x7C,
// 	}
// 	ViewTile(data, img, 0, 0)

// 	f, _ := os.Create("/tmp/image.png")
// 	png.Encode(f, img)

// }
