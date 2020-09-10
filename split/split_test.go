package split

import (
	"testing"
)

func TestSplitMultiSep(t *testing.T) {
	s := "Python,JavaScript Twitter"
	result := SplitMultiSep(s, []string{"&",","," "})
	t.Logf("result is %v",result)
	for i := range result{
		t.Log(result[i])
	}
}