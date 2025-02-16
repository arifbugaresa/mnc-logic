package logic

import "testing"

func TestIsValidBracketString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "testcase 1",
			args: "{{[<>[{{}}]]}}",
			want: true,
		},
		{
			name: "testcase 2",
			args: "{<{[[{{[]<{{[{[]<>}]}}<>>}}]]}>}",
			want: true,
		},
		{
			name: "testcase 3",
			args: "{{[{<[[{<{<<<[{{{[]{<{[<[[<{{[[[[[<{[{<[<<[[<<{[[{[<<<<<<<[{[{[{{<{[[<{<<<{<{[<>]}>}>>[]>}>]]}>}}]}]}]>>>>>>>]}]]}>>]]>>]>}]}>]]]]]}}>]]>]}>}}}}]>>>}>}]]>}]}}",
			want: true,
		},
		{
			name: "testcase 4",
			args: "[<{<{[{[{}[[<[<{{[<[<[[[<{{[<<<[[[<[<{{[<<{{<{<{<[<{[{{[{{{{[<<{{{<{[{[[[{<<<[{[<{<<<>>>}>]}]>>>}]]]}]}>}}}>>]}}}}]}}]}>]>}>}>}}>>]}}>]>]]]>>>]}}>]]]>]>]}}>]>]]]}]}>}>]",
			want: true,
		},
		{
			name: "testcase 5",
			args: "[[{[[<{{{{[[[<[{[[<{{{{[{[{[[[<<{<{[{<<<[[<[{[<{{[\n{[<[[<<[{<<[[[{<[{[[{{<<>[<<{{<<{[[[<{}{[{{{[[{{[[<[{}]>]]}}]]}}}]}>]]]}>>}}>>]>}}]]}]>}]]]>>}]>>]]>]}]}}>]}]>]]>>>}]}>}>>]]]}]}]}}}}>]]}]>]]]}}}}>]]}]]",
			want: true,
		},
		{
			name: "testcase 6",
			args: "[{}<>]",
			want: true,
		},
		{
			name: "testcase 7",
			args: "]",
			want: false,
		},
		{
			name: "testcase 8",
			args: "][",
			want: false,
		},
		{
			name: "testcase 9",
			args: "[>]",
			want: false,
		},
		{
			name: "testcase 10",
			args: "[>",
			want: false,
		},
		{
			name: "testcase 11",
			args: "{<{[[{{[]<{[{[]<>}]}}<>>}}]]}>}",
			want: false,
		},
		{
			name: "testcase 12",
			args: "{{[{<[[{<{<<<[{{{[]{<{[<[[<{{[[[[<{[{<[<<[[<<{[[{[<<<<<<<[{[{[{{<{[[<{<<<{<{[<>]}>}>>[]>}>]]}",
			want: false,
		},
		{
			name: "testcase 13",
			args: ">}}]}]}]>>>>>>]}]]}>>]]>>]>}]}>]]]]]}}>]]>]}>}}}}]>>>}>}]]>}]}}",
			want: false,
		},
		{
			name: "testcase 14",
			args: "[<{<{[{[{}[[<[<{{[<[<[[[<{{[<<<[[[<[<{{[<<{{<{<{<[<{[{{[{{{{[<<{{{<{[{[[[{<<<[{[<{<<>>[]}]>>>}]]]}]}>}}}>>]}}}}]}}]}>]>}>}>}}>>]}}>]>]]]>>>]}}>]]]>]>]}}>]>]]]}]}>}>]",
			want: false,
		},
		{
			name: "testcase 15",
			args: "[{}<[>]",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBracketString(tt.args); got != tt.want {
				t.Errorf("IsValidBracketString() = %v, want %v", got, tt.want)
			}
		})
	}
}
