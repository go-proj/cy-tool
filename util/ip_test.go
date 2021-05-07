package util

import (
	"testing"
)

func TestSvcIpip(t *testing.T) {
	tests := []struct {
		name string
		param string
		want string
	}{
		{"local_ip", "", "当前 IP：58.246.102.65  来自于：中国 上海 上海  联通"},
		{"global_ip", "58.246.102.133", "[\"中国\",\"上海\",\"上海\",\"\",\"\"]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SvcIpip(tt.param); got != tt.want {
				t.Errorf("SvcIpip('%v') = %v, want %v", tt.param, got, tt.want)
			} else {
				t.Logf("SvcIpip('%v') = %v, want %v", tt.param, got, tt.want)
			}
		})
	}
}

func TestSvcIpInfo(t *testing.T) {
	tests := []struct {
		name string
		param string
		want string
	}{
		{"local_ip", "", "当前 IP：58.246.102.65  来自于：中国 上海 上海  联通"},
		{"global_ip", "58.246.102.133", "[\"中国\",\"上海\",\"上海\",\"\",\"\"]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SvcIpInfo(tt.param); got != tt.want {
				t.Errorf("SvcIpInfo('%v') = %v, want %v", tt.param, got, tt.want)
			} else {
				t.Logf("SvcIpInfo('%v') = %v, want %v", tt.param, got, tt.want)
			}
		})
	}
}

func TestSvcIpFm(t *testing.T) {
	tests := []struct {
		name string
		param string
		want string
	}{
		{"local_ip", "", "当前 IP：58.246.102.65  来自于：中国 上海 上海  联通"},
		{"global_ip", "58.246.102.133", "[\"中国\",\"上海\",\"上海\",\"\",\"\"]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SvcIpFm(tt.param); got != tt.want {
				t.Errorf("SvcIpFm('%v') = %v, want %v", tt.param, got, tt.want)
			} else {
				t.Logf("SvcIpFm('%v') = %v, want %v", tt.param, got, tt.want)
			}
		})
	}
}
