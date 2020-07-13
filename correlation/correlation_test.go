package correlation

import (
	"fmt"
	"testing"
)

func TestPearson_Correlate(t *testing.T) {
	pearson := &Pearson{}
	r := pearson.Correlate(Vector{data: []float64{2, 1.2, 3, 1.2, 1.5}}, Vector{[]float64{2.2, 1.3, 2, 1.5, 2.5}})
	fmt.Println(r)

	r = pearson.Correlate(Vector{data: []float64{1,2,3}}, Vector{[]float64{3,2,1}})
	fmt.Println(r)
}
