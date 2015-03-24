
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

// Search the system for all connected FTDI devices.
// Returns a slice of `DeviceInfo` objects for each.
func GetDeviceList() []DeviceInfo {}

// Open the device described by DeviceInfo
func Open(di *DeviceInfo) *Device, error {}

// Close the device
// Return nil on success, error otherwise.
func (d *Device) Close() error {}

// Read from the device. Implements io.Reader.
func (d *Device) Read([]byte) int, error {}
// Write to the device. Implements io.Writer.
func (d *Device) Write([]byte) int, error {}

// Set the baudrate/bitrate in bits-per-second.
// Return nil on success, error otherwise.
func (d *Device) SetBaudrate(baud uint) error {}

// Set the device's bit mode.
func (d *Device) SetBitMode(mode BitMode) error {}
func (d *Device) SetFlowControl() error{}
func (d *Device) SetLatency() {}
func (d *Device) SetLineProperty() {}
func (d *Device) SetTimeout() {}
func (d *Device) SetTransferSize() {}
func (d *Device) SetChars() {}

// Reset the device. Returns nil on success,
// error otherwise.
func (d *Device) Reset() error {}
func (d *Device) Purge() error {}


// Others...
func (d *Device) GetStatus() (rx_queue, tx_queue, events int32, e error) {}



type MPSSEDevice uintptr

func (m MPSSEDevice) Initialize(device) (*)
func (m MPSSEDevice) SetGPIO(device) (*)
func (m MPSSEDevice) ReadGPIO(device) (*)
func (m MPSSEDevice) WriteGPIO(device) (*)

func (m MPSSEDevice) SetMode(device) (*)
func (m MPSSEDevice) SetClk(device) (*)
func (m MPSSEDevice) Write(device) (*)
func (m MPSSEDevice) Read(device) (*)
func (m MPSSEDevice) Close(device) (*)

