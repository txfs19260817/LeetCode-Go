package generaterandomnft

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNFT(t *testing.T) {
	cases := []struct {
		name   string
		config Config
		n      int
	}{
		{
			name: "config-1",
			n:    5,
			config: Config{
				Name: "config-1",
				Size: "large",
				Traits: map[string][]TraitValue{
					"nose": {
						{Name: "pointy", Weight: 1},
						{Name: "tiny", Weight: 1},
						{Name: "flat", Weight: 1},
					},
					"mouth": {
						{Name: "small", Weight: 1},
						{Name: "wide", Weight: 1},
						{Name: "thin", Weight: 1},
					},
					"eyes": {
						{Name: "blue", Weight: 1},
						{Name: "green", Weight: 1},
						{Name: "brown", Weight: 1},
					},
				},
			},
		},
		{
			name: "config-2",
			n:    3,
			config: Config{
				Name: "config-2",
				Size: "small",
				Traits: map[string][]TraitValue{
					"color": {
						{Name: "red", Weight: 1},
						{Name: "blue", Weight: 1},
						{Name: "green", Weight: 1},
					},
					"shape": {
						{Name: "circle", Weight: 1},
						{Name: "square", Weight: 1},
					},
				},
			},
		},
		{
			name: "config-3",
			n:    3,
			config: Config{
				Name: "config-3",
				Size: "large",
				Traits: map[string][]TraitValue{
					"color": {
						{Name: "red", Weight: 1},
						{Name: "blue", Weight: 1},
						{Name: "green", Weight: 1},
						{Name: "yellow", Weight: 1},
						{Name: "purple", Weight: 1},
					},
					"texture": {
						{Name: "smooth", Weight: 1},
						{Name: "rough", Weight: 1},
						{Name: "grainy", Weight: 1},
					},
					"size": {
						{Name: "tiny", Weight: 1},
						{Name: "small", Weight: 1},
						{Name: "medium", Weight: 1},
						{Name: "large", Weight: 1},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			solution := Solution{}
			result := solution.GenerateNFT(tc.config, tc.n)
			assert.Len(t, result, tc.n)
			assertNFTsValid(t, result, tc.config.Traits)
			assertNFTsUnique(t, result)
		})
	}
}

func TestGenerateNFTWeighted(t *testing.T) {
	config := Config{
		Name: "config-weighted",
		Size: "large",
		Traits: map[string][]TraitValue{
			"nose": {
				{Name: "pointy", Weight: 1},
				{Name: "tiny", Weight: 2},
				{Name: "flat", Weight: 3},
			},
			"mouth": {
				{Name: "small", Weight: 1000},
				{Name: "wide", Weight: 1},
				{Name: "thin", Weight: 1},
			},
			"eyes": {
				{Name: "blue", Weight: 10},
				{Name: "green", Weight: 2},
				{Name: "brown", Weight: 1},
			},
		},
	}

	solution := Solution{}
	result := solution.GenerateNFT(config, 5)
	assert.Len(t, result, 5)
	assertNFTsValid(t, result, config.Traits)
	assertNFTsUnique(t, result)
}

func TestGenerateNFTTooMany(t *testing.T) {
	config := Config{
		Name: "simple",
		Size: "small",
		Traits: map[string][]TraitValue{
			"color": {
				{Name: "red", Weight: 1},
				{Name: "blue", Weight: 1},
				{Name: "green", Weight: 1},
			},
			"shape": {
				{Name: "circle", Weight: 1},
				{Name: "square", Weight: 1},
			},
		},
	}

	solution := Solution{}
	assert.Panics(t, func() {
		solution.GenerateNFT(config, 10)
	})
}

func assertNFTsValid(t *testing.T, results []map[string]string, traits map[string][]TraitValue) {
	allowed := make(map[string]map[string]struct{}, len(traits))
	for trait, values := range traits {
		set := make(map[string]struct{}, len(values))
		for _, value := range values {
			set[value.Name] = struct{}{}
		}
		allowed[trait] = set
	}

	for _, nft := range results {
		assert.Len(t, nft, len(traits))
		for trait, values := range allowed {
			value, ok := nft[trait]
			assert.True(t, ok)
			_, allowedValue := values[value]
			assert.True(t, allowedValue)
		}
		for trait := range nft {
			_, ok := allowed[trait]
			assert.True(t, ok)
		}
	}
}

func assertNFTsUnique(t *testing.T, results []map[string]string) {
	seen := make(map[string]struct{}, len(results))
	for _, nft := range results {
		serialized := serializeNFT(nft)
		_, exists := seen[serialized]
		assert.False(t, exists)
		seen[serialized] = struct{}{}
	}
}

func serializeNFT(nft map[string]string) string {
	keys := make([]string, 0, len(nft))
	for key := range nft {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	out := make([]byte, 0, len(keys)*12)
	for _, key := range keys {
		out = append(out, key...)
		out = append(out, '=')
		out = append(out, nft[key]...)
		out = append(out, ';')
	}
	return string(out)
}
