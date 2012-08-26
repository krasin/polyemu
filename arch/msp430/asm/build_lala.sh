#!/bin/sh

set -eu

msp430-gcc -c -o lala asm/lala.s
msp430-objdump -D lala
