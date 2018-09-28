package kata

import "bytes"

func Accum(s string) string {
    // your code
    var buf bytes.Buffer
    for i:=0;i<len(s);i++ {
      if i > 0 {
        buf.WriteString("-")
      }
      var ch rune
      if s[i] >= 'a' {
        ch = rune(s[i] - ('a'-'A'))
      } else {
        ch = rune(s[i])
      }
      for j:=0;j<=i;j++ {
        if j == 0 {
          buf.WriteRune(ch)
        } else {
          buf.WriteRune(ch+('a'-'A'))
        }
      }
    }
    return buf.String()
}
