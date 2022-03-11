package async

import (
	"context"
	"errors"
	"reflect"
)

type runner int

// PipeLine defined TODO
type PipeLine struct {
	fsm    *simpleFSM
	ctx    context.Context
	cancel context.CancelFunc
	runner chan runner
	pipe   chan interface{}
}

func NewPipeLine(ctx context.Context, worker int) *PipeLine {
	x, c := context.WithCancel(ctx)
	line := PipeLine{pipe: make(chan interface{}, 1)}
	line.ctx = x
	line.runner = make(chan runner, worker)
	line.cancel = c
	line.fsm = newSimpleFSM()
	line.fsm.actEvent(running)
	return &line
}

// Product defined TODO
func (p *PipeLine) Product(i interface{}) error {
	iValue := reflect.ValueOf(i)
	iType := iValue.Type()
	ikind := iType.Kind()

	if p.isFinished() {
		return errors.New("isFinished")
	}

	if ikind == reflect.Slice || ikind == reflect.Array {
		for i := 0; i < iValue.Len(); i++ {
			p.pipe <- iValue.Index(i).Interface()
		}
	} else {
		p.pipe <- i
	}
	return nil
}

// ProductUntil defined TODO
func (p *PipeLine) ProductUntil(funk interface{}, init interface{}) error {
	funcValue := reflect.ValueOf(funk)
	initValue := reflect.ValueOf(init)
	funcType := funcValue.Type()
	if funcType.NumIn() != 1 || funcType.NumOut() != 2 {
		panic("invalid funk")
	}
	if funcType.Out(1).Kind() != reflect.Bool {
		panic("invalid funk")
	}
	result := funcValue.Call([]reflect.Value{initValue})
	i, b := result[0], result[1]

	if err := p.Product(i.Interface()); err != nil {
		return err
	}
	if !b.Interface().(bool) {
		p.ProductUntil(funk, init)
	}
	return nil
}

// Consumer defined TODO
func (p *PipeLine) Consumer(funk interface{}) {
	funcValue := reflect.ValueOf(funk)
	funcType := funcValue.Type()
	if funcType.NumIn() != 1 {
		panic("invalid funk")
	}
	go func() {
		for v := range p.pipe {
			p.dispatch(funcValue, reflect.ValueOf(v))
		}
		p.cancel()
	}()
}

func (p *PipeLine) dispatch(funk reflect.Value, v reflect.Value) {
	p.runner <- runner(0)
	go func() {
		defer func() {
			<-p.runner
		}()
		funk.Call([]reflect.Value{v})
	}()
}

func (p *PipeLine) isFinished() bool {
	return p.fsm.Current() >= shutdown
}

func (p *PipeLine) done() {
	if !p.isFinished() {
		p.fsm.actEvent(shutdown)
		close(p.pipe)
	}
}

func (p *PipeLine) Finish() {
	p.done()
}

func (p *PipeLine) Close() {
	p.done()
}

func (p *PipeLine) Wait() error {
	defer close(p.runner)
	if p.isFinished() {
		return nil
	}
	p.done()
	for range p.ctx.Done() {
		return p.ctx.Err()
	}
	return nil
}
