/**
 * @Author: DPY
 * @Description:
 * @File:  policy.go
 * @Version: 1.0.0
 * @Date: 2022/4/14 15:03
 */

package ac

import (
	"net/http"
)

type PolicyService struct {
	AC
}

// GetNetPolicy 获取设备已有上网策略信息
func (d PolicyService) GetNetPolicy() ([]NetPolicy, error) {
	var r []NetPolicy
	err := d.request(http.MethodGet, `policy/netpolicy`).Do(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetFluxPolicy 获取设备已有流控策略信息
func (d PolicyService) GetFluxPolicy() ([]FluxPolicy, error) {
	var r []FluxPolicy
	err := d.request(http.MethodGet, `policy/fluxpolicy`).Do(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
