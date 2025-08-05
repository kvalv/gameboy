package gameboy

import (
	"image"
	"image/png"
	"log/slog"
	"os"
	"testing"
)

func TestBootLoader(t *testing.T) {
	// TODO: LD  ...

	t.Run("binary", func(t *testing.T) {
		cpu := CPU{
			limit: 50,
		}
		mem := NewMemory().Write(bootROM)
		cpu.WithHook(func(cpu *CPU, instr Instruction, log *slog.Logger) {
			cpu.Dump(os.Stderr)
		})
		if err := Run(&cpu, mem, logger(true)); err != nil && err != ErrNoMoreInstructions {
			t.Fatal(err)
		}
		t.Logf("CPU state after boot: %+v", cpu)
		t.Logf("instruction count: %d", cpu.instrCount)
		// fmt.Printf("data; \n%s\n", hex.Dump(cpu.mem.VRAM().TileData2))

	})

	t.Run("full boot", func(t *testing.T) {
		cpu := CPU{}
		mem := NewMemory().Write(GetBootCode())
		if err := Run(&cpu, mem, logger(true)); err != nil && err != ErrNoMoreInstructions {
			t.Fatal(err)
		}
	})

	t.Run("individual blocks", func(t *testing.T) {
		for i, blk := range GetBootCode() {
			cpu := CPU{}
			blk.Offset = 0
			mem := NewMemory().Write(blk)
			t.Log("Running block", i, "at offset", blk.Offset, "with data", blk.Data)
			if err := Run(&cpu, mem, logger(true)); err != nil && err != ErrNoMoreInstructions {
				t.Errorf("block %d: %v", i, err)
			}
		}
	})

}

func TestViewTileset(t *testing.T) {

	width, height := 64, 64
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{width, height},
	})

	ViewManyTiles(DataLogo[:16], img)
	// ViewTile(DataLogo[:16], img, 0, 0)

	f, _ := os.Create("/tmp/image.png")
	png.Encode(f, img)
}

func TestViewTile(t *testing.T) {

	width, height := 64, 8
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{width, height},
	})

	data := []byte{
		0x3C, 0x7E, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x7E, 0x5E, 0x7E, 0x0A, 0x7C, 0x56, 0x38, 0x7C,
	}
	ViewTile(data, img, 0, 0)

	f, _ := os.Create("/tmp/image.png")
	png.Encode(f, img)

}
