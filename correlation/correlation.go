package correlation

import (
	"fmt"
	"math"
)

type Vector struct {
	data []float64
}

func (v Vector) len() int {
	return len(v.data)
}

func (v Vector) get(i int) float64 {
	if i >= v.len() {
		fmt.Println(fmt.Sprintf("invalid index: %d", i))
		return 0
	}
	return v.data[i]
}

type Correlator interface {
	Correlate(Vector, Vector) float64
}

type Pearson struct {
}

var pow = func(n float64) float64 {
	return math.Pow(n, float64(2))
}

//皮尔逊相关性
func (p *Pearson) Correlate(v1, v2 Vector) float64 {
	if v1.len() != v2.len() {
		fmt.Println("v1 , v2 has diff length")
		return 0
	}
	var sum1, sum2, sumSq1, sumSq2 float64
	var pSum float64 //乘积之和
	size := v1.len()
	for i := 0; i < size;i++ {
		sum1 += v1.get(i)
		sumSq1 += pow(v1.get(i))
		sum2 += v2.get(i)
		sumSq2 += pow(v2.get(i))

		pSum += v1.get(i) * v2.get(i)
	}

	num := pSum - (sum1 * sum2 / float64(size))
	den := math.Sqrt((sumSq1 - pow(sum1)/float64(size)) * (sumSq2 - pow(sum2)/float64(size)))
	if den == 0 {
		return 0
	}

	return num/den
}
