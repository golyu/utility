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
	"sort"
	"strconv"
	"strings"
)

// AppendStr appends string to slice with no duplicates.
func AppendStr(strs []string, str string) []string {
	for _, s := range strs {
		if s == str {
			return strs
		}
	}
	return append(strs, str)
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements and order are both the same.
func CompareSliceStr(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements are the same, and ignores the order.
func CompareSliceStrU(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				s2 = append(s2[:j], s2[j+1:]...)
				break
			}
		}
	}
	if len(s2) > 0 {
		return false
	}
	return true
}

// IsSliceContainsStr returns true if the string exists in given slice, ignore case.
func IsSliceContainsStr(sl []string, str string) bool {
	str = strings.ToLower(str)
	for _, s := range sl {
		if strings.ToLower(s) == str {
			return true
		}
	}
	return false
}

// IsSliceContainsInt64 returns true if the int64 exists in given slice.
func IsSliceContainsInt64(sl []int64, i int64) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

// IntSliceDeduplication int的切片去重复
func IntSliceDeduplication(ints []int) []int {
	temp := make(map[int]int)
	for i, _ := range ints {
		temp[ints[i]] = 0
	}
	ints = make([]int, 0)
	for k := range temp {
		ints = append(ints, k)
	}
	return ints
}

//[]string 转 []int
func Strings2Ints(src []string) (dst []int, err error) {
	if len(src) == 0 {
		return nil, errors.New("string convert to int error:not found src ")
	}
	for _, v := range src {
		temp, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("string convert to int error:not found src")
		}
		dst = append(dst, temp)
	}
	return
}

// Ints2Strings []int转[]string
func Ints2Strings(src []int) []string {
	dst := make([]string, len(src))
	for i, v := range src {
		dst[i] = strconv.Itoa(v)
	}
	return dst
}

//[]string 转 []int 排序从小到大
func Strings2IntsAsc(src []string) (dst []int, err error) {
	dst, err = Strings2Ints(src)
	if err != nil {
		return
	}
	sort.Ints(dst)
	return
}
