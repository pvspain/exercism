package raindrops

import "fmt"

type Factors struct {
	three, five, seven, n int
}

func Convert(n int) string {
	return _convert(Factors{n % 3, n % 5, n % 7, n})
}

func _convert(f Factors) string {
  result := ""

  if f.n > 0 {
    if f.three == 0 {
      result += "Pling"
    }
    if f.five == 0 {
      result += "Plang"
    }
    if f.seven == 0 {
      result += "Plong"
    }
  }

  if result == "" {
    result = fmt.Sprintf("%d", f.n)
  }
  return result
}
