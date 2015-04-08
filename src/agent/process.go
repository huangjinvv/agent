package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var Process *Proc = NewProc()

type Proc struct {
	BaseInfo *ProcBase
	Cpu      *ProcCpu
	Mem      *ProcMem
}

func NewProc() *Proc {
	return &Proc{}
}

func (this *Proc) Init() {
	this.BaseInfo = &ProcBase{}
	this.Cpu = &ProcCpu{}
	this.Mem = &ProcMem{}
	this.BaseInfo.GetProcInfo()
	this.Mem.Update()
	this.Cpu.Update()
}

type ProcMem struct {
	VmSize int
	VmRss  int
	VmData int
	VmStk  int
	VmExe  int
	VmLib  int
}

func (this *ProcMem) Update() error {
	file, err := os.Open("/proc/" + Process.BaseInfo.Pid + "/status")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fields := strings.Fields(line)
		if len(fields) != 3 {
			continue
		}
		if strings.Trim(fields[1], " ") == "0" {
			continue
		}
		switch strings.Trim(fields[0], ":") {
		case "VmSize":
			this.VmSize, _ = strconv.Atoi(fields[1])
		case "VmRSS":
			this.VmRss, _ = strconv.Atoi(fields[1])
		case "VmData":
			this.VmData, _ = strconv.Atoi(fields[1])
		case "VmStk":
			this.VmStk, _ = strconv.Atoi(fields[1])
		case "VmExe":
			this.VmExe, _ = strconv.Atoi(fields[1])
		case "VmLib":
			this.VmLib, _ = strconv.Atoi(fields[1])
		}
	}
	return nil
}

func (this *ProcMem) ReSet() {
	this.VmData = 0
	this.VmSize = 0
	this.VmRss = 0
	this.VmLib = 0
	this.VmExe = 0
	this.VmStk = 0
}

func (this *ProcMem) String() string {
	return fmt.Sprintf("VIRT:%d KB, RES:%d KB, Data:%d KB, Stack:%d KB, Text Segment:%d KB, Lib:%d KB",
		this.VmSize, this.VmRss, this.VmData, this.VmStk, this.VmExe, this.VmLib)
}

type ProcCpu struct {
	Utime     uint64
	Stime     uint64
	Cutime    uint64
	Cstime    uint64
	StartTime uint64
}

func (this *ProcCpu) Update() error {
	file, err := os.Open("/proc/" + Process.BaseInfo.Pid + "/stat")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil
	}
	fields := strings.Fields(line)
	if utime, err := strconv.ParseUint(fields[13], 10, 64); err == nil {
		this.Utime = utime
	}
	if stime, err := strconv.ParseUint(fields[14], 10, 64); err == nil {
		this.Stime = stime
	}
	if cutime, err := strconv.ParseUint(fields[15], 10, 64); err == nil {
		this.Cutime = cutime
	}
	if cstime, err := strconv.ParseUint(fields[16], 10, 64); err == nil {
		this.Cstime = cstime
	}
	if starttime, err := strconv.ParseUint(fields[21], 10, 64); err == nil {
		this.StartTime = starttime
	}
	return nil
}

func (this *ProcCpu) ReSet() {
	this.Utime = 0
	this.Stime = 0
	this.Cutime = 0
	this.Cstime = 0
	this.StartTime = 0
}

func (this *ProcCpu) String() string {
	this.Update()
	totalTime := this.Utime + this.Stime
	uptime := Machine.GetUptime()
	seconds := uptime - float64(this.StartTime)/float64(Machine.Hertz)
	pcpu := (float64(totalTime) * 1000 / float64(Machine.Hertz)) / seconds / 10
	return fmt.Sprintf("Cpu:%0.2f%%", pcpu)
}

type ProcBase struct {
	Pid     string
	PPid    string
	Command string
	State   string
}

func (this *ProcBase) GetProcInfo() error {
	this.Pid = strconv.Itoa(os.Getpid())
	file, err := os.Open("/proc/" + this.Pid + "/stat")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil
	}
	fields := strings.Fields(line)
	this.PPid = fields[3]
	this.Command = this.GetCommand()
	this.State = fields[2]
	return nil
}

func (this *ProcBase) GetCommand() string {
	command, _ := ioutil.ReadFile("/proc/" + this.Pid + "/cmdline")
	return string(command)
}
