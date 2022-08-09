package sol

import "testing"

func BenchmarkTestXOR(b *testing.B) {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	for idx := 0; idx < b.N; idx++ {
		missingNumberXOR(nums)
	}
}
func Test_missingNumberXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [3,0,1]",
			args: args{nums: []int{3, 0, 1}},
			want: 2,
		},
		{
			name: "nums = [0,1]",
			args: args{nums: []int{0, 1}},
			want: 2,
		},
		{
			name: "nums = [9,6,4,2,3,5,7,0,1]",
			args: args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumberXOR(tt.args.nums); got != tt.want {
				t.Errorf("missingNumberXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
