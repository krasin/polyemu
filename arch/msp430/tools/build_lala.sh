#!/bin/sh

set -eu

msp430-gcc -c -o lala tools/lala.s
msp430-objdump -D lala
