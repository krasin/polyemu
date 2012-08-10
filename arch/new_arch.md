Adding New Architecture

1. Find the spec, an emulator, assembler and disassembler. It might be that some of these components don't exist, then skip them.

2. Add tests for all instructions (see arch/dcpu16/emulator_tests.go)

3. Implement an emulator to pass tests (see arch/dcpu16/emulator.go)

