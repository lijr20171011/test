package main

import (
	"errors"
	"flag"
	"fmt"
	"my_project/test/utils"
	"strings"
	"time"
)

type interval []time.Duration

//实现String接口
func (i *interval) String() string {
	// utils.Info("String")
	// return fmt.Sprintf("%v", *i)
	return fmt.Sprintf("%v", *i)
}

//实现Set接口,Set接口决定了如何解析flag的值
func (i *interval) Set(value string) error {
	utils.Info("Set")
	//此处决定命令行是否可以设置多次-deltaT
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

var intervalFlag interval

func init() {
	utils.Info("init")
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
	utils.Info("main")
	flag.Parse()
	// utils.Info("parse")
	// fmt.Println(intervalFlag)
	// utils.Info("intervalflag")
}
