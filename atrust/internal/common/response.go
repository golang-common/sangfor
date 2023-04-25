package common

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Response struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data RespData `json:"data"`
}

func (r Response) ErrorCheck() error {
	if r.Code == 0 {
		return nil
	}
	if r.Msg != "" {
		return errors.New(fmt.Sprintf("code=%d,msg=%s", r.Code, r.Msg))
	}
	if len(r.Data) > 0 {
		return errors.New(string(r.Data))
	}
	return errors.New("response check failed")
}

type RespData json.RawMessage

func (r *RespData) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	return *r, nil
}

func (r *RespData) Into(key string) error {
	var temp map[string]json.RawMessage
	err := json.Unmarshal(*r, &temp)
	if err != nil {
		return err
	}
	v, ok := temp[key]
	if !ok {
		return errors.New("not found")
	}
	*r = RespData(v)
	return nil
}

func (r *RespData) UnmarshalJSON(data []byte) error {
	if r == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*r = append((*r)[0:0], data...)
	return nil
}

func (r *RespData) Parse(rst any) error {
	return json.Unmarshal(*r, rst)
}

func (r *RespData) ParseData(title string, rst any) error {
	var temp map[string]json.RawMessage
	var err error
	err = json.Unmarshal(*r, &temp)
	if err != nil {
		return err
	}
	v, ok := temp[title]
	if !ok {
		return errors.New("not found")
	}
	err = json.Unmarshal(v, rst)
	if err != nil {
		return err
	}
	return nil
}

func (r *RespData) ParseDataWithPager(title string, rst any) (Pager, error) {
	var (
		temp  = make(map[string]json.RawMessage)
		pager Pager
		err   error
	)
	err = json.Unmarshal(*r, &temp)
	if err != nil {
		return Pager{}, err
	}
	err = json.Unmarshal(*r, &pager)
	if err != nil {
		return Pager{}, err
	}
	v, ok := temp[title]
	if !ok {
		return Pager{}, errors.New("not found")
	}
	err = json.Unmarshal(v, &rst)
	if err != nil {
		return Pager{}, err
	}
	return pager, nil
}

type Pager struct {
	PageCount int
	PageSize  int
	PageIndex int
	Count     int
}
