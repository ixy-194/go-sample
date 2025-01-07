package fizzbuzz

import (
	"fmt"
)

// FizzBuzz は整数 n を受け取り、以下のルールに従って文字列を返します。
// - n が 3 で割り切れる場合は "Fizz" を返します。
// - n が 5 で割り切れる場合は "Buzz" を返します。
// - 上記のいずれにも該当しない場合は、n をそのまま文字列に変換して返します。
func FizzBuzz(n int) string {
	switch {
	// - n が 3 で割り切れる場合は "Fizz" を返します。
	case n%3 == 0:
		return "Fizz"
	// - n が 5 で割り切れる場合は "Buzz" を返します。
	case n%5 == 0:
		return "Buzz"
	// - 上記のいずれにも該当しない場合は、n をそのまま文字列に変換して返します。
	default:
		return fmt.Sprintf("%d", n)
	}
}
