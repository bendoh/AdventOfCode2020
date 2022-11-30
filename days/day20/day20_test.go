package day20

import "testing"
import "github.com/stretchr/testify/assert"

/*
  5   2
1 # . . 4
  . # #
3 # # . 6
  5   2
*/

func TestMakeTile(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		lines  []string
		result tile
	}{
		{
			lines: []string{
				"#..",
				".##",
				"##.",
			},
			result: tile{
				borders: [8]int{4, 2, 3, 5, 5, 6, 2, 1},
				body:    body{[]bool{true}},
			},
		},
		{
			lines: []string{
				"##.",
				"...",
				".##",
			},
			result: tile{
				borders: [8]int{6, 1, 6, 1, 4, 3, 4, 3},
				body:    body{[]bool{false}},
			},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(testCase.result, MakeTile(testCase.lines))
	}

}
