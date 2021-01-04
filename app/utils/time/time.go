package time

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
	MysqlLayout = "2006-01-02 15:04:05 +0800 CST"
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, time.Time(t).Format(TimeLayout))
	return []byte(s), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
	// 因为实际接收到值是""2018-11-25 20:04:51""格式的
	now, err := time.ParseInLocation(`"`+TimeLayout+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return err
}

func (t JsonTime) Value() (driver.Value, error) {
	tm := time.Time(t)
	return tm.Format(TimeLayout), nil
}

func (t *JsonTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch st := value.(type) {
	case time.Time:
		*t = JsonTime(st)
	case string:
		tm, err := time.Parse(TimeLayout, st)
		if err != nil {
			return err
		}
		*t = JsonTime(tm)
	}
	return nil
}

func (t JsonTime) String() string {
	return time.Time(t).Format(TimeLayout)
}
