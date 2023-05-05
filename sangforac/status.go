/**
 * @Author: DPY
 * @Description: AC状态接口
 * @File:  ac_status.go
 * @Version: 1.0.0
 * @Date: 2022/4/14 14:34
 */

package sangforac

import (
	"net/http"
	"time"
)

type StatusService struct {
	AC
}

// Version 获取版本信息
func (d StatusService) Version() (string, error) {
	var ver string
	err := d.request(http.MethodGet, `status/version`).Do(&ver)
	if err != nil {
		return "", err
	}
	return ver, nil
}

// OnlineUserCount 获取在线用户计数
func (d StatusService) OnlineUserCount() (int, error) {
	var count int
	err := d.request(http.MethodGet, `status/online-user`).Do(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// SessionCount 获取当前设备的会话数
func (d StatusService) SessionCount() (int, error) {
	var snum int
	err := d.request(http.MethodGet, `status/session-num`).Do(&snum)
	if err != nil {
		return 0, err
	}
	return snum, nil
}

// InsideLibs 获取设备内置库版本信息,包含病毒库,URL库等模块
func (d StatusService) InsideLibs() ([]InsideLib, error) {
	var libs []InsideLib
	err := d.request(http.MethodGet, `status/insidelib`).Do(&libs)
	if err != nil {
		return nil, err
	}
	return libs, nil
}

// LogCount 获取日志计数统计(拦截日志,记录日志)
func (d StatusService) LogCount() (LogCount, error) {
	var lcount LogCount
	err := d.request(http.MethodGet, `status/log`).Do(&lcount)
	if err != nil {
		return LogCount{}, err
	}
	return lcount, nil
}

// CpuUsage 获取设备的实时CPU使用率(百分比整数)
func (d StatusService) CpuUsage() (int, error) {
	var usage int
	err := d.request(http.MethodGet, `status/cpu-usage`).Do(&usage)
	if err != nil {
		return 0, err
	}
	return usage, nil
}

// MemUsage 获取设备的实时内存使用率(百分比整数)
func (d StatusService) MemUsage() (int, error) {
	var usage int
	err := d.request(http.MethodGet, `status/mem-usage`).Do(&usage)
	if err != nil {
		return 0, err
	}
	return usage, nil
}

// DiskUsage 获取设备的磁盘使用率(百分比整数)
func (d StatusService) DiskUsage() (int, error) {
	var usage int
	err := d.request(http.MethodGet, `status/disk-usage`).Do(&usage)
	if err != nil {
		return 0, err
	}
	return usage, nil
}

// BandwidthUsage 获取带宽使用率
func (d StatusService) BandwidthUsage() (int, error) {
	var r int
	err := d.request(http.MethodGet, `status/bandwidth-usage`).Do(&r)
	if err != nil {
		return 0, err
	}
	return r, nil
}

// SysTime 获取设备的当前系统时间(e.g:2017-12-13 17:52:11)
func (d StatusService) SysTime() (time.Time, error) {
	var timeString string
	err := d.request(http.MethodGet, `status/sys-time`).Do(&timeString)
	if err != nil {
		return time.Time{}, err
	}
	t, err := time.Parse(`2006-01-02 15:04:05`, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// Throughput 获取设备当前上行和下行流量,bitUnit为true时使用bit单位,ifName匹配接口名称,默认统计所有WAN扣
func (d StatusService) Throughput(filter ...ThroughputFilter) (Throughput, error) {
	var tput Throughput
	var f struct {
		Filter ThroughputFilter `json:"filter,omitempty"`
	}
	if len(filter) > 0 {
		f.Filter = filter[0]
	}
	err := d.request(http.MethodPost, `status/throughput`).
		AddQuery("_method", "GET").
		SetBody(f).
		Do(&tput)
	if err != nil {
		return Throughput{}, err
	}
	return tput, nil
}

// UserRank 获取用户流量排行
func (d StatusService) UserRank(filter ...UserRankFilter) ([]UserRank, error) {
	var rst []UserRank
	var f struct {
		Filter UserRankFilter `json:"filter,omitempty"`
	}
	if len(filter) > 0 {
		f.Filter = filter[0]
	}
	err := d.request(http.MethodPost, `status/user-rank`).
		AddQuery("_method", "GET").
		SetBody(f).
		Do(&rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

// AppRank 获取应用流量排行
func (d StatusService) AppRank(filter ...AppRankFilter) ([]AppRank, error) {
	var rst []AppRank
	var f struct {
		Filter AppRankFilter `json:"filter,omitempty"`
	}
	if len(filter) > 0 {
		f.Filter = filter[0]
	}
	err := d.request(http.MethodPost, `status/app-rank`).
		AddQuery("_method", "GET").
		SetBody(f).
		Do(&rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}
