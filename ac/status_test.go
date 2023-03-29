package ac

import (
	"testing"
)

func TestVersion(t *testing.T) {
	ver, err := AClient.Status().Version()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(ver)
}

func TestOnlineUserCount(t *testing.T) {
	c, err := AClient.Status().OnlineUserCount()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c)
}

func TestSessionCount(t *testing.T) {
	c, err := AClient.Status().SessionCount()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c)
}

func TestInsideLibs(t *testing.T) {
	libs, err := AClient.Status().InsideLibs()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", libs)
}

func TestLogCount(t *testing.T) {
	lc, err := AClient.Status().LogCount()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", lc)
}

func TestCpuUsage(t *testing.T) {
	usg, err := AClient.Status().CpuUsage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(usg)
}

func TestMemUsage(t *testing.T) {
	usg, err := AClient.Status().MemUsage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(usg)
}

func TestDiskUsage(t *testing.T) {
	usg, err := AClient.Status().DiskUsage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(usg)
}

func TestBandwidthUsage(t *testing.T) {
	usg, err := AClient.Status().BandwidthUsage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(usg)
}

func TestSysTime(t *testing.T) {
	ts, err := AClient.Status().SysTime()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ts)
}

func TestThroughput(t *testing.T) {
	tput, err := AClient.Status().Throughput()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", tput)
}

func TestUserRank(t *testing.T) {
	r, err := AClient.Status().UserRank()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(IndentJson(r))
}

func TestAppRank(t *testing.T) {
	r, err := AClient.Status().AppRank(AppRankFilter{Top: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(IndentJson(r))
}
