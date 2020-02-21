// package outputdevice acts as a client for the org_kde_kwin_outputdevice wayland protocol.

// generated by wl-scanner
// https://github.com/dkolbly/wl-scanner
// from: outputdevice.xml
// on 2020-01-20 20:33:23 +0800
package outputdevice

import (
	"sync"

	"github.com/dkolbly/wl"
	"golang.org/x/net/context"
)

type OutputdeviceGeometryEvent struct {
	EventContext   context.Context
	X              int32
	Y              int32
	PhysicalWidth  int32
	PhysicalHeight int32
	Subpixel       int32
	Make           string
	Model          string
	Transform      int32
}

type OutputdeviceGeometryHandler interface {
	HandleOutputdeviceGeometry(OutputdeviceGeometryEvent)
}

func (p *Outputdevice) AddGeometryHandler(h OutputdeviceGeometryHandler) {
	if h != nil {
		p.mu.Lock()
		p.geometryHandlers = append(p.geometryHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveGeometryHandler(h OutputdeviceGeometryHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.geometryHandlers {
		if e == h {
			p.geometryHandlers = append(p.geometryHandlers[:i], p.geometryHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceModeEvent struct {
	EventContext context.Context
	Flags        uint32
	Width        int32
	Height       int32
	Refresh      int32
	ModeId       int32
}

type OutputdeviceModeHandler interface {
	HandleOutputdeviceMode(OutputdeviceModeEvent)
}

func (p *Outputdevice) AddModeHandler(h OutputdeviceModeHandler) {
	if h != nil {
		p.mu.Lock()
		p.modeHandlers = append(p.modeHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveModeHandler(h OutputdeviceModeHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.modeHandlers {
		if e == h {
			p.modeHandlers = append(p.modeHandlers[:i], p.modeHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceDoneEvent struct {
	EventContext context.Context
}

type OutputdeviceDoneHandler interface {
	HandleOutputdeviceDone(OutputdeviceDoneEvent)
}

func (p *Outputdevice) AddDoneHandler(h OutputdeviceDoneHandler) {
	if h != nil {
		p.mu.Lock()
		p.doneHandlers = append(p.doneHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveDoneHandler(h OutputdeviceDoneHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.doneHandlers {
		if e == h {
			p.doneHandlers = append(p.doneHandlers[:i], p.doneHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceScaleEvent struct {
	EventContext context.Context
	Factor       int32
}

type OutputdeviceScaleHandler interface {
	HandleOutputdeviceScale(OutputdeviceScaleEvent)
}

func (p *Outputdevice) AddScaleHandler(h OutputdeviceScaleHandler) {
	if h != nil {
		p.mu.Lock()
		p.scaleHandlers = append(p.scaleHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveScaleHandler(h OutputdeviceScaleHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.scaleHandlers {
		if e == h {
			p.scaleHandlers = append(p.scaleHandlers[:i], p.scaleHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceEdidEvent struct {
	EventContext context.Context
	Raw          string
}

type OutputdeviceEdidHandler interface {
	HandleOutputdeviceEdid(OutputdeviceEdidEvent)
}

func (p *Outputdevice) AddEdidHandler(h OutputdeviceEdidHandler) {
	if h != nil {
		p.mu.Lock()
		p.edidHandlers = append(p.edidHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveEdidHandler(h OutputdeviceEdidHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.edidHandlers {
		if e == h {
			p.edidHandlers = append(p.edidHandlers[:i], p.edidHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceEnabledEvent struct {
	EventContext context.Context
	Enabled      int32
}

type OutputdeviceEnabledHandler interface {
	HandleOutputdeviceEnabled(OutputdeviceEnabledEvent)
}

func (p *Outputdevice) AddEnabledHandler(h OutputdeviceEnabledHandler) {
	if h != nil {
		p.mu.Lock()
		p.enabledHandlers = append(p.enabledHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveEnabledHandler(h OutputdeviceEnabledHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.enabledHandlers {
		if e == h {
			p.enabledHandlers = append(p.enabledHandlers[:i], p.enabledHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceUuidEvent struct {
	EventContext context.Context
	Uuid         string
}

type OutputdeviceUuidHandler interface {
	HandleOutputdeviceUuid(OutputdeviceUuidEvent)
}

func (p *Outputdevice) AddUuidHandler(h OutputdeviceUuidHandler) {
	if h != nil {
		p.mu.Lock()
		p.uuidHandlers = append(p.uuidHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveUuidHandler(h OutputdeviceUuidHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.uuidHandlers {
		if e == h {
			p.uuidHandlers = append(p.uuidHandlers[:i], p.uuidHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceScalefEvent struct {
	EventContext context.Context
	Factor       float32
}

type OutputdeviceScalefHandler interface {
	HandleOutputdeviceScalef(OutputdeviceScalefEvent)
}

func (p *Outputdevice) AddScalefHandler(h OutputdeviceScalefHandler) {
	if h != nil {
		p.mu.Lock()
		p.scalefHandlers = append(p.scalefHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveScalefHandler(h OutputdeviceScalefHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.scalefHandlers {
		if e == h {
			p.scalefHandlers = append(p.scalefHandlers[:i], p.scalefHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceColorcurvesEvent struct {
	EventContext context.Context
	Red          []int32
	Green        []int32
	Blue         []int32
}

type OutputdeviceColorcurvesHandler interface {
	HandleOutputdeviceColorcurves(OutputdeviceColorcurvesEvent)
}

func (p *Outputdevice) AddColorcurvesHandler(h OutputdeviceColorcurvesHandler) {
	if h != nil {
		p.mu.Lock()
		p.colorcurvesHandlers = append(p.colorcurvesHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveColorcurvesHandler(h OutputdeviceColorcurvesHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.colorcurvesHandlers {
		if e == h {
			p.colorcurvesHandlers = append(p.colorcurvesHandlers[:i], p.colorcurvesHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceSerialNumberEvent struct {
	EventContext context.Context
	SerialNumber string
}

type OutputdeviceSerialNumberHandler interface {
	HandleOutputdeviceSerialNumber(OutputdeviceSerialNumberEvent)
}

func (p *Outputdevice) AddSerialNumberHandler(h OutputdeviceSerialNumberHandler) {
	if h != nil {
		p.mu.Lock()
		p.serialNumberHandlers = append(p.serialNumberHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveSerialNumberHandler(h OutputdeviceSerialNumberHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.serialNumberHandlers {
		if e == h {
			p.serialNumberHandlers = append(p.serialNumberHandlers[:i], p.serialNumberHandlers[i+1:]...)
			break
		}
	}
}

type OutputdeviceEisaIdEvent struct {
	EventContext context.Context
	EisaId       string
}

type OutputdeviceEisaIdHandler interface {
	HandleOutputdeviceEisaId(OutputdeviceEisaIdEvent)
}

func (p *Outputdevice) AddEisaIdHandler(h OutputdeviceEisaIdHandler) {
	if h != nil {
		p.mu.Lock()
		p.eisaIdHandlers = append(p.eisaIdHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Outputdevice) RemoveEisaIdHandler(h OutputdeviceEisaIdHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.eisaIdHandlers {
		if e == h {
			p.eisaIdHandlers = append(p.eisaIdHandlers[:i], p.eisaIdHandlers[i+1:]...)
			break
		}
	}
}

func (p *Outputdevice) Dispatch(ctx context.Context, event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.geometryHandlers) > 0 {
			ev := OutputdeviceGeometryEvent{}
			ev.EventContext = ctx
			ev.X = event.Int32()
			ev.Y = event.Int32()
			ev.PhysicalWidth = event.Int32()
			ev.PhysicalHeight = event.Int32()
			ev.Subpixel = event.Int32()
			ev.Make = event.String()
			ev.Model = event.String()
			ev.Transform = event.Int32()
			p.mu.RLock()
			for _, h := range p.geometryHandlers {
				h.HandleOutputdeviceGeometry(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.modeHandlers) > 0 {
			ev := OutputdeviceModeEvent{}
			ev.EventContext = ctx
			ev.Flags = event.Uint32()
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			ev.Refresh = event.Int32()
			ev.ModeId = event.Int32()
			p.mu.RLock()
			for _, h := range p.modeHandlers {
				h.HandleOutputdeviceMode(ev)
			}
			p.mu.RUnlock()
		}
	case 2:
		if len(p.doneHandlers) > 0 {
			ev := OutputdeviceDoneEvent{}
			ev.EventContext = ctx
			p.mu.RLock()
			for _, h := range p.doneHandlers {
				h.HandleOutputdeviceDone(ev)
			}
			p.mu.RUnlock()
		}
	case 3:
		if len(p.scaleHandlers) > 0 {
			ev := OutputdeviceScaleEvent{}
			ev.EventContext = ctx
			ev.Factor = event.Int32()
			p.mu.RLock()
			for _, h := range p.scaleHandlers {
				h.HandleOutputdeviceScale(ev)
			}
			p.mu.RUnlock()
		}
	case 4:
		if len(p.edidHandlers) > 0 {
			ev := OutputdeviceEdidEvent{}
			ev.EventContext = ctx
			ev.Raw = event.String()
			p.mu.RLock()
			for _, h := range p.edidHandlers {
				h.HandleOutputdeviceEdid(ev)
			}
			p.mu.RUnlock()
		}
	case 5:
		if len(p.enabledHandlers) > 0 {
			ev := OutputdeviceEnabledEvent{}
			ev.EventContext = ctx
			ev.Enabled = event.Int32()
			p.mu.RLock()
			for _, h := range p.enabledHandlers {
				h.HandleOutputdeviceEnabled(ev)
			}
			p.mu.RUnlock()
		}
	case 6:
		if len(p.uuidHandlers) > 0 {
			ev := OutputdeviceUuidEvent{}
			ev.EventContext = ctx
			ev.Uuid = event.String()
			p.mu.RLock()
			for _, h := range p.uuidHandlers {
				h.HandleOutputdeviceUuid(ev)
			}
			p.mu.RUnlock()
		}
	case 7:
		if len(p.scalefHandlers) > 0 {
			ev := OutputdeviceScalefEvent{}
			ev.EventContext = ctx
			ev.Factor = event.Float32()
			p.mu.RLock()
			for _, h := range p.scalefHandlers {
				h.HandleOutputdeviceScalef(ev)
			}
			p.mu.RUnlock()
		}
	case 8:
		if len(p.colorcurvesHandlers) > 0 {
			ev := OutputdeviceColorcurvesEvent{}
			ev.EventContext = ctx
			ev.Red = event.Array()
			ev.Green = event.Array()
			ev.Blue = event.Array()
			p.mu.RLock()
			for _, h := range p.colorcurvesHandlers {
				h.HandleOutputdeviceColorcurves(ev)
			}
			p.mu.RUnlock()
		}
	case 9:
		if len(p.serialNumberHandlers) > 0 {
			ev := OutputdeviceSerialNumberEvent{}
			ev.EventContext = ctx
			ev.SerialNumber = event.String()
			p.mu.RLock()
			for _, h := range p.serialNumberHandlers {
				h.HandleOutputdeviceSerialNumber(ev)
			}
			p.mu.RUnlock()
		}
	case 10:
		if len(p.eisaIdHandlers) > 0 {
			ev := OutputdeviceEisaIdEvent{}
			ev.EventContext = ctx
			ev.EisaId = event.String()
			p.mu.RLock()
			for _, h := range p.eisaIdHandlers {
				h.HandleOutputdeviceEisaId(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type Outputdevice struct {
	wl.BaseProxy
	mu                   sync.RWMutex
	geometryHandlers     []OutputdeviceGeometryHandler
	modeHandlers         []OutputdeviceModeHandler
	doneHandlers         []OutputdeviceDoneHandler
	scaleHandlers        []OutputdeviceScaleHandler
	edidHandlers         []OutputdeviceEdidHandler
	enabledHandlers      []OutputdeviceEnabledHandler
	uuidHandlers         []OutputdeviceUuidHandler
	scalefHandlers       []OutputdeviceScalefHandler
	colorcurvesHandlers  []OutputdeviceColorcurvesHandler
	serialNumberHandlers []OutputdeviceSerialNumberHandler
	eisaIdHandlers       []OutputdeviceEisaIdHandler
}

func NewOutputdevice(ctx *wl.Context) *Outputdevice {
	ret := new(Outputdevice)
	ctx.Register(ret)
	return ret
}

const (
	OutputdeviceSubpixelUnknown       = 0
	OutputdeviceSubpixelNone          = 1
	OutputdeviceSubpixelHorizontalRgb = 2
	OutputdeviceSubpixelHorizontalBgr = 3
	OutputdeviceSubpixelVerticalRgb   = 4
	OutputdeviceSubpixelVerticalBgr   = 5
)

const (
	OutputdeviceTransformNormal     = 0
	OutputdeviceTransform90         = 1
	OutputdeviceTransform180        = 2
	OutputdeviceTransform270        = 3
	OutputdeviceTransformFlipped    = 4
	OutputdeviceTransformFlipped90  = 5
	OutputdeviceTransformFlipped180 = 6
	OutputdeviceTransformFlipped270 = 7
)

const (
	OutputdeviceModeCurrent   = 0x1
	OutputdeviceModePreferred = 0x2
)

const (
	OutputdeviceEnablementDisabled = 0
	OutputdeviceEnablementEnabled  = 1
)