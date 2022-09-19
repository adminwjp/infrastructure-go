package utils

import (
	"log"
	"os"
	"strings"
	"time"
)

type TimerFlag int
const  (
	Min TimerFlag=iota
	Hour
	Day
	Week
	Month
	Year
)

type TimerLogConfig struct {
	Flag TimerFlag
	Value int
}
func TimerLogFile(config TimerLogConfig)  {
	file:=LogFile(config.Flag)
	//if file!=nil{
	//	defer file.Close()
	//}
	log.Println("splider test")
	//return

	go func() {

		t:=time.Now()
		create:=func (){
			if file!=nil{
				file.Close()
			}
			file=LogFile(config.Flag)
			t=time.Now()
		}
		for  {
			switch config.Flag {
			case Min:
				if t.Add(time.Minute*time.Duration(config.Value)).Unix()<time.Now().Unix(){
					create()
				}
				break
			case Hour:
				if t.Add(time.Hour*time.Duration(config.Value)).Unix()<time.Now().Unix(){
					create()
				}
				break
			case Day:
				if t.Add(time.Hour*time.Duration(config.Value*24)).Unix()<time.Now().Unix(){
					create()
				}
				break
			case Week:
				if t.Add(time.Hour*time.Duration(config.Value*24*7)).Unix()<time.Now().Unix(){
					create()
				}
				break
			case Month:
				if t.Add(time.Hour*time.Duration(config.Value*24*7*30)).Unix()<time.Now().Unix(){
					create()
				}
				break
			case Year:
				if t.Add(time.Hour*time.Duration(config.Value*24*7*365)).Unix()<time.Now().Unix(){

				}
				break
			default:
				break
			}
			if t.Add(time.Minute*10).Unix()<time.Now().Unix(){

			}
			time.Sleep(time.Minute)
		}

	}()
}
func LogFile(flag TimerFlag) *os.File {
	t:=time.Now().Format(time.RFC3339)
	t=strings.Replace(t,"T","-",-1)
	t=strings.Replace(t,"Z","",-1)
	t=strings.Replace(t,":","-",-1)
	t=strings.Split(t,"+")[0]
	l:=len(t)
	//yyyy-mm-dd-hh-mm-ss
	/* error result
	case 1:
	case 2:
	break
    right result
	case 1:break
	case 2:
	break
	*/
	switch flag {
	case Min:l=15
		break
	case Hour:l=13
		break
	case Day:l=10
		break
	case Week:l=10
		break
	case Month:l=10
		break
	case Year:l=10
	break
		
	}
	log.Println(l)
	var cs []rune=make([]rune,l)
	var vs=[]rune(t)
	for i := 0; i < l; i++ {
		cs[i]=vs[i]
	}
	//copy([]rune(t),cs) //ex
	file, err := os.OpenFile("logs/log-"+string(cs)+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Printf("create log file fail,err=>%s",err.Error())
		return nil
	}
	log.Printf("create log file suc")


	log.SetOutput(file)
	log.SetFlags(log.Llongfile|log.LstdFlags)
	log.Println("test")
	return file
}
