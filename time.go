// Copyright 2013 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package utility

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	LAYOUT_FORMAT19 = "2006-01-02 15:04:05"
	LAYOUT_FORMAT14 = "20060102150405"
	LAYOUT_FORMAT10 = "2006-01-02"
	LAYOUT_FORMAT8  = "20060102"
	LAYOUT_FORMAT17 = "20060102 15:04:05"
)

var (
	loc, _ = time.LoadLocation("Asia/Chongqing")
)

// Format unix time int64 to string
func Date(ti int64, format string) string {
	t := time.Unix(int64(ti), 0)
	return DateT(t, format)
}

// Format unix time string to string
func DateS(ts string, format string) string {
	i, _ := strconv.ParseInt(ts, 10, 64)
	return Date(i, format)
}

// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func DateT(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

// DateFormat pattern rules.
var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", //A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// Parse Date use PHP time format.
func DateParse(dateString, format string) (time.Time, error) {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return time.ParseInLocation(format, dateString, time.Local)
}

// GetNowTime get current time
func GetNowTime() time.Time {
	return time.Now()
}

// GetNowTimestamp get current timestamp
func GetNowTimestamp() int {
	return int(time.Now().Unix())
}

// Format19 format date YYYY-MM-DD HH:mm:ss
func Format19(time time.Time) string {
	return time.Format(LAYOUT_FORMAT19)
}

// Format17 format date YYYYMMDD HH:mm:ss
func Format17(time time.Time) string {
	return time.Format(LAYOUT_FORMAT17)
}

// Format14 format date YYYYMMDDHHmmss
func Format14(time time.Time) string {
	return time.Format(LAYOUT_FORMAT14)
}

// Format10 format date YYYY-MM-DD
func Format10(time time.Time) string {
	return time.Format(LAYOUT_FORMAT10)
}

// Format8 format date YYYYMMDD
func Format8(time time.Time) string {
	return time.Format(LAYOUT_FORMAT8)
}

// TimeIntervalDay 计算两个时间相隔多少天
func TimeIntervalDay(t1, t2 time.Time) int {
	t1 = t1.UTC().Truncate(24 * time.Hour)
	t2 = t2.UTC().Truncate(24 * time.Hour)
	result := int(t1.Sub(t2).Hours() / 24)
	if result < 0 {
		result *= -1
	}
	return result
}

// TimeIntervalSecond 两个时间相隔秒数
func TimeIntervalSecond(time1 time.Time, time2 time.Time) int {
	d := time2.Sub(time1)
	return int(d / 1e9)
}

// Format19ToTimestamp 格式化的时间(YYYY-MM-DD HH:mm:ss)转时间戳
func Format19ToTimestamp(formatTime string) (int, error) {
	if len(formatTime) != 19 {
		return 0, errors.New("formatTime length error")
	} else {
		sTemp, err := time.ParseInLocation(LAYOUT_FORMAT19, formatTime, loc)
		if err != nil {
			return 0, err
		}
		return int(sTemp.Unix()), nil
	}
}

// Format17ToTimestamp 格式化的时间(YYYYMMDD HH:mm:ss)转时间戳
func Format17ToTimestamp(formatTime string) (int, error) {
	if len(formatTime) != 17 {
		return 0, errors.New("formatTime length error")
	} else {
		sTemp, err := time.ParseInLocation(LAYOUT_FORMAT17, formatTime, loc)
		if err != nil {
			return 0, err
		}
		return int(sTemp.Unix()), nil
	}
}

// Format14ToTimestamp 格式化的时间(YYYYMMDDHHmmss)转时间戳
func Format14ToTimestamp(formatTime string) (int, error) {
	if len(formatTime) != 14 {
		return 0, errors.New("formatTime length error")
	} else {
		sTemp, err := time.ParseInLocation(LAYOUT_FORMAT14, formatTime, loc)
		if err != nil {
			return 0, err
		}
		return int(sTemp.Unix()), nil
	}
}

// Format10ToTimestamp 格式化的时间(YYYY-MM-DD)转时间戳
func Format10ToTimestamp(formatTime string) (int, error) {
	if len(formatTime) != 10 {
		return 0, errors.New("formatTime length error")
	} else {
		sTemp, err := time.ParseInLocation(LAYOUT_FORMAT10, formatTime, loc)
		if err != nil {
			return 0, err
		}
		return int(sTemp.Unix()), nil
	}
}

// Format8ToTimestamp 格式化的时间(YYYYMMDD)转时间戳
func Format8ToTimestamp(formatTime string) (int, error) {
	if len(formatTime) != 8 {
		return 0, errors.New("formatTime length error")
	} else {
		sTemp, err := time.ParseInLocation(LAYOUT_FORMAT8, formatTime, loc)
		if err != nil {
			return 0, err
		}
		return int(sTemp.Unix()), nil
	}
}

// GetTimeIntervalDay 与当前时间的时间间隔 单位天 如 GetTimeIntervalDay(1) 为延后一天
func GetTimeIntervalDay(day int) int {
	d, _ := time.ParseDuration("-24h")
	return int(time.Now().Add(30 * d).Unix())
}

// GetNightTimestamp 获取前一天的时间戳：day = -1 , 获取今天晚上的时间戳：0 ，获取明天的凌晨时间戳：1
func GetNightTimestamp(day int) int {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, day)
	d := yesTime.Format(LAYOUT_FORMAT8)
	res, _ := Format8ToTimestamp(d)
	return res
}

// 计算两个时间相隔多少天
func TimeSub(t1, t2 time.Time) int {
	t1 = t1.UTC().Truncate(24 * time.Hour)
	t2 = t2.UTC().Truncate(24 * time.Hour)
	result := int(t1.Sub(t2).Hours() / 24)
	if result < 0 {
		result *= -1
	}
	return result
}
