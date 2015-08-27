package proxy_test

import (
	"net"
	"testing"

	"github.com/wandoulabs/codis/pkg/models"
	"github.com/wandoulabs/codis/pkg/proxy"
	"github.com/wandoulabs/codis/pkg/utils"
	"github.com/wandoulabs/codis/pkg/utils/assert"
)

func openProxy() (*proxy.Proxy, string) {
	l, err := net.Listen("tcp", "0.0.0.0:0")
	assert.MustNoError(err)

	config := proxy.NewDefaultConfig()
	config.ProxyAddr = "0.0.0.0:0"
	config.AdminAddr = l.Addr().String()

	l.Close()

	s, err := proxy.NewWithConfig(config)
	assert.MustNoError(err)

	return s, s.GetConfig().AdminAddr
}

func TestInfo(x *testing.T) {
	s, addr := openProxy()
	defer s.Close()

	var c = proxy.NewApiClient(addr)

	info, err := c.GetInfo()
	assert.MustNoError(err)
	assert.Must(info.Version == utils.Version)
	assert.Must(info.Compile == utils.Compile)
	assert.Must(info.Token == s.GetToken())
}

func TestStats(x *testing.T) {
	s, addr := openProxy()
	defer s.Close()

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
	s, addr := openProxy()
	defer s.Close()

	var c = proxy.NewApiClient(addr)

	expect := make(map[int]*models.SlotInfo)

	for i := 0; i < 16; i++ {
		slot := &models.SlotInfo{
			Id:          i,
			Locked:      i%2 == 0,
			BackendAddr: "x.x.x.x:xxxx",
		}
		assert.MustNoError(c.FillSlot(s.GetToken(), slot))
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
	assert.MustNoError(c.FillSlot(s.GetToken(), slots...))
	verifySlots(c, expect)
}

func TestOnlineAndShutdown(x *testing.T) {
	s, addr := openProxy()
	defer s.Close()

	var c = proxy.NewApiClient(addr)

	expect := make(map[int]*models.SlotInfo)

	for i := 0; i < 16; i++ {
		slot := &models.SlotInfo{
			Id:          i,
			BackendAddr: "x.x.x.x:xxxx",
		}
		assert.MustNoError(c.FillSlot(s.GetToken(), slot))
		expect[i] = slot
	}
	verifySlots(c, expect)

	err1 := c.Online(s.GetToken())
	assert.MustNoError(err1)

	err2 := c.Shutdown(s.GetToken())
	assert.MustNoError(err2)

	err3 := c.Online(s.GetToken())
	assert.Must(err3 != nil)
}
