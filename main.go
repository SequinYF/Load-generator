package Load_generator

import (
	"context"
	"github.com/lunny/log"
	"time"
)
import "./lglib"

type LoadGen struct {
	lps uint32 //每秒载荷量
	durationNS time.Duration //载荷持续时间
	timeoutNS time.Duration //处理超时时间

	resultCh chan *lglib.CallResult

	ctx context.Context
	cancelFunc context.CancelFunc

	caller lglib.Caller
	tickets lglib.Gotickets

	concurrency uint32 //载荷并发量
	callcount int64
	status uint32
}
//caller lglib.Caller, lps uint32, durationNS, timeoutNS time.Duration, resultCh chan *lglib.CallResult

func NewLoadGen(pset lglib.ParamSet) (lglib.LoadGenerator, error) {
	log.Info("new a load generator...")
	if err := pset.Check(); err != nil {
		return nil, err
	}
	gen := &LoadGen{
		lps:pset.Lps,
		caller:pset.Caller,
		durationNS:pset.DurationNS,
		timeoutNS:pset.TimeoutNS,
		resultCh:pset.ResultCh,
		status:lglib.STATUS_ORIGINAL,
	}
	if err := gen.init(); err != nil {
		return nil, err
	}

	return gen ,nil
}

func (gen *LoadGen)Start() bool {
	return true
}
func (gen *LoadGen)Stop() bool {
	return true
}
func (gen *LoadGen)Status() uint32 {
	return 1
}
func (gen *LoadGen)Callcount() int64 {
	return 1
}