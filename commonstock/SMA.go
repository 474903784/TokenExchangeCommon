package commonstock

import (
	"github.com/idoall/TokenExchangeCommon/commonmodels"
	"time"
)

// SMA struct
type SMA struct {
	Period int //默认计算几天的SMA
	alpha  int //权重，一般取作1
	points []SMAPoint
	kline  []*commonmodels.Kline
}

type SMAPoint struct {
	point
}

// NewSMA new Func
func NewSMA(list []*commonmodels.Kline, period int) *SMA {
	m := &SMA{kline: list, Period: period, alpha: 1}
	return m
}

// Calculation Func
func (e *SMA) Calculation() *SMA {
	for _, v := range e.kline {
		e.add(v.KlineTime, v.Close)
	}
	return e
}

// GetPoints Func
func (e *SMA) GetPoints() []SMAPoint {
	return e.points
}

// Add adds a new Value to Sma
// 使用方法，先添加最早日期的数据,最后一条应该是当前日期的数据
func (e *SMA) add(timestamp time.Time, value float64) {
	p := SMAPoint{}
	p.Time = timestamp

	smaTminusOne := value
	if len(e.points) > 0 {
		smaTminusOne = e.points[len(e.points)-1].Value
	}

	// 计算 SMA指数
	emaT := (float64(e.alpha)*value + float64(e.Period-e.alpha)*smaTminusOne) / float64(e.Period)
	p.Value = emaT
	e.points = append(e.points, p)
}
