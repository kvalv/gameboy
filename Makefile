hello:
	rgbasm -o main.o hello.asm 
	rgblink -o game.gb main.o 
	rgbfix -t hello -v -p 0 game.gb
