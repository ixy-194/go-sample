package fizzbuzz_test

import (
	"go-sample/fizzbuzz"
	"testing"
)

// テストの定義
func TestFizzBuzz(t *testing.T) {
	// テストケースの定義。input が入力値、expected が期待される出力です。
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},         // 1を入力した場合、出力は "1"
		{3, "Fizz"},      // 3を入力した場合、出力は "Fizz"
		{5, "Buzz"},      // 5を入力した場合、出力は "Buzz"
		{15, "FizzBuzz"}, // 15を入力した場合、出力は "FizzBuzz"
		{7, "7"},         // 7を入力した場合、出力は "7"
	}

	// 各テストケースをループで回します。
	for _, test := range tests {
		// fizzbuzz パッケージの FizzBuzz 関数を呼び出し、結果を result に格納します。
		result := fizzbuzz.FizzBuzz(test.input)

		// 結果が期待される出力と異なる場合、エラーメッセージを表示します。
		if result != test.expected {
			t.Errorf("FizzBuzz(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
