package procIO

import (
	"fmt"
	"io"
)

type Process interface {
	io.Closer
	fmt.Stringer

	PID() int
	Handle() interface{}
	MemorySegments() ([]*MemorySegmentInfo, error)
	Suspend() error
	Resume() error
}

type CachingProcess interface {
	Process
	InvalidateCache()
}

func OpenProcess(pid int) (CachingProcess, error) {
	proc, err := open(pid)
	return &cachingProcess{
		proc:  proc,
		cache: nil,
	}, err
}

type cachingProcess struct {
	proc  Process
	cache []*MemorySegmentInfo
}

func (c cachingProcess) Close() error {
	return c.proc.Close()
}

func (c cachingProcess) String() string {
	return c.proc.String()
}

func (c cachingProcess) PID() int {
	return c.proc.PID()
}

func (c cachingProcess) Handle() interface{} {
	return c.proc.Handle()
}

func (c cachingProcess) Suspend() error {
	return c.proc.Suspend()
}

func (c cachingProcess) Resume() error {
	return c.proc.Resume()
}

func (c cachingProcess) MemorySegments() ([]*MemorySegmentInfo, error) {
	var err error
	if c.cache == nil {
		c.cache, err = c.proc.MemorySegments()
	}
	return c.cache, err
}

func (c cachingProcess) InvalidateCache() {
	c.cache = nil
}
