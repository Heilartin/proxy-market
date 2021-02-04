package proxy_market

import (
	"database/sql"
	"database/sql/driver"
	"github.com/araddon/dateparse"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const (
	MyTimeFormat          = "2006-01-02T15:04:05"
	ProxyMarketTimeFormat = "2021-02-27 21:00:00"
)

type ProxyMarketTime sql.NullTime

func (t ProxyMarketTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(sql.NullTime(t).Time.Format(ProxyMarketTimeFormat))), nil
	//return []byte(sql.NullTime(t).Time.Format(time.RFC3339)), nil
}
func (t *ProxyMarketTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	s = strings.Replace(s, "\"", "", -1)
	if s == "null" || s == "" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = dateparse.ParseLocal(s)
	if err != nil {
		logrus.Error(err)
	}
	return
	//
	//s = strings.Split(string(b), " ")[0]
	//if strings.Contains(s, "/") {
	//	s = strings.Split(s, "/")[1]
	//}
	//if s == "" || s == "null" {
	//	t.Time = time.Time{}
	//	return
	//}

}

func (t ProxyMarketTime) String() string {
	return sql.NullTime(t).Time.Format(ProxyMarketTimeFormat)
	//return sql.NullTime(t).Time.Format(time.RFC3339)
}

// Scan implements the Scanner interface.
func (nt *ProxyMarketTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt ProxyMarketTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
