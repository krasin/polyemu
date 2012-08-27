#!/bin/sh

# Compiles asm snippet and disassembles it.
# Useful for preparing binary tests.
# Usage: ./tools/build_snippet.sh
set -eu

msp430-gcc -c -o asm-snippet tools/asm-snippet.s
msp430-objdump -D asm-snippet
