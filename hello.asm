; hello.asm
; Minimal Hello World for DMG/CGB

; ==== Header ====
SECTION "Header", ROM0[$100]
    nop
    jp  Start           ; entry point at 0x0100

    ; Nintendo logo + metadata (rgbfix will patch this later)
    ds $150 - @, 0      ; fill up until 0x0150 with zeros

; ==== Code ====
SECTION "Code", ROM0[$150]

Start:
    ; disable interrupts
    di
    ld sp, $FFFE        ; init stack pointer

    ; Turn off LCD before messing with VRAM
    ld a, 0
    ld [$FF40], a

    ; Load "Hello World!" into tile data
    ld hl, HelloTiles
    ld de, $9000        ; start of tile data
    ld bc, HelloTilesEnd - HelloTiles
CopyTiles:
    ld a, [hl+]
    ld [de], a
    inc de
    dec bc
    ld a, b
    or c
    jr nz, CopyTiles

    ; Load tilemap
    ld hl, HelloMap
    ld de, $9800        ; BG Map 0
    ld bc, HelloMapEnd - HelloMap
CopyMap:
    ld a, [hl+]
    ld [de], a
    inc de
    dec bc
    ld a, b
    or c
    jr nz, CopyMap

    ; Turn LCD back on (BG enabled)
    ld a, %10000001
    ld [$FF40], a

    ld a, 1

Forever:
    ; inc a
    ; ld [$ff01], a
    ; ld a, $81
    ; ldh [$ff02], a
    ; ld a, $80
    ; ldh [$ff02], a
    jr Forever

; ==== Data ====
SECTION "Tiles", ROM0
HelloTiles:
    INCBIN "font.chr"   ; 1bpp font tiles (you supply this, 8x8 chars)
HelloTilesEnd:

SECTION "Tilemap", ROM0
HelloMap:
    db "HELLO WORLD!"   ; each byte = tile index (depends on your font)
HelloMapEnd:

