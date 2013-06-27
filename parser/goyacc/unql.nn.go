package goyacc

import (
	"log"
)
import (
	"strconv"
)
import ("bufio";"io";"strings")
type dfa struct {
  acc []bool
  f []func(rune) int
  id int
}
type family struct {
  a []dfa
  endcase int
}
var a0 [51]dfa
var a []family
func init() {
a = make([]family, 1)
{
var acc [18]bool
var fun [18]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 117: return 2
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 2
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[14] = func(r rune) int {
  switch(r) {
  case 34: return 4
  case 114: return 2
  case 98: return 15
  case 117: return 2
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 15
  case 116: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 15
    case 65 <= r && r <= 70: return 15
    case 97 <= r && r <= 102: return 15
    default: return 2
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 117: return 2
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 2
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 117: return 2
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 2
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 34: return -1
  case 114: return -1
  case 98: return -1
  case 117: return -1
  case 47: return -1
  case 92: return -1
  case 110: return -1
  case 102: return -1
  case 116: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return -1
    case 65 <= r && r <= 70: return -1
    case 97 <= r && r <= 102: return -1
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 92: return 3
  case 110: return 2
  case 102: return 2
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 2
  case 117: return 2
  case 47: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 116: return 5
  case 34: return 6
  case 114: return 7
  case 98: return 8
  case 117: return 9
  case 47: return 10
  case 92: return 11
  case 110: return 12
  case 102: return 13
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 34: return 6
  case 114: return 7
  case 98: return 8
  case 117: return 9
  case 47: return 10
  case 92: return 11
  case 110: return 12
  case 102: return 13
  case 116: return 5
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 92: return 3
  case 110: return 2
  case 102: return 2
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 2
  case 117: return 2
  case 47: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 114: return 2
  case 98: return 2
  case 117: return 2
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 2
  case 116: return 2
  case 34: return 4
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[12] = func(r rune) int {
  switch(r) {
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 2
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 2
  case 117: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[16] = func(r rune) int {
  switch(r) {
  case 98: return 17
  case 117: return 2
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 17
  case 116: return 2
  case 34: return 4
  case 114: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 17
    case 65 <= r && r <= 70: return 17
    case 97 <= r && r <= 102: return 17
    default: return 2
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 117: return 2
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 14
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 14
  default:
    switch {
    case 48 <= r && r <= 57: return 14
    case 65 <= r && r <= 70: return 14
    case 97 <= r && r <= 102: return 14
    default: return 2
    }
  }
  panic("unreachable")
}
fun[13] = func(r rune) int {
  switch(r) {
  case 92: return 3
  case 110: return 2
  case 102: return 2
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 2
  case 117: return 2
  case 47: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 117: return 2
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 2
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 116: return -1
  case 34: return 1
  case 114: return -1
  case 98: return -1
  case 117: return -1
  case 47: return -1
  case 92: return -1
  case 110: return -1
  case 102: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return -1
    case 65 <= r && r <= 70: return -1
    case 97 <= r && r <= 102: return -1
    default: return -1
    }
  }
  panic("unreachable")
}
fun[17] = func(r rune) int {
  switch(r) {
  case 102: return 2
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 2
  case 117: return 2
  case 47: return 2
  case 92: return 3
  case 110: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[15] = func(r rune) int {
  switch(r) {
  case 47: return 2
  case 92: return 3
  case 110: return 2
  case 102: return 16
  case 116: return 2
  case 34: return 4
  case 114: return 2
  case 98: return 16
  case 117: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 16
    case 65 <= r && r <= 70: return 16
    case 97 <= r && r <= 102: return 16
    default: return 2
    }
  }
  panic("unreachable")
}
a0[0].acc = acc[:]
a0[0].f = fun[:]
a0[0].id = 0
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 46: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 46: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[1].acc = acc[:]
a0[1].f = fun[:]
a0[1].id = 1
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 43: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 43: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[2].acc = acc[:]
a0[2].f = fun[:]
a0[2].id = 2
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 45: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 45: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[3].acc = acc[:]
a0[3].f = fun[:]
a0[3].id = 3
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 42: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 42: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[4].acc = acc[:]
a0[4].f = fun[:]
a0[4].id = 4
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 47: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 47: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[5].acc = acc[:]
a0[5].f = fun[:]
a0[5].id = 5
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 37: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 37: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[6].acc = acc[:]
a0[6].f = fun[:]
a0[6].id = 6
}
{
var acc [4]bool
var fun [4]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 110: return 2
  case 68: return -1
  case 100: return -1
  case 65: return -1
  case 78: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[3] = true
fun[3] = func(r rune) int {
  switch(r) {
  case 68: return -1
  case 100: return -1
  case 65: return -1
  case 78: return -1
  case 97: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 97: return 1
  case 110: return -1
  case 68: return -1
  case 100: return -1
  case 65: return 1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 68: return 3
  case 100: return 3
  case 65: return -1
  case 78: return -1
  case 97: return -1
  case 110: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[7].acc = acc[:]
a0[7].f = fun[:]
a0[7].id = 7
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 79: return 1
  case 114: return -1
  case 111: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 82: return 2
  case 79: return -1
  case 114: return 2
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 114: return -1
  case 111: return -1
  case 82: return -1
  case 79: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[8].acc = acc[:]
a0[8].f = fun[:]
a0[8].id = 8
}
{
var acc [3]bool
var fun [3]func(rune) int
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 61: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 61: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[9].acc = acc[:]
a0[9].f = fun[:]
a0[9].id = 9
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 61: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[10].acc = acc[:]
a0[10].f = fun[:]
a0[10].id = 10
}
{
var acc [3]bool
var fun [3]func(rune) int
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 33: return -1
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 33: return 1
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 33: return -1
  case 61: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[11].acc = acc[:]
a0[11].f = fun[:]
a0[11].id = 11
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 60: return 1
  case 62: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 60: return -1
  case 62: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 60: return -1
  case 62: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[12].acc = acc[:]
a0[12].f = fun[:]
a0[12].id = 12
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 60: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 60: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[13].acc = acc[:]
a0[13].f = fun[:]
a0[13].id = 13
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 60: return 1
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 60: return -1
  case 61: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 60: return -1
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[14].acc = acc[:]
a0[14].f = fun[:]
a0[14].id = 14
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 62: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 62: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[15].acc = acc[:]
a0[15].f = fun[:]
a0[15].id = 15
}
{
var acc [3]bool
var fun [3]func(rune) int
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 61: return -1
  case 62: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 62: return 1
  case 61: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 62: return -1
  case 61: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[16].acc = acc[:]
a0[16].f = fun[:]
a0[16].id = 16
}
{
var acc [4]bool
var fun [4]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 110: return 1
  case 111: return -1
  case 79: return -1
  case 116: return -1
  case 84: return -1
  case 78: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 79: return -1
  case 116: return 3
  case 84: return 3
  case 78: return -1
  case 110: return -1
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 111: return 2
  case 79: return 2
  case 116: return -1
  case 84: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[3] = true
fun[3] = func(r rune) int {
  switch(r) {
  case 79: return -1
  case 116: return -1
  case 84: return -1
  case 78: return -1
  case 110: return -1
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[17].acc = acc[:]
a0[17].f = fun[:]
a0[17].id = 17
}
{
var acc [5]bool
var fun [5]func(rune) int
fun[3] = func(r rune) int {
  switch(r) {
  case 107: return -1
  case 69: return 4
  case 73: return -1
  case 75: return -1
  case 105: return -1
  case 108: return -1
  case 76: return -1
  case 101: return 4
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 75: return -1
  case 105: return -1
  case 108: return 1
  case 76: return 1
  case 101: return -1
  case 107: return -1
  case 69: return -1
  case 73: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 107: return -1
  case 69: return -1
  case 73: return -1
  case 75: return -1
  case 105: return -1
  case 108: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 107: return 3
  case 69: return -1
  case 73: return -1
  case 75: return 3
  case 105: return -1
  case 108: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 107: return -1
  case 69: return -1
  case 73: return 2
  case 75: return -1
  case 105: return 2
  case 108: return -1
  case 76: return -1
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[18].acc = acc[:]
a0[18].f = fun[:]
a0[18].id = 18
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 115: return 2
  case 105: return -1
  case 73: return -1
  case 83: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 73: return -1
  case 83: return -1
  case 115: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 105: return 1
  case 73: return 1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[19].acc = acc[:]
a0[19].f = fun[:]
a0[19].id = 19
}
{
var acc [8]bool
var fun [8]func(rune) int
fun[3] = func(r rune) int {
  switch(r) {
  case 83: return 4
  case 71: return -1
  case 77: return -1
  case 110: return -1
  case 115: return 4
  case 105: return -1
  case 103: return -1
  case 109: return -1
  case 78: return -1
  case 73: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 77: return -1
  case 110: return 6
  case 115: return -1
  case 105: return -1
  case 103: return -1
  case 109: return -1
  case 78: return 6
  case 73: return -1
  case 83: return -1
  case 71: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 77: return -1
  case 110: return -1
  case 115: return -1
  case 105: return 5
  case 103: return -1
  case 109: return -1
  case 78: return -1
  case 73: return 5
  case 83: return -1
  case 71: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[7] = true
fun[7] = func(r rune) int {
  switch(r) {
  case 103: return -1
  case 109: return -1
  case 78: return -1
  case 73: return -1
  case 83: return -1
  case 71: return -1
  case 77: return -1
  case 110: return -1
  case 115: return -1
  case 105: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 103: return -1
  case 109: return -1
  case 78: return -1
  case 73: return 2
  case 83: return -1
  case 71: return -1
  case 77: return -1
  case 110: return -1
  case 115: return -1
  case 105: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 115: return 3
  case 105: return -1
  case 103: return -1
  case 109: return -1
  case 78: return -1
  case 73: return -1
  case 83: return 3
  case 71: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 71: return 7
  case 77: return -1
  case 110: return -1
  case 115: return -1
  case 105: return -1
  case 103: return 7
  case 109: return -1
  case 78: return -1
  case 73: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 71: return -1
  case 77: return 1
  case 110: return -1
  case 115: return -1
  case 105: return -1
  case 103: return -1
  case 109: return 1
  case 78: return -1
  case 73: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[20].acc = acc[:]
a0[20].f = fun[:]
a0[20].id = 20
}
{
var acc [7]bool
var fun [7]func(rune) int
fun[5] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 68: return 6
  case 69: return -1
  case 108: return -1
  case 101: return -1
  case 86: return -1
  case 85: return -1
  case 117: return -1
  case 76: return -1
  case 100: return 6
  case 118: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 100: return -1
  case 118: return 1
  case 97: return -1
  case 65: return -1
  case 68: return -1
  case 69: return -1
  case 108: return -1
  case 101: return -1
  case 86: return 1
  case 85: return -1
  case 117: return -1
  case 76: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 86: return -1
  case 85: return -1
  case 117: return -1
  case 76: return 3
  case 100: return -1
  case 118: return -1
  case 97: return -1
  case 65: return -1
  case 68: return -1
  case 69: return -1
  case 108: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 65: return -1
  case 68: return -1
  case 69: return 5
  case 108: return -1
  case 101: return 5
  case 86: return -1
  case 85: return -1
  case 117: return -1
  case 76: return -1
  case 100: return -1
  case 118: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 85: return -1
  case 117: return -1
  case 76: return -1
  case 100: return -1
  case 118: return -1
  case 97: return 2
  case 65: return 2
  case 68: return -1
  case 69: return -1
  case 108: return -1
  case 101: return -1
  case 86: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 108: return -1
  case 101: return -1
  case 86: return -1
  case 85: return -1
  case 117: return -1
  case 76: return -1
  case 100: return -1
  case 118: return -1
  case 97: return -1
  case 65: return -1
  case 68: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 65: return -1
  case 68: return -1
  case 69: return -1
  case 108: return -1
  case 101: return -1
  case 86: return -1
  case 85: return 4
  case 117: return 4
  case 76: return -1
  case 100: return -1
  case 118: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[21].acc = acc[:]
a0[21].f = fun[:]
a0[21].id = 21
}
{
var acc [7]bool
var fun [7]func(rune) int
fun[2] = func(r rune) int {
  switch(r) {
  case 116: return -1
  case 108: return 3
  case 67: return -1
  case 76: return 3
  case 84: return -1
  case 83: return -1
  case 101: return -1
  case 99: return -1
  case 115: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 76: return -1
  case 84: return -1
  case 83: return -1
  case 101: return 2
  case 99: return -1
  case 115: return -1
  case 69: return 2
  case 116: return -1
  case 108: return -1
  case 67: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 101: return 4
  case 99: return -1
  case 115: return -1
  case 69: return 4
  case 116: return -1
  case 108: return -1
  case 67: return -1
  case 76: return -1
  case 84: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 115: return 1
  case 69: return -1
  case 116: return -1
  case 108: return -1
  case 67: return -1
  case 76: return -1
  case 84: return -1
  case 83: return 1
  case 101: return -1
  case 99: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 67: return 5
  case 76: return -1
  case 84: return -1
  case 83: return -1
  case 101: return -1
  case 99: return 5
  case 115: return -1
  case 69: return -1
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 69: return -1
  case 116: return 6
  case 108: return -1
  case 67: return -1
  case 76: return -1
  case 84: return 6
  case 83: return -1
  case 101: return -1
  case 99: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 101: return -1
  case 99: return -1
  case 115: return -1
  case 69: return -1
  case 116: return -1
  case 108: return -1
  case 67: return -1
  case 76: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[22].acc = acc[:]
a0[22].f = fun[:]
a0[22].id = 22
}
{
var acc [3]bool
var fun [3]func(rune) int
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 65: return -1
  case 97: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 83: return 2
  case 115: return 2
  case 65: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 115: return -1
  case 65: return 1
  case 97: return 1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[23].acc = acc[:]
a0[23].f = fun[:]
a0[23].id = 23
}
{
var acc [5]bool
var fun [5]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 102: return -1
  case 111: return -1
  case 77: return -1
  case 109: return -1
  case 70: return -1
  case 79: return -1
  case 114: return 2
  case 82: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 79: return -1
  case 114: return -1
  case 82: return -1
  case 102: return 1
  case 111: return -1
  case 77: return -1
  case 109: return -1
  case 70: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 102: return -1
  case 111: return -1
  case 77: return 4
  case 109: return 4
  case 70: return -1
  case 79: return -1
  case 114: return -1
  case 82: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 77: return -1
  case 109: return -1
  case 70: return -1
  case 79: return 3
  case 114: return -1
  case 82: return -1
  case 102: return -1
  case 111: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 114: return -1
  case 82: return -1
  case 102: return -1
  case 111: return -1
  case 77: return -1
  case 109: return -1
  case 70: return -1
  case 79: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[24].acc = acc[:]
a0[24].f = fun[:]
a0[24].id = 24
}
{
var acc [6]bool
var fun [6]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 72: return 2
  case 119: return -1
  case 114: return -1
  case 87: return -1
  case 69: return -1
  case 82: return -1
  case 104: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 101: return 5
  case 72: return -1
  case 119: return -1
  case 114: return -1
  case 87: return -1
  case 69: return 5
  case 82: return -1
  case 104: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[5] = true
fun[5] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 72: return -1
  case 119: return -1
  case 114: return -1
  case 87: return -1
  case 69: return -1
  case 82: return -1
  case 104: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 72: return -1
  case 119: return -1
  case 114: return -1
  case 87: return -1
  case 69: return 3
  case 82: return -1
  case 104: return -1
  case 101: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 114: return 4
  case 87: return -1
  case 69: return -1
  case 82: return 4
  case 104: return -1
  case 101: return -1
  case 72: return -1
  case 119: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 104: return -1
  case 101: return -1
  case 72: return -1
  case 119: return 1
  case 114: return -1
  case 87: return 1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[25].acc = acc[:]
a0[25].f = fun[:]
a0[25].id = 25
}
{
var acc [6]bool
var fun [6]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 111: return -1
  case 79: return -1
  case 82: return 2
  case 100: return -1
  case 101: return -1
  case 68: return -1
  case 114: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[5] = true
fun[5] = func(r rune) int {
  switch(r) {
  case 82: return -1
  case 100: return -1
  case 101: return -1
  case 68: return -1
  case 114: return -1
  case 69: return -1
  case 111: return -1
  case 79: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 114: return -1
  case 69: return -1
  case 111: return 1
  case 79: return 1
  case 82: return -1
  case 100: return -1
  case 101: return -1
  case 68: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 68: return 3
  case 114: return -1
  case 69: return -1
  case 111: return -1
  case 79: return -1
  case 82: return -1
  case 100: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 111: return -1
  case 79: return -1
  case 82: return 5
  case 100: return -1
  case 101: return -1
  case 68: return -1
  case 114: return 5
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 69: return 4
  case 111: return -1
  case 79: return -1
  case 82: return -1
  case 100: return -1
  case 101: return 4
  case 68: return -1
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[26].acc = acc[:]
a0[26].f = fun[:]
a0[26].id = 26
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 98: return -1
  case 66: return -1
  case 89: return 2
  case 121: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 98: return -1
  case 66: return -1
  case 89: return -1
  case 121: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 98: return 1
  case 66: return 1
  case 89: return -1
  case 121: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[27].acc = acc[:]
a0[27].f = fun[:]
a0[27].id = 27
}
{
var acc [4]bool
var fun [4]func(rune) int
acc[3] = true
fun[3] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 115: return -1
  case 97: return -1
  case 83: return -1
  case 67: return -1
  case 65: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 115: return -1
  case 97: return 1
  case 83: return -1
  case 67: return -1
  case 65: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 97: return -1
  case 83: return 2
  case 67: return -1
  case 65: return -1
  case 99: return -1
  case 115: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 65: return -1
  case 99: return 3
  case 115: return -1
  case 97: return -1
  case 83: return -1
  case 67: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[28].acc = acc[:]
a0[28].f = fun[:]
a0[28].id = 28
}
{
var acc [5]bool
var fun [5]func(rune) int
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 99: return -1
  case 115: return -1
  case 67: return -1
  case 101: return -1
  case 68: return -1
  case 100: return -1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 99: return -1
  case 115: return 3
  case 67: return -1
  case 101: return -1
  case 68: return -1
  case 100: return -1
  case 69: return -1
  case 83: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 99: return -1
  case 115: return -1
  case 67: return -1
  case 101: return -1
  case 68: return 1
  case 100: return 1
  case 69: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 99: return 4
  case 115: return -1
  case 67: return 4
  case 101: return -1
  case 68: return -1
  case 100: return -1
  case 69: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 99: return -1
  case 115: return -1
  case 67: return -1
  case 101: return 2
  case 68: return -1
  case 100: return -1
  case 69: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[29].acc = acc[:]
a0[29].f = fun[:]
a0[29].id = 29
}
{
var acc [6]bool
var fun [6]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 105: return 2
  case 116: return -1
  case 108: return -1
  case 77: return -1
  case 73: return 2
  case 84: return -1
  case 76: return -1
  case 109: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 109: return -1
  case 105: return -1
  case 116: return -1
  case 108: return 1
  case 77: return -1
  case 73: return -1
  case 84: return -1
  case 76: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 73: return -1
  case 84: return 5
  case 76: return -1
  case 109: return -1
  case 105: return -1
  case 116: return 5
  case 108: return -1
  case 77: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 77: return 3
  case 73: return -1
  case 84: return -1
  case 76: return -1
  case 109: return 3
  case 105: return -1
  case 116: return -1
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 108: return -1
  case 77: return -1
  case 73: return 4
  case 84: return -1
  case 76: return -1
  case 109: return -1
  case 105: return 4
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[5] = true
fun[5] = func(r rune) int {
  switch(r) {
  case 105: return -1
  case 116: return -1
  case 108: return -1
  case 77: return -1
  case 73: return -1
  case 84: return -1
  case 76: return -1
  case 109: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[30].acc = acc[:]
a0[30].f = fun[:]
a0[30].id = 30
}
{
var acc [7]bool
var fun [7]func(rune) int
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 84: return -1
  case 116: return -1
  case 111: return -1
  case 101: return -1
  case 115: return -1
  case 79: return -1
  case 102: return -1
  case 70: return -1
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 116: return -1
  case 111: return -1
  case 101: return -1
  case 115: return 4
  case 79: return -1
  case 102: return -1
  case 70: return -1
  case 83: return 4
  case 69: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 69: return -1
  case 84: return -1
  case 116: return -1
  case 111: return 1
  case 101: return -1
  case 115: return -1
  case 79: return 1
  case 102: return -1
  case 70: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 84: return -1
  case 116: return -1
  case 111: return -1
  case 101: return -1
  case 115: return -1
  case 79: return -1
  case 102: return 3
  case 70: return 3
  case 83: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 116: return -1
  case 111: return -1
  case 101: return -1
  case 115: return -1
  case 79: return -1
  case 102: return 2
  case 70: return 2
  case 83: return -1
  case 69: return -1
  case 84: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 83: return -1
  case 69: return -1
  case 84: return 6
  case 116: return 6
  case 111: return -1
  case 101: return -1
  case 115: return -1
  case 79: return -1
  case 102: return -1
  case 70: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 101: return 5
  case 115: return -1
  case 79: return -1
  case 102: return -1
  case 70: return -1
  case 83: return -1
  case 69: return 5
  case 84: return -1
  case 116: return -1
  case 111: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[31].acc = acc[:]
a0[31].f = fun[:]
a0[31].id = 31
}
{
var acc [3]bool
var fun [3]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 124: return 2
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 124: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 124: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[32].acc = acc[:]
a0[32].f = fun[:]
a0[32].id = 32
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 40: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 40: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[33].acc = acc[:]
a0[33].f = fun[:]
a0[33].id = 33
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 41: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 41: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[34].acc = acc[:]
a0[34].f = fun[:]
a0[34].id = 34
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 123: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 123: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[35].acc = acc[:]
a0[35].f = fun[:]
a0[35].id = 35
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 125: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 125: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[36].acc = acc[:]
a0[36].f = fun[:]
a0[36].id = 36
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 44: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 44: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[37].acc = acc[:]
a0[37].f = fun[:]
a0[37].id = 37
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 58: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 58: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[38].acc = acc[:]
a0[38].f = fun[:]
a0[38].id = 38
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 91: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 91: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[39].acc = acc[:]
a0[39].f = fun[:]
a0[39].id = 39
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 93: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 93: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[40].acc = acc[:]
a0[40].f = fun[:]
a0[40].id = 40
}
{
var acc [2]bool
var fun [2]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 34: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 34: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[41].acc = acc[:]
a0[41].f = fun[:]
a0[41].id = 41
}
{
var acc [5]bool
var fun [5]func(rune) int
fun[3] = func(r rune) int {
  switch(r) {
  case 69: return 4
  case 85: return -1
  case 117: return -1
  case 116: return -1
  case 84: return -1
  case 82: return -1
  case 101: return 4
  case 114: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 84: return -1
  case 82: return 2
  case 101: return -1
  case 114: return 2
  case 69: return -1
  case 85: return -1
  case 117: return -1
  case 116: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 114: return -1
  case 69: return -1
  case 85: return -1
  case 117: return -1
  case 116: return -1
  case 84: return -1
  case 82: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 84: return 1
  case 82: return -1
  case 101: return -1
  case 114: return -1
  case 69: return -1
  case 85: return -1
  case 117: return -1
  case 116: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 114: return -1
  case 69: return -1
  case 85: return 3
  case 117: return 3
  case 116: return -1
  case 84: return -1
  case 82: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[42].acc = acc[:]
a0[42].f = fun[:]
a0[42].id = 42
}
{
var acc [6]bool
var fun [6]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 102: return 1
  case 69: return -1
  case 115: return -1
  case 108: return -1
  case 76: return -1
  case 70: return 1
  case 83: return -1
  case 65: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 102: return -1
  case 69: return 5
  case 115: return -1
  case 108: return -1
  case 76: return -1
  case 70: return -1
  case 83: return -1
  case 65: return -1
  case 97: return -1
  case 101: return 5
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[5] = true
fun[5] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 108: return -1
  case 76: return -1
  case 70: return -1
  case 83: return -1
  case 65: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 108: return -1
  case 76: return -1
  case 70: return -1
  case 83: return -1
  case 65: return 2
  case 97: return 2
  case 101: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return 4
  case 108: return -1
  case 76: return -1
  case 70: return -1
  case 83: return 4
  case 65: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 101: return -1
  case 102: return -1
  case 69: return -1
  case 115: return -1
  case 108: return 3
  case 76: return 3
  case 70: return -1
  case 83: return -1
  case 65: return -1
  case 97: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[43].acc = acc[:]
a0[43].f = fun[:]
a0[43].id = 43
}
{
var acc [5]bool
var fun [5]func(rune) int
fun[0] = func(r rune) int {
  switch(r) {
  case 76: return -1
  case 110: return 1
  case 85: return -1
  case 78: return 1
  case 108: return -1
  case 117: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 108: return 4
  case 117: return -1
  case 76: return 4
  case 110: return -1
  case 85: return -1
  case 78: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 117: return -1
  case 76: return 3
  case 110: return -1
  case 85: return -1
  case 78: return -1
  case 108: return 3
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 117: return -1
  case 76: return -1
  case 110: return -1
  case 85: return -1
  case 78: return -1
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 117: return 2
  case 76: return -1
  case 110: return -1
  case 85: return 2
  case 78: return -1
  case 108: return -1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[44].acc = acc[:]
a0[44].f = fun[:]
a0[44].id = 44
}
{
var acc [12]bool
var fun [12]func(rune) int
fun[1] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 43: return -1
  case 101: return -1
  case 69: return -1
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 2
    case 49 <= r && r <= 57: return 3
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 43: return -1
  case 101: return -1
  case 69: return -1
  case 45: return 1
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 2
    case 49 <= r && r <= 57: return 3
    default: return -1
    }
  }
  panic("unreachable")
}
acc[8] = true
fun[8] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 43: return -1
  case 101: return 7
  case 69: return 7
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 8
    case 49 <= r && r <= 57: return 8
    default: return -1
    }
  }
  panic("unreachable")
}
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 43: return -1
  case 101: return -1
  case 69: return -1
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 11
    case 49 <= r && r <= 57: return 11
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 43: return -1
  case 101: return -1
  case 69: return -1
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 10
    case 49 <= r && r <= 57: return 10
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 43: return 9
  case 101: return -1
  case 69: return -1
  case 45: return 9
  default:
    switch {
    case 48 <= r && r <= 48: return 10
    case 49 <= r && r <= 57: return 10
    default: return -1
    }
  }
  panic("unreachable")
}
acc[11] = true
fun[11] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 43: return -1
  case 101: return -1
  case 69: return -1
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 11
    case 49 <= r && r <= 57: return 11
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 43: return -1
  case 101: return -1
  case 69: return -1
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 6
    case 49 <= r && r <= 57: return 6
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 45: return -1
  case 46: return 4
  case 43: return -1
  case 101: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 5
    case 49 <= r && r <= 57: return 5
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 46: return 4
  case 43: return -1
  case 101: return -1
  case 69: return -1
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return -1
    case 49 <= r && r <= 57: return -1
    default: return -1
    }
  }
  panic("unreachable")
}
acc[6] = true
fun[6] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 43: return -1
  case 101: return 7
  case 69: return 7
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 8
    case 49 <= r && r <= 57: return 8
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 46: return 4
  case 43: return -1
  case 101: return -1
  case 69: return -1
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 5
    case 49 <= r && r <= 57: return 5
    default: return -1
    }
  }
  panic("unreachable")
}
a0[45].acc = acc[:]
a0[45].f = fun[:]
a0[45].id = 45
}
{
var acc [12]bool
var fun [12]func(rune) int
acc[11] = true
fun[11] = func(r rune) int {
  switch(r) {
  case 45: return -1
  case 43: return -1
  case 101: return -1
  case 46: return -1
  case 69: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 11
    case 49 <= r && r <= 57: return 11
    default: return -1
    }
  }
  panic("unreachable")
}
fun[4] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 45: return 9
  case 43: return 9
  case 101: return -1
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 10
    case 49 <= r && r <= 57: return 10
    default: return -1
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 45: return -1
  case 43: return -1
  case 101: return -1
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 10
    case 49 <= r && r <= 57: return 10
    default: return -1
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 101: return 4
  case 46: return 5
  case 69: return 4
  case 45: return -1
  case 43: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 6
    case 49 <= r && r <= 57: return 6
    default: return -1
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 69: return 4
  case 45: return -1
  case 43: return -1
  case 101: return 4
  case 46: return 5
  default:
    switch {
    case 48 <= r && r <= 48: return 6
    case 49 <= r && r <= 57: return 6
    default: return -1
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 46: return -1
  case 69: return 4
  case 45: return -1
  case 43: return -1
  case 101: return 4
  default:
    switch {
    case 48 <= r && r <= 48: return 8
    case 49 <= r && r <= 57: return 8
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 45: return 1
  case 43: return -1
  case 101: return -1
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 2
    case 49 <= r && r <= 57: return 3
    default: return -1
    }
  }
  panic("unreachable")
}
acc[10] = true
fun[10] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 45: return -1
  case 43: return -1
  case 101: return -1
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 11
    case 49 <= r && r <= 57: return 11
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 69: return -1
  case 45: return -1
  case 43: return -1
  case 101: return -1
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 2
    case 49 <= r && r <= 57: return 3
    default: return -1
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 69: return 4
  case 45: return -1
  case 43: return -1
  case 101: return 4
  case 46: return 5
  default:
    switch {
    case 48 <= r && r <= 48: return -1
    case 49 <= r && r <= 57: return -1
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 69: return 4
  case 45: return -1
  case 43: return -1
  case 101: return 4
  case 46: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 8
    case 49 <= r && r <= 57: return 8
    default: return -1
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 43: return -1
  case 101: return -1
  case 46: return -1
  case 69: return -1
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 7
    case 49 <= r && r <= 57: return 7
    default: return -1
    }
  }
  panic("unreachable")
}
a0[46].acc = acc[:]
a0[46].f = fun[:]
a0[46].id = 46
}
{
var acc [5]bool
var fun [5]func(rune) int
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return -1
    case 49 <= r && r <= 57: return -1
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 45: return 1
  default:
    switch {
    case 48 <= r && r <= 48: return 2
    case 49 <= r && r <= 57: return 3
    default: return -1
    }
  }
  panic("unreachable")
}
acc[3] = true
fun[3] = func(r rune) int {
  switch(r) {
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 4
    case 49 <= r && r <= 57: return 4
    default: return -1
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 2
    case 49 <= r && r <= 57: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 45: return -1
  default:
    switch {
    case 48 <= r && r <= 48: return 4
    case 49 <= r && r <= 57: return 4
    default: return -1
    }
  }
  panic("unreachable")
}
a0[47].acc = acc[:]
a0[47].f = fun[:]
a0[47].id = 47
}
{
var acc [2]bool
var fun [2]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 9: return 1
  case 10: return 1
  case 32: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 9: return 1
  case 10: return 1
  case 32: return 1
  default:
    switch {
    default: return -1
    }
  }
  panic("unreachable")
}
a0[48].acc = acc[:]
a0[48].f = fun[:]
a0[48].id = 48
}
{
var acc [3]bool
var fun [3]func(rune) int
acc[1] = true
fun[1] = func(r rune) int {
  switch(r) {
  case 45: return 2
  case 95: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 90: return 2
    case 97 <= r && r <= 122: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 45: return -1
  case 95: return 1
  default:
    switch {
    case 48 <= r && r <= 57: return -1
    case 65 <= r && r <= 90: return 1
    case 97 <= r && r <= 122: return 1
    default: return -1
    }
  }
  panic("unreachable")
}
acc[2] = true
fun[2] = func(r rune) int {
  switch(r) {
  case 45: return 2
  case 95: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 90: return 2
    case 97 <= r && r <= 122: return 2
    default: return -1
    }
  }
  panic("unreachable")
}
a0[49].acc = acc[:]
a0[49].f = fun[:]
a0[49].id = 49
}
{
var acc [18]bool
var fun [18]func(rune) int
fun[12] = func(r rune) int {
  switch(r) {
  case 116: return 2
  case 96: return 4
  case 110: return 2
  case 102: return 2
  case 92: return 3
  case 98: return 2
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[16] = func(r rune) int {
  switch(r) {
  case 102: return 17
  case 92: return 3
  case 98: return 17
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  case 110: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 17
    case 65 <= r && r <= 70: return 17
    case 97 <= r && r <= 102: return 17
    default: return 2
    }
  }
  panic("unreachable")
}
acc[4] = true
fun[4] = func(r rune) int {
  switch(r) {
  case 110: return -1
  case 102: return -1
  case 92: return -1
  case 98: return -1
  case 114: return -1
  case 117: return -1
  case 47: return -1
  case 34: return -1
  case 116: return -1
  case 96: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return -1
    case 65 <= r && r <= 70: return -1
    case 97 <= r && r <= 102: return -1
    default: return -1
    }
  }
  panic("unreachable")
}
fun[7] = func(r rune) int {
  switch(r) {
  case 110: return 5
  case 102: return 6
  case 92: return 7
  case 98: return 8
  case 114: return 9
  case 117: return 10
  case 47: return 11
  case 34: return 12
  case 116: return 13
  case 96: return 4
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[5] = func(r rune) int {
  switch(r) {
  case 102: return 2
  case 92: return 3
  case 98: return 2
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  case 110: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[17] = func(r rune) int {
  switch(r) {
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  case 110: return 2
  case 102: return 2
  case 92: return 3
  case 98: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[9] = func(r rune) int {
  switch(r) {
  case 102: return 2
  case 92: return 3
  case 98: return 2
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  case 110: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[6] = func(r rune) int {
  switch(r) {
  case 92: return 3
  case 98: return 2
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  case 110: return 2
  case 102: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[1] = func(r rune) int {
  switch(r) {
  case 102: return 2
  case 92: return 3
  case 98: return 2
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return -1
  case 110: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[8] = func(r rune) int {
  switch(r) {
  case 116: return 2
  case 96: return 4
  case 110: return 2
  case 102: return 2
  case 92: return 3
  case 98: return 2
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[3] = func(r rune) int {
  switch(r) {
  case 96: return 4
  case 110: return 5
  case 102: return 6
  case 92: return 7
  case 98: return 8
  case 114: return 9
  case 117: return 10
  case 47: return 11
  case 34: return 12
  case 116: return 13
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[14] = func(r rune) int {
  switch(r) {
  case 110: return 2
  case 102: return 15
  case 92: return 3
  case 98: return 15
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  default:
    switch {
    case 48 <= r && r <= 57: return 15
    case 65 <= r && r <= 70: return 15
    case 97 <= r && r <= 102: return 15
    default: return 2
    }
  }
  panic("unreachable")
}
fun[0] = func(r rune) int {
  switch(r) {
  case 114: return -1
  case 117: return -1
  case 47: return -1
  case 34: return -1
  case 116: return -1
  case 96: return 1
  case 110: return -1
  case 102: return -1
  case 92: return -1
  case 98: return -1
  default:
    switch {
    case 48 <= r && r <= 57: return -1
    case 65 <= r && r <= 70: return -1
    case 97 <= r && r <= 102: return -1
    default: return -1
    }
  }
  panic("unreachable")
}
fun[10] = func(r rune) int {
  switch(r) {
  case 102: return 14
  case 92: return 3
  case 98: return 14
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  case 110: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 14
    case 65 <= r && r <= 70: return 14
    case 97 <= r && r <= 102: return 14
    default: return 2
    }
  }
  panic("unreachable")
}
fun[2] = func(r rune) int {
  switch(r) {
  case 110: return 2
  case 102: return 2
  case 92: return 3
  case 98: return 2
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[11] = func(r rune) int {
  switch(r) {
  case 110: return 2
  case 102: return 2
  case 92: return 3
  case 98: return 2
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
fun[15] = func(r rune) int {
  switch(r) {
  case 98: return 16
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  case 110: return 2
  case 102: return 16
  case 92: return 3
  default:
    switch {
    case 48 <= r && r <= 57: return 16
    case 65 <= r && r <= 70: return 16
    case 97 <= r && r <= 102: return 16
    default: return 2
    }
  }
  panic("unreachable")
}
fun[13] = func(r rune) int {
  switch(r) {
  case 102: return 2
  case 92: return 3
  case 98: return 2
  case 114: return 2
  case 117: return 2
  case 47: return 2
  case 34: return 2
  case 116: return 2
  case 96: return 4
  case 110: return 2
  default:
    switch {
    case 48 <= r && r <= 57: return 2
    case 65 <= r && r <= 70: return 2
    case 97 <= r && r <= 102: return 2
    default: return 2
    }
  }
  panic("unreachable")
}
a0[50].acc = acc[:]
a0[50].f = fun[:]
a0[50].id = 50
}
a[0].endcase = 51
a[0].a = a0[:]
}
func getAction(c *frame) int {
  if -1 == c.match { return -1 }
  c.action = c.fam.a[c.match].id
  c.match = -1
  return c.action
}
type frame struct {
  atEOF bool
  action, match, matchn, n int
  buf []rune
  text string
  in *bufio.Reader
  state []int
  fam family
}
func newFrame(in *bufio.Reader, index int) *frame {
  f := new(frame)
  f.buf = make([]rune, 0, 128)
  f.in = in
  f.match = -1
  f.fam = a[index]
  f.state = make([]int, len(f.fam.a))
  return f
}
type Lexer []*frame
func NewLexer(in io.Reader) Lexer {
  stack := make([]*frame, 0, 4)
  stack = append(stack, newFrame(bufio.NewReader(in), 0))
  return stack
}
func (stack Lexer) isDone() bool {
  return 1 == len(stack) && stack[0].atEOF
}
func (stack Lexer) nextAction() int {
  c := stack[len(stack) - 1]
  for {
    if c.atEOF { return c.fam.endcase }
    if c.n == len(c.buf) {
      r,_,er := c.in.ReadRune()
      switch er {
      case nil: c.buf = append(c.buf, r)
      case io.EOF:
	c.atEOF = true
	if c.n > 0 {
	  c.text = string(c.buf)
	  return getAction(c)
	}
	return c.fam.endcase
      default: panic(er.Error())
      }
    }
    jammed := true
    r := c.buf[c.n]
    for i, x := range c.fam.a {
      if -1 == c.state[i] { continue }
      c.state[i] = x.f[c.state[i]](r)
      if -1 == c.state[i] { continue }
      jammed = false
      if x.acc[c.state[i]] {
	if -1 == c.match || c.matchn < c.n+1 || c.match > i {
	  c.match = i
	  c.matchn = c.n+1
	}
      }
    }
    if jammed {
      a := getAction(c)
      if -1 == a { c.matchn = c.n + 1 }
      c.n = 0
      for i, _ := range c.state { c.state[i] = 0 }
      c.text = string(c.buf[:c.matchn])
      copy(c.buf, c.buf[c.matchn:])
      c.buf = c.buf[:len(c.buf) - c.matchn]
      return a
    }
    c.n++
  }
  panic("unreachable")
}
func (stack Lexer) push(index int) Lexer {
  c := stack[len(stack) - 1]
  return append(stack,
      newFrame(bufio.NewReader(strings.NewReader(c.text)), index))
}
func (stack Lexer) pop() Lexer {
  return stack[:len(stack) - 1]
}
func (stack Lexer) Text() string {
  c := stack[len(stack) - 1]
  return c.text
}
func (yylex Lexer) Error(e string) {
  panic(e)
}
func (yylex Lexer) Lex(lval *yySymType) int {
  for !yylex.isDone() {
    switch yylex.nextAction() {
    case -1:
    case 0:  //\"((\\\")|(\\\\)|(\\\/)|(\\b)|(\\f)|(\\n)|(\\r)|(\\t)|(\\u[0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F])|[^\"])*\"/
{ 
                    lval.s = yylex.Text()[1:len(yylex.Text())-1]
                    logDebugTokens("%s", lval.s)
                    logDebugTokens("STRING"); 
                    return STRING 
                  }
    case 1:  //\./
{ logDebugTokens("DOT"); return DOT }
    case 2:  //\+/
{ logDebugTokens("PLUS"); return PLUS }
    case 3:  //-/
{ logDebugTokens("MINUS"); return MINUS }
    case 4:  //\*/
{ logDebugTokens("MULT"); return MULT }
    case 5:  //\//
{ logDebugTokens("DIV"); return DIV }
    case 6:  //%/
{ logDebugTokens("MOD"); return MOD }
    case 7:  //[aA][nN][dD]/
{ logDebugTokens("AND"); return AND }
    case 8:  //[oO][rR]/
{ logDebugTokens("OR"); return OR }
    case 9:  //\=\=/
{ logDebugTokens("EQ"); return EQ }
    case 10:  //\=/
{ logDebugTokens("EQ"); return EQ }
    case 11:  //\!\=/
{ logDebugTokens("NE"); return NE }
    case 12:  //\<\>/
{ logDebugTokens("NE"); return NE }
    case 13:  //\</
{ logDebugTokens("LT"); return LT }
    case 14:  //\<\=/
{ logDebugTokens("LTE"); return LTE }
    case 15:  //\>/
{ logDebugTokens("GT"); return GT }
    case 16:  //\>\=/
{ logDebugTokens("GTE"); return GTE }
    case 17:  //[nN][oO][tT]/
{ logDebugTokens("NOT"); return NOT }
    case 18:  //[lL][iI][kK][eE]/
{ logDebugTokens("LIKE"); return LIKE }
    case 19:  //[iI][sS]/
{ logDebugTokens("IS"); return IS }
    case 20:  //[mM][iI][sS][sS][iI][nN][gG]/
{ logDebugTokens("MISSING"); return MISSING }
    case 21:  //[vV][aA][lL][uU][eE][dD]/
{ logDebugTokens("VALUED"); return VALUED }
    case 22:  //[sS][eE][lL][eE][cC][tT]/
{ logDebugTokens("SELECT"); return SELECT }
    case 23:  //[aA][sS]/
{ logDebugTokens("AS"); return AS }
    case 24:  //[fF][rR][oO][mM]/
{ logDebugTokens("FROM"); return FROM }
    case 25:  //[wW][hH][eE][rR][eE]/
{ logDebugTokens("WHERE"); return WHERE }
    case 26:  //[oO][rR][dD][eE][rR]/
{ logDebugTokens("ORDER"); return ORDER }
    case 27:  //[bB][yY]/
{ logDebugTokens("BY"); return BY }
    case 28:  //[aA][sS][cC]/
{ logDebugTokens("ASC"); return ASC }
    case 29:  //[dD][eE][sS][cC]/
{ logDebugTokens("DESC"); return DESC }
    case 30:  //[lL][iI][mM][iI][tT]/
{ logDebugTokens("LIMIT"); return LIMIT }
    case 31:  //[oO][fF][fF][sS][eE][tT]/
{ logDebugTokens("OFFSET"); return OFFSET }
    case 32:  //\|\|/
{ logDebugTokens("CONCAT"); return CONCAT }
    case 33:  //\(/
{ logDebugTokens("LPAREN"); return LPAREN }
    case 34:  //\)/
{ logDebugTokens("RPAREN"); return RPAREN }
    case 35:  //\{/
{ logDebugTokens("LBRACE"); return LBRACE }
    case 36:  //\}/
{ logDebugTokens("RBRACE"); return RBRACE }
    case 37:  //\,/
{ logDebugTokens("COMMA"); return COMMA }
    case 38:  //\:/
{ logDebugTokens("COLON"); return COLON }
    case 39:  //\[/
{ logDebugTokens("LBRACKET"); return LBRACKET }
    case 40:  //\]/
{ logDebugTokens("RBRACKET"); return RBRACKET }
    case 41:  //"/
{ logDebugTokens("DQUOTE"); return DQUOTE }
    case 42:  //[tT][rR][uU][eE]/
{ logDebugTokens("TRUE"); return TRUE}
    case 43:  //[fF][aA][lL][sS][eE]/
{ logDebugTokens("FALSE"); return FALSE}
    case 44:  //[nN][uU][lL][lL]/
{ logDebugTokens("NULL"); return NULL}
    case 45:  //-?([0-9]|[1-9][0-9]*)(\.[0-9][0-9]*)([eE][+\-]?[0-9][0-9]*)?/
{
                  // there are 2 separate rules for NUMBER
                  // instead of 1 with two optional components
                  // to differntiate it from plan INT
                    lval.f,_ = strconv.ParseFloat(yylex.Text(), 64);
                    logDebugTokens("%f", lval.f)
                    logDebugTokens("NUMBER"); 
                    return NUMBER 
                  }
    case 46:  //-?([0-9]|[1-9][0-9]*)(\.[0-9][0-9]*)?([eE][+\-]?[0-9][0-9]*)/
{ 
                    lval.f,_ = strconv.ParseFloat(yylex.Text(), 64);
                    logDebugTokens("%f", lval.f)
                    logDebugTokens("NUMBER"); 
                    return NUMBER 
                  }
    case 47:  //-?[0-9]|[1-9][0-9]*/
{ 
                    lval.n,_ = strconv.Atoi(yylex.Text());
                    logDebugTokens("INT"); 
                    return INT 
                  }
    case 48:  //[ \t\n]+/
{ logDebugTokens("WHITESPACE (count=%d)", len(yylex.Text())) /* eat up whitespace */ }
    case 49:  //[a-zA-Z_][a-zA-Z0-9\-_]*/
{ 
                    lval.s = yylex.Text();
                    logDebugTokens("%s", lval.s)
                    logDebugTokens("IDENTIFIER"); 
                    return IDENTIFIER 
                  }
    case 50:  //`((\\\")|(\\\\)|(\\\/)|(\\b)|(\\f)|(\\n)|(\\r)|(\\t)|(\\u[0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F])|[^`])+`/
{
                    //this rule allows for a wider range of identifiers by escaping them
                    lval.s = yylex.Text()[1:len(yylex.Text())-1]
                    logDebugTokens("%s", lval.s)
                    logDebugTokens("IDENTIFIER"); 
                    return IDENTIFIER 
                  }
    case 51:  ///
// [END]
    }
  }
  return 0
}
func logDebugTokens(format string, v ...interface{}) {
    if DebugTokens && len(v) > 0 {
        log.Printf("DEBUG TOKEN " + format, v)
    } else if DebugTokens {
        log.Printf("DEBUG TOKEN " + format)
    }
}