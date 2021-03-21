package tushare

import "math"

type Boll struct {
	n int
	k float64
}

//NewBoll(20, 2)
func NewBoll(n, k int) *Boll {
	return &Boll{n: n, k: float64(k)}
}
func (this *Boll) sma(lines []*ProBarData) float64 {
	s := len(lines)
	var sum float64 = 0
	for i := 0; i < s; i++ {
		sum += float64(lines[i].Close)
	}
	return sum / float64(s)
}
func (this *Boll) dma(lines []*ProBarData, ma float64) float64 {
	s := len(lines)
	//log.Println(s)
	var sum float64 = 0
	for i := 0; i < s; i++ {
		sum += (lines[i].Close - ma) * (lines[i].Close - ma)
	}
	return math.Sqrt(sum / float64(this.n-1))
}

func (this *Boll) Boll(lines []*ProBarData) {
	l := len(lines)
	if l < this.n {
		return
	}
	for i := l - 1; i > this.n-1; i-- {
		ps := lines[(i - this.n + 1) : i+1]
		lines[i].BOLL.MID = this.sma(ps)
		dm := this.k * this.dma(ps, lines[i].BOLL.MID)
		lines[i].BOLL.UP = lines[i].BOLL.MID + dm
		lines[i].BOLL.LOW = lines[i].BOLL.MID - dm
	}

	return
}
