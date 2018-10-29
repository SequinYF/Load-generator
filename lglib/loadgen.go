package lglib

import (
	"time"
)

const (
	STATUS_ORIGINAL uint32 = 0
	STATUS_STARTING uint32 = 1
	STATUS_STARTED uint32 = 2
	STATUS_STOPPINT uint32 = 3
	STATUS_STOPPED uint32 = 4
)


type Caller interface {
	BuildReq() RawReq
	Call(req RawReq, timeoutNs time.Duration)([]byte, error)
	CheckResp(rawReq RawReq, rawrResp RawResp) *CallResult
}

type LoadGenerator interface {
	Start() bool
	Stop() bool
	Status() uint32
	Callcount() int64
}


type CallResult struct {
	ID int64
	Req RawReq
	Resp RawResp
	//Code RetCode
	Status uint32
	Msg string
	Elapse time.Duration
}

type RawReq struct {
	ID int
	Req []byte
}

type RawResp struct {
	ID int
	Resp []byte
	Err error
	Elapse time.Duration
}

type ParamSet struct {
	Lps uint32 //每秒载荷量
	DurationNS time.Duration //载荷持续时间
	TimeoutNS time.Duration //处理超时时间
	ResultCh chan *CallResult
	Caller Caller
}

func (ps *ParamSet) Check() error {
	return nil
}

