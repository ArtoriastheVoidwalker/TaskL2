package main

import "testing"

func Test_unpacking(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    string
		wantErr bool
	}{
		{
			name: "Test №1",
			s:    "a4bc2d5e",
			want: "aaaabccddddde",
		},
		{
			name: "Test №2",
			s:    "abcd",
			want: "abcd",
		},
		{
			name:    "Test №3",
			s:       "45",
			want:    "",
			wantErr: true,
		},
		{
			name: "Test №4",
			s:    "",
			want: "",
		},
		{
			name: "Test №5",
			s:    `qwe\4\5`,
			want: "qwe45",
		},
		{
			name: "Test №6",
			s:    `qwe\45`,
			want: "qwe44444",
		},
		{
			name: "Test №7",
			s:    `qwe\\5`,
			want: `qwe\\\\\`,
		},
		{
			name: "Test №8",
			s:    `кыв0ер\\в4`,
			want: `кыер\вввв`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpacking(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpacking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unpacking() got = %v, want %v", got, tt.want)
			}
		})
	}
}
