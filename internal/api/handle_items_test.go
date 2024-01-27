package api

import "testing"

func TestValidateItemName(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"regular name": {
			input: "Macbook Pro",
			want:  true,
		},
		"empty name": {
			input: "",
			want:  false,
		},
		"over char limit name": {
			input: "Factory New Macbook Pro NEW BEST CONDITION M2 chip APPLE BEST DEAL BEST DEAL",
			want:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := validateItemName(test.input); got != test.want {
				t.Fatalf("|TEST: %20s| got: %5t, want: %5t", name, got, test.want)
			}
		})
	}
}

func TestValidateItemDescription(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"regular description": {
			input: "Factory new, M2 chip, Apple",
			want:  true,
		},
		"empty description": {
			input: "",
			want:  false,
		},
		"over char limit description": {
			input: `at tellus at urna condimentum mattis pellentesque id nibh tortor id 
				aliquet lectus proin nibh nisl condimentum id venenatis a condimentum vitae 
				sapien pellentesque habitant morbi tristique senectus et netus et malesuada 
				fames ac turpis egestas sed tempus urna et pharetra pharetra massa massa 
				ultricies mi quis hendrerit dolor magna eget est lorem ipsum dolor sit amet
				consectetur adipiscing elit pellentesque habitant morbi tristique senectus et 
				netus et malesuada fames ac turpis egestas integer eget aliquet nibh praesent 
				tristique magna sit amet purus gravida quis blandit turpis cursus in hac 
				habitasse platea dictumst quisque sagittis purus sit amet volutpat consequat
				aliquet lectus proin nibh nisl condimentum id venenatis a condimentum vitae 
				sapien pellentesque habitant morbi tristique senectus et netus et malesuada 
				fames ac turpis egestas sed tempus urna et pharetra pharetra massa massa`,
			want: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := validateItemDescription(test.input); got != test.want {
				t.Fatalf("|TEST: %20s| got: %5t, want: %5t", name, got, test.want)
			}
		})
	}
}
