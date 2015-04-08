package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type MachineCpu struct {
	User        uint64
	Nice        uint64
	System      uint64
	Idle        uint64
	Iowait      uint64
	Irq         uint64
	SoftIrq     uint64
	Stealstolen uint64
	Guest       uint64
}

func (this *MachineCpu) Update() error {
	file, err := os.Open("/proc/stat")
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

	if user, err := strconv.ParseUint(fields[1], 10, 64); err == nil {
		this.User = user
	}
	if nice, err := strconv.ParseUint(fields[2], 10, 64); err == nil {
		this.Nice = nice
	}
	if system, err := strconv.ParseUint(fields[3], 10, 64); err == nil {
		this.System = system
	}
	if idle, err := strconv.ParseUint(fields[4], 10, 64); err == nil {
		this.Idle = idle
	}
	if iowait, err := strconv.ParseUint(fields[5], 10, 64); err == nil {
		this.Iowait = iowait
	}
	if irq, err := strconv.ParseUint(fields[6], 10, 64); err == nil {
		this.Irq = irq
	}
	if softirq, err := strconv.ParseUint(fields[7], 10, 64); err == nil {
		this.SoftIrq = softirq
	}
	if stealstolen, err := strconv.ParseUint(fields[8], 10, 64); err == nil {
		this.Stealstolen = stealstolen
	}
	if guest, err := strconv.ParseUint(fields[9], 10, 64); err == nil {
		this.Guest = guest
	}
	return nil
}

func (this *MachineCpu) ReSet() {
	this.User = 0
	this.Nice = 0
	this.System = 0
	this.Idle = 0
	this.Iowait = 0
	this.Irq = 0
	this.SoftIrq = 0
	this.Stealstolen = 0
	this.Guest = 0
}

var Machine *MachineInfo = NewMachineInfo()

type MachineInfo struct {
	Uptime float64
	Hertz  int
	Cpu    *MachineCpu
}

func NewMachineInfo() *MachineInfo {
	return &MachineInfo{Hertz: 100}
}

func (this *MachineInfo) GetUptime() float64 {
	if uptime, err := ioutil.ReadFile("/proc/uptime"); err == nil {
		fields := strings.Fields(string(uptime))
		this.Uptime, _ = strconv.ParseFloat(fields[0], 64)
	}
	return this.Uptime
}
