package proxy_test

import (
	"net"
	"testing"

	"github.com/wandoulabs/codis/pkg/models"
	"github.com/wandoulabs/codis/pkg/proxy"
	"github.com/wandoulabs/codis/pkg/utils"
	"github.com/wandoulabs/codis/pkg/utils/assert"
)

var s *proxy.Proxy

func openProxy() (net.Listener, string) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	assert.MustNoError(err)

	s = proxy.New()
	go s.Serve(l)

	return l, l.Addr().String()
}

func TestProxyInfo(x *testing.T) {
	l, addr := openProxy()
	defer l.Close()

	var c = proxy.NewApiClient(addr)

	info, err := c.GetInfo()
	assert.MustNoError(err)
	assert.Must(info.Version == utils.Version)
	assert.Must(info.Compile == utils.Compile)
	assert.Must(info.Token == s.GetToken())
}

func TestProxyPing(x *testing.T) {
	l, addr := openProxy()
	defer l.Close()

	var c = proxy.NewApiClient(addr)

	assert.Must(c.Ping("Bad Token.") != nil)
	assert.Must(c.Ping(s.GetToken()) == nil)
}

func verifySlots(c *proxy.ApiClient, expect map[int]*models.SlotInfo) {
	slots, err := c.GetSlots()
	assert.MustNoError(err)
	assert.Must(len(slots) == models.MaxSlotNum)

	for i, slot := range expect {
		if slot != nil {
			assert.Must(slots[i].Id == i)
			assert.Must(slot.Target == slots[i].Target)
			assert.Must(slot.Locked == slots[i].Locked)
			assert.Must(slot.MigrateFrom == slots[i].MigrateFrom)
		}
	}
}

func TestLockSlot(x *testing.T) {
	l, addr := openProxy()
	defer l.Close()

	token := s.GetToken()

	var c = proxy.NewApiClient(addr)

	expect := make(map[int]*models.SlotInfo)

	for i := 0; i < 16; i++ {
		assert.MustNoError(c.LockSlot(token, i))
		assert.MustNoError(c.LockSlot(token, i))
		expect[i] = &models.SlotInfo{Locked: true}
	}
	verifySlots(c, expect)

	assert.Must(c.LockSlot(token, -1) != nil)
	assert.Must(c.LockSlot(token, models.MaxSlotNum) != nil)
}

func TestFillSlot(x *testing.T) {
	l, addr := openProxy()
	defer l.Close()

	token := s.GetToken()

	var c = proxy.NewApiClient(addr)

	expect := make(map[int]*models.SlotInfo)

	for i := 0; i < 16; i++ {
		assert.MustNoError(c.LockSlot(token, i))
		expect[i] = &models.SlotInfo{Locked: true}
	}
	verifySlots(c, expect)

	for i := 0; i < 16; i++ {
		addr := "x.x.x.x:xxxx"
		assert.MustNoError(c.FillSlot(token, i, addr, ""))
		expect[i] = &models.SlotInfo{Locked: false, Target: addr}
	}
	verifySlots(c, expect)

	for i := 0; i < 16; i++ {
		assert.MustNoError(c.LockSlot(token, i))
		expect[i].Locked = true
	}
	verifySlots(c, expect)

	for i := 0; i < 16; i++ {
		addr := "y.y.y.y:yyyy"
		from := "z.z.z.z:zzzz"
		assert.MustNoError(c.FillSlot(token, i, addr, from))
		expect[i] = &models.SlotInfo{Locked: false, Target: addr, MigrateFrom: from}
	}
	verifySlots(c, expect)

	assert.Must(c.LockSlot(token, -1) != nil)
	assert.Must(c.LockSlot(token, models.MaxSlotNum) != nil)
}

func TestShutdown(x *testing.T) {
	l, addr := openProxy()
	defer l.Close()

	token := s.GetToken()

	var c = proxy.NewApiClient(addr)

	expect := make(map[int]*models.SlotInfo)

	for i := 0; i < 16; i++ {
		assert.MustNoError(c.LockSlot(token, i))
		expect[i] = &models.SlotInfo{Locked: true}
	}
	verifySlots(c, expect)

	err := c.Shutdown(token)
	assert.MustNoError(err)

	for i := 0; i < models.MaxSlotNum; i++ {
		expect[i] = &models.SlotInfo{}
	}
	verifySlots(c, expect)
}
