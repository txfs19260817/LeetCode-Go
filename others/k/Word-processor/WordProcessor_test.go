package k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wrapLines(t *testing.T) {
	type args struct {
		words  []string
		maxLen int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "words1_maxLen13",
			args: args{
				words:  []string{"The", "day", "began", "as", "still", "as", "the", "night", "abruptly", "lighted", "with", "brilliant", "flame"},
				maxLen: 13,
			},
			wantAns: []string{"The-day-began", "as-still-as", "the-night", "abruptly", "lighted-with", "brilliant", "flame"},
		},
		{
			name: "words1_maxLen20",
			args: args{
				words:  []string{"The", "day", "began", "as", "still", "as", "the", "night", "abruptly", "lighted", "with", "brilliant", "flame"},
				maxLen: 20,
			},
			wantAns: []string{"The-day-began-as", "still-as-the-night", "abruptly-lighted", "with-brilliant-flame"},
		},
		{
			name: "words2_single_word",
			args: args{
				words:  []string{"Hello"},
				maxLen: 5,
			},
			wantAns: []string{"Hello"},
		},
		{
			name: "words3_two_words_each_line",
			args: args{
				words:  []string{"Hello", "world"},
				maxLen: 5,
			},
			wantAns: []string{"Hello", "world"},
		},
		{
			name: "words4_three_words_each_line",
			args: args{
				words:  []string{"Well", "Hello", "world"},
				maxLen: 5,
			},
			wantAns: []string{"Well", "Hello", "world"},
		},
		{
			name: "words5_mixed_lengths",
			args: args{
				words:  []string{"Hello", "HelloWorld", "Hello", "Hello"},
				maxLen: 20,
			},
			wantAns: []string{"Hello-HelloWorld", "Hello-Hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantAns, wrapLines(tt.args.words, tt.args.maxLen))
		})
	}
}

func Test_reflowAndJustify(t *testing.T) {
	type args struct {
		lines  []string
		maxLen int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "example_width_24",
			args: args{
				lines: []string{
					"The day began as still as the",
					"night abruptly lighted with",
					"brilliant flame",
				},
				maxLen: 24,
			},
			wantAns: []string{
				"The--day--began-as-still",
				"as--the--night--abruptly",
				"lighted--with--brilliant",
				"flame",
			},
		},
		{
			name: "example_width_25",
			args: args{
				lines: []string{
					"The day began as still as the",
					"night abruptly lighted with",
					"brilliant flame",
				},
				maxLen: 25,
			},
			wantAns: []string{
				"The-day-began-as-still-as",
				"the-----night----abruptly",
				"lighted---with--brilliant",
				"flame",
			},
		},
		{
			name: "example_width_26",
			args: args{
				lines: []string{
					"The day began as still as the",
					"night abruptly lighted with",
					"brilliant flame",
				},
				maxLen: 26,
			},
			wantAns: []string{
				"The--day-began-as-still-as",
				"the-night-abruptly-lighted",
				"with----brilliant----flame",
			},
		},
		{
			name: "example_width_40",
			args: args{
				lines: []string{
					"The day began as still as the",
					"night abruptly lighted with",
					"brilliant flame",
				},
				maxLen: 40,
			},
			wantAns: []string{
				"The--day--began--as--still--as-the-night",
				"abruptly--lighted--with--brilliant-flame",
			},
		},
		{
			name: "example_width_14",
			args: args{
				lines: []string{
					"The day began as still as the",
					"night abruptly lighted with",
					"brilliant flame",
				},
				maxLen: 14,
			},
			wantAns: []string{
				"The--day-began",
				"as---still--as",
				"the------night",
				"abruptly",
				"lighted---with",
				"brilliant",
				"flame",
			},
		},
		{
			name: "all_words_fit_in_one_line_exact",
			args: args{
				lines: []string{
					"The day began as still as the",
					"night abruptly lighted with",
					"brilliant flame",
				},
				// sum(len(word)) + (numWords - 1) = 61 + 12 = 73
				maxLen: 73,
			},
			wantAns: []string{
				"The-day-began-as-still-as-the-night-abruptly-lighted-with-brilliant-flame",
			},
		},
		{
			name: "single_word_not_padded_even_if_maxLen_bigger",
			args: args{
				lines:  []string{"flame"},
				maxLen: 10,
			},
			// 题目要求：单词独占一行时，不做 padding
			wantAns: []string{
				"flame",
			},
		},
		{
			name: "single_line_two_words_should_be_justified",
			args: args{
				lines:  []string{"hello world"},
				maxLen: 13,
			},
			// "hello" + 3 * "-" + "world" => 5 + 3 + 5 = 13
			wantAns: []string{
				"hello---world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantAns, reflowAndJustify(tt.args.lines, tt.args.maxLen))
		})
	}
}

func Test_balancedWrapLines(t *testing.T) {
	type args struct {
		text   string
		maxLen int
	}
	tests := []struct {
		name      string
		args      args
		wantLines []string
	}{
		// text0 from prompt
		{
			name: "text0_example_max7",
			args: args{
				text:   "one two add a three four",
				maxLen: 7,
			},
			// 最优 score = 9，题目给的正确答案
			wantLines: []string{
				"one",
				"two",
				"add-a",
				"three",
				"four",
			},
		},
		// text1 from prompt
		{
			name: "text1_example_max15",
			args: args{
				text:   "Seven six eight thirteen four five sixteen",
				maxLen: 15,
			},
			// 最优 score = 13，题目给的示例
			wantLines: []string{
				"Seven-six-eight",
				"thirteen-four",
				"five-sixteen",
			},
		},
		// text2 from prompt
		{
			name: "text2_example_max9",
			args: args{
				text:   "XXXXXX X XXXX X X XXXXX",
				maxLen: 9,
			},
			// 最优 score = 5，这个是题目给的两种正确答案之一
			wantLines: []string{
				"XXXXXX",
				"X-XXXX-X",
				"X-XXXXX",
			},
		},
		// ---- 扩展 / 边界用例 ----

		// 单个单词，肯定就一行
		{
			name: "single_word_fits_in_line",
			args: args{
				text:   "Hello",
				maxLen: 10,
			},
			wantLines: []string{
				"Hello",
			},
		},
		// 两个单词合起来超过 maxLen，只能拆两行，score = 0
		{
			name: "two_words_each_on_own_line",
			args: args{
				text:   "Hello world",
				maxLen: 5,
			},
			wantLines: []string{
				"Hello",
				"world",
			},
		},
		// 三个短词，最佳是每个一行（更平衡）
		{
			name: "three_short_words_each_line",
			args: args{
				text:   "aa bbb ccc",
				maxLen: 7,
			},
			// 可能的划分：
			// ["aa","bbb","ccc"]    -> 长度 [2,3,3], M=3, score=1
			// ["aa-bbb","ccc"]      -> [6,3],       M=6, score=9
			// ["aa","bbb-ccc"]      -> [2,7],       M=7, score=25
			// 最优是每个一行
			wantLines: []string{
				"aa",
				"bbb",
				"ccc",
			},
		},
		// 稍微复杂一点的平衡例子
		{
			name: "balanced_example_custom",
			args: args{
				text:   "a bb c ddd eee",
				maxLen: 7,
			},
			// 最优划分：
			// ["a-bb-c", "ddd-eee"]
			// 长度 [7,7], M=7, score=0
			wantLines: []string{
				"a-bb-c",
				"ddd-eee",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantLines, balancedWrapLines(tt.args.text, tt.args.maxLen))
		})
	}
}
