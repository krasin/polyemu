#!/bin/bash

# Create a stub msp430 ELF file to load it into mspdebug sim.
# Usage:
# ./tools/build_stub.sh

set -ue

msp430-gcc -o tools/stub.elf tools/stub.c
