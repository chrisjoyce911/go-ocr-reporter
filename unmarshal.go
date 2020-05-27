package ocr

import (
	"strconv"
	"strings"
	"time"
)

type SpecialBool struct {
	Bool bool
}

type SpecialDate struct {
	time.Time
}

type SpecialUnix struct {
	time.Time
}

type SpecialInt struct {
	int
}

func (sd *SpecialDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)

	var newTime time.Time
	var err error
	if strInput != "null" {
		newTime, err = time.Parse("2006-01-02 15:04:05", strInput)
		if err != nil {
			return err
		}
	}
	sd.Time = newTime
	return nil
}

func (sb *SpecialBool) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	switch strInput {
	case "0":
		sb.Bool = false
	case "1":
		sb.Bool = true
	case "No":
		sb.Bool = false
	case "Yes":
		sb.Bool = true
	}
	return nil
}

func (st *SpecialUnix) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)

	i, err := strconv.ParseInt(strInput, 10, 64)
	if err != nil {
		return err
	}
	st.Time = time.Unix(i, 0)
	return nil
}

func (si *SpecialInt) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)

	var i int
	var err error

	i, err = strconv.Atoi(strInput)
	if err != nil {
		return err
	}

	si.int = i
	return nil
}
