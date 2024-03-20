# Paintf\*\*k

Paintfuck is an esoteric programming language that operates on an infinite two-dimensional grid of cells. Each cell can hold an integer value and has a color associated with it. The language provides commands for navigating the grid, changing the values and colors of cells, and controlling the program flow. 

## Features:
Valid commands in Paintfuck include:

* n \- Move data pointer north (up)
* e \- Move data pointer east (right)
* s \- Move data pointer south (down)
* w \- Move data pointer west (left)
* \* \- Flip the bit at the current cell (same as in Smallfuck)
* [ \- Jump past matching ] if bit under current pointer is 0 (same as in Smallfuck)
* ] \- Jump back to the matching [ (if bit under current pointer is nonzero) (same as in Smallfuck)

## Usage:
Paintfuck programs typically involve manipulating cells and colors to achieve specific patterns or effects on the grid. Due to its esoteric nature, Paintfuck is mainly used for artistic and recreational purposes rather than practical programming tasks.

## Example:
```paintfuck
[]++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.
```
This example prints "Hello, world!" when interpreted.

