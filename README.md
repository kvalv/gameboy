Starting date ish: July 25

CPU: LR35902. It's based on the Z80, but has fewer instructions and fewer registers.
Has the same registers as Intel 8080.

SP = next unused memory address in stack
PC = next instruction to execute

Read byterange
```
xxd -g1 -s 0x100 -l 0x4f pokemon.gb
```

- Video RAM
- Work RAM, cartridge RAM, Save RAM

CPU <-> RAM <-> Cartridge

CPU access memory in RAM, and if it's within one of the RAM banks, then
access data from cartridge. 


## Ideas
- Debug viewer that loads tile data and tile maps. Lets me see the graphics state.

# Memory
The simplest cartridges contained 32kB of space (0x0000 - 0x7ffff). The entire game fits into it. For largeer games, an MBC (Memory Bank Controller)
swaps out 16kB blocks of RAM from the cartridge.

References:
- https://retrocomputing.stackexchange.com/questions/11732/how-does-the-gameboys-memory-bank-switching-work
- 

The Game Boy includes a small embedded boot ROOM, which can be mapped to 0x0000-0x00FF memory area.
The ROM checks whether the cartridge is valid. If it is, the Boot ROM "unmaps itself" (not sure what that means)
before execution of the cartridge ROM starting at 0x0100.

- 0x0000 - 0x3fff -  16kB - ROM Bank #0
- 0x4000 - 0x7fff -  16kB - Switchable ROM Bank
- 0x8000 - 0xafff -   8kB - Video RAM (VRAM)
- 0xa000 - 0xbfff -   8kB - Switchable RAM Bank
- 0xc000 - 0xdfff -   8kB - Internal RAM
- 0xe000 - 0xfdff - 7.5kB - Echo of Internal RAM
- 0xfe00 - 0xfe9f -  160B - Sprite Attribute Memory  (or Object Attribute Memory, OAM)
- 0xFEA0 - 0xFEFF -   96B - Unused
- 0xFF00 - 0xFF4B -   76B - Special IO   **(yes, very special)**
- 0xFF4C - 0xFF4C -  527B - Unused
- 0xFF80 - 0xFFFE -  127B - Internal RAM
- 0xFFFF - 0xFFFF -    1B - Interrupt Enable Register

Video RAM
1kb = 256 * 4 -> 0x0200
2kb = 256 * 8 -> 0x0400
4kb -> 0x0800

0x800 + 4kB = 0x8000 + 0x0800 -> 0x8800 -> but sub 1 -> 0x87ff

# Display
Screen view 160x144 px (width, height)
Tile view is larger than screen view, so the developer must choose the section to display.
The `SCY` and `SCX` "registers", located at 0xFF42, 0xFF43

Uses scanline rendering: one line at a time, starting from top and moving down. 

Data is put in Video RAM. There are two Tile Views (first and second). 

Each tile is 8x8 pixels, and has a 2-bit color palette (4 variants of gray).
So each tile is 16 bytes in total. Tiles can be moved around in 8x8 steps.

The game boy has two static views, each covering 32x32 tiles (1kB size) - 1 byte each?? 32*8 = 256 bytes. So each tile is 8x8,
so the total area covered is 256px * 256px. The tile view is larger than the display.

- 0x8000 - 0x8fff - 4kB - Tile data 1
- 0x8800 - 0x97ff - 4kB - tile data 2 (note overlap with above)
- 0x9800 - 0x9bff - 1kB - Tile view 1
- 0x9c00 - 0x9fff - 1kB - Tile view 2

So we have two views:
- Background = used for the map
- Window = used for the UI
- (Sprites) = used for moving stuff

If a tile overlap in both Window and Background, it will be occluded by Window. In other words, Window has 100% opacity.

## Sprites
If we only have tiles, the game would be pretty boring and low-res. On top of that, we have sprites.
Relevant to this is the OAM (Object Access Memory)

Sprites can be referenced in either 8x8 pixels or 8x16 pixels. 
They can be moved in 1px increments.
They can be transparent and put on top of both tile views.
OAM does not store sprite data itself, but just stores references to them.

### DMA to OAM
Games can have a sprite table for each scene. To speed up "loading" these sprite tables for new
scenes, we have DMA that writes 168 bytes diretly to OAM. To do this,
write a value `$AB` to register `0xFF46` . Then data `0xAB00 - 0xAB9F` 

### Display registers
- 0xFF40 LCD Display Register -- tile stuff
- 0xFF41 LCD Status (`STAT`) -- no idea
- 0xFF42 - 0xFF4A --- bunch of more positioning / scrolling registers

## Render loop
From top to bottom (repeat 144 times):
1. Load row `LY` of the Background tile view to the line buffer
2. Overwrite the line buffer with row `LY` from the window tile view.
3. Sprite engine generates a 1-pixel section of the sprites, where they intersect `LY` and overwrites the line buffer with this.
4. Increment `LY` 



## Registers
- `0xFF50`: BOOT - Boot ROM lock register. Last bit is disabled by default, and if set then boot is off.

### Graphics
- `FF47`: BG Palette Data (assign gray shades to background tiles)

### Audio
- `0xFF26`: Audio master register

# (Function) calls
Return addresses are pushed to stack using `CALL` and popped using `RET`. The stack
moves downwards, so SP starts at 0xFFFF. When writing an u16 to the stack, the
MSB is at the highest address (i.e. closer to 0xFFFF) and the LSB is at the lowest address.

Function call = push current PC to the stack, then move
PC to the another address. When the call returns, pop from stack
to retrieve previous location.




# End
References: 
- https://gekkio.fi/files/gb-docs/gbctr.pdf
- https://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html
- https://gbdev.io/resources.html#tools

List of Boot ROMs, raw binary
- https://gbdev.gg8.se/files/roms/bootroms/
