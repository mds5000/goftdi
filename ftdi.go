
// The real API

type DeviceInfo struct {
	index uint64
	flags uint64
	dtype uint64
	id    uint64
	location uint64
	serial_number string
	description string
	handle unsafe.Pointer
}

func GetDeviceList() []DeviceInfo {}

func Open(*DeviceInfo) *Device {}

// Implements ReadWriteCloser
func (d *Device) Close() error {}
func (d *Device) Read([]byte) int, error {}
func (d *Device) Write([]byte) int, error {}
func (d *Device) SetBaudrate(int) {}
func (d *Device) SetBitmode() {}
func (d *Device) SetFlowControl() {}
func (d *Device) SetLatency() {}
func (d *Device) SetLineProperty() {}
func (d *Device) SetTimeout() {}
func (d *Device) SetTransferSize() {}
func (d *Device) Reset() {}



type MpsseDevice struct {}
func
