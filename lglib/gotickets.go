package lglib

import (
	"errors"
	"fmt"
)

type Gotickets interface {
	Take() bool
	Return() bool
	Active() bool
	Total() uint32
	Remainder() uint32
}

type GoTickets struct {
	total uint32
	ticketCh chan struct{}
	active bool
}

func (gt *GoTickets)Take() bool {
	return true
}

func (gt *GoTickets)Return() bool {
	return true
}
func (gt *GoTickets)Active() bool {
	return true
}
func (gt *GoTickets)Total() uint32 {
	return 1
}
func (gt *GoTickets)Remainder() uint32 {
	return 1
}

func (gt *GoTickets)init(total uint32) bool {
	if gt.active == true {
		return false
	}
	if total == 0 {
		return  false
	}

	num := int(total)
	ch := make(chan struct{}, total)
	for i := 0; i < num ; i++ {
		ch <- struct{}{}
	}
	gt.total = total
	gt.active = true
	gt.ticketCh = ch
	return true
}

func NewGoTickets(total uint32) (Gotickets, error) {
	gt := GoTickets{}
	if !gt.init(total) {
		errMsg := fmt.Sprintf("init a Go tickets error with total %d\n", total)
		return nil, errors.New(errMsg)
	}
	return &gt, nil
}


