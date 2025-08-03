Starting date ish: July 25

CPU: LR35902. It's based on the Z80, but has fewer instructions and fewer registers.
Has the same registers as Intel 8080.

SP = next unused memory address in stack
PC = next instruction to execute


# Memory
The simplest cartridges contained 32kB of space (0x0000 - 0x7ffff). The entire
game fits into the. 
Memory Bank Controller (MBC)

References:
- https://retrocomputing.stackexchange.com/questions/11732/how-does-the-gameboys-memory-bank-switching-work
- 

# (Function) calls
Return addresses are pushed to stack using `CALL` and popped using `RET`. The stack
moves downwards, so SP starts at 0xFFFF. When writing an u16 to the stack, the
MSB is at the highest address (i.e. closer to 0xFFFF) and the LSB is at the lowest address.

Function call = push current PC to the stack, then move
PC to the another address. When the call returns, pop from stack
to retrieve previous location.


# Display
Screen view 160x144 px
Tile view is larger than screen view, so the developer must choose the section to display, using `SCY` and `SCX` "registers", located at 0xFF42, 0xFF43


# End
References: 
- https://gekkio.fi/files/gb-docs/gbctr.pdf
- https://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html

