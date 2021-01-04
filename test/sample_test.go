package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Test struct {
	Name *string
	Age int
}

type MyTime time.Time

const localDateTimeFormat string = "2006-01-02 15:04:05"


func (l MyTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localDateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(l).AppendFormat(b, localDateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (l *MyTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+localDateTimeFormat+`"`, string(b), time.Local)
	*l = MyTime(now)
	return err
}

func (l MyTime) String() string {
	return time.Time(l).Format(localDateTimeFormat)
}

func TestFunc(t *testing.T) {
	//test := Test{}
	//*test.Name = "abc"
	//fmt.Println(*test.Name)
	dateTime, err := time.Parse("2006-01-02 15:04:05", "1971-02-01 00:04:00")
	tt := MyTime(dateTime)
	//cc := (time.Time)(dateTime)
	//fmt.Println("tt", dateTime, tt, cc, )
	kk, _ := json.Marshal(tt)
	fmt.Println("kk...", kk)
	fmt.Println(string(kk))
	fmt.Println(dateTime, err)
	//fmt.Println("tt", tt, dateTime)
	st, err := time.Parse("2006-01-02 15:04:05", "2020-01-02 15:04:05")
	fmt.Println(st, err)

}
