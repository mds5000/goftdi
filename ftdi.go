package ftdi

type BitMode byte

const (
	RESET         BitMode = 0x00
	ASYNC_BITBANG         = 0x01
	MPSSE                 = 0x02
	SYNC_BITBANG          = 0x04
	HOST_EMU              = 0x08
	FAST_OPTO             = 0x10
	CBUS_BITBANG          = 0x20
	SYNCHRONOUS           = 0x40
)

type FlowControl uint16

const (
	DISABLED FlowControl = 0x0000
	RTS_CTS              = 0x0100
	DTR_DSR              = 0x0200
	XON_XOFF             = 0x0400
)

type LineProperties struct {
	Bits     bitsPerWord
	StopBits stopBits
	Parity   parity
}
type bitsPerWord byte
type stopBits byte
type parity byte

const (
	BITS_8 bitsPerWord = 8
	BIST_7 bitsPerWord = 7
	STOP_1 stopBits    = 0
	STOP_2 stopBits    = 2
	NONE   parity      = 0
	ODD    parity      = 1
	EVEN   parity      = 2
	MARK   parity      = 3
	SPACE  parity      = 4
)

/* API
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
*/
