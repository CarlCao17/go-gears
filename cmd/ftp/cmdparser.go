package main

import (
	"io"

	"github.com/CarlCao17/go-gears/pkg/bufferpool"
)

type CmdType int

const (
	CmdCd CmdType = iota
	CmdList
	CmdMkDir
	CmdRmDir

	CmdFileSize
	CmdModTime
	CmdDelete
	CmdRename
	CmdGet
	CmdPut
	CmdAppend
)

type parser struct {
	p bufferpool.Bytes
}

func (p *parser) Parse(r io.Reader) {
	
}
