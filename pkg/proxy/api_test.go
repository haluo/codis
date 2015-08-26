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
	go s.ServeHTTP(l)

	return l, l.Addr().String()
}

func TestInfo(x *testing.T) {
	l, addr := openProxy()
	defer l.Close()

	var c = proxy.NewApiClient(addr)

	info, err := c.GetInfo()
	assert.MustNoError(err)
	assert.Must(info.Version == utils.Version)
	assert.Must(info.Compile == utils.Compile)
	assert.Must(info.Token == s.GetToken())
}

func TestStats(x *testing.T) {
	l, addr := openProxy()
	defer l.Close()

	var c = proxy.NewApiClient(addr)

	_, err1 := c.GetStats("Bad Token.")
	assert.Must(err1 != nil)

	_, err2 := c.GetStats(s.GetToken())
	assert.MustNoError(err2)
}

func verifySlots(c *proxy.ApiClient, expect map[int]*models.SlotInfo) {
	info, err := c.GetInfo()
	assert.MustNoError(err)

	slots := info.Slots
	assert.Must(len(slots) == models.MaxSlotNum)

	for i, slot := range expect {
		if slot != nil {
			assert.Must(slots[i].Id == i)
			assert.Must(slot.Locked == slots[i].Locked)
			assert.Must(slot.BackendAddr == slots[i].BackendAddr)
			assert.Must(slot.MigrateFrom == slots[i].MigrateFrom)
		}
	}
}

func TestFillSlot(x *testing.T) {
	l, addr := openProxy()
	defer l.Close()

	token := s.GetToken()

	var c = proxy.NewApiClient(addr)

	expect := make(map[int]*models.SlotInfo)

	for i := 0; i < 16; i++ {
		slot := &models.SlotInfo{
			Id:          i,
			Locked:      i%2 == 0,
			BackendAddr: "x.x.x.x:xxxx",
		}
		assert.MustNoError(c.FillSlot(token, slot))
		expect[i] = slot
	}
	verifySlots(c, expect)

	slots := []*models.SlotInfo{}
	for i := 0; i < 16; i++ {
		slot := &models.SlotInfo{
			Id:          i,
			Locked:      i%2 != 0,
			BackendAddr: "y.y.y.y:yyyy",
			MigrateFrom: "x.x.x.x:xxxx",
		}
		slots = append(slots, slot)
		expect[i] = slot
	}
	assert.MustNoError(c.FillSlot(token, slots...))
	verifySlots(c, expect)
}

func TestStartAndShutdown(x *testing.T) {
	l, addr := openProxy()
	defer l.Close()

	token := s.GetToken()

	var c = proxy.NewApiClient(addr)

	expect := make(map[int]*models.SlotInfo)

	for i := 0; i < 16; i++ {
		slot := &models.SlotInfo{
			Id:          i,
			BackendAddr: "x.x.x.x:xxxx",
		}
		assert.MustNoError(c.FillSlot(token, slot))
		expect[i] = slot
	}
	verifySlots(c, expect)

	err1 := c.Start(token)
	assert.MustNoError(err1)

	err2 := c.Shutdown(token)
	assert.MustNoError(err2)

	err3 := c.Start(token)
	assert.Must(err3 != nil)
}
