package main

import "fmt"
import "bytes"
import "errors"
import "unsafe"

// #cgo CFLAGS: -I/usr/local/Cellar/libftdi/1.1/include/libftdi1/
// #cgo LDFLAGS: -lftdi1 -L/usr/local/Cellar/libftdi/1.1/lib/
// #include <ftdi.h>
import "C"

// Return Library version, formatted to match D2XX
func GetLibraryVersion() uint32 {
	v := C.ftdi_get_library_version()
	return uint32(v.major&0xFF<<16 +
		v.minor&0xFF<<8 +
		v.micro&0xFF)
}

// Return the number of connected FTDI USB devices
func CreateDeviceInfoList() (n int64, e error) {
	ctx := C.ftdi_new()
	defer C.ftdi_free(ctx)
	if ctx == nil {
		return 0, errors.New("Failed to create FTDI context")
	}

	var dev_list *C.struct_ftdi_device_list
	defer C.ftdi_list_free(&dev_list)

	num := C.ftdi_usb_find_all(ctx, &dev_list, 0, 0)
	if num < 0 {
		return 0, getErr(ctx)
	}

	return int64(num), nil
}

type DeviceDescriptor struct {
	index         uint64
	flags         uint64 // not used in linux
	dtype         uint64 // not used in linux
	id            uint64 // not used
	location      uint64 // not used
	serial_number string
	description   string
	handle        uintptr // the libusb device pointer
}

// Return a description of the USB device
func GetDeviceInfoDetail(index uint32) (descriptor *DeviceDescriptor, e error) {
	ctx := C.ftdi_new()
	defer C.ftdi_free(ctx)
	if ctx == nil {
		return nil, errors.New("Failed to create FTDI context")
	}

	var dev_list *C.struct_ftdi_device_list
	defer C.ftdi_list_free(&dev_list)

	num := C.ftdi_usb_find_all(ctx, &dev_list, 0, 0)
	if num < 0 {
		return nil, getErr(ctx)
	}

	// walk through the linked list
	for i := uint32(0); i < index; i++ {
		dev_list = dev_list.next
	}

	const CHAR_SZ = 256
	var mnf_char, desc_char, ser_char [CHAR_SZ]C.char

	ret := C.ftdi_usb_get_strings(ctx, dev_list.dev,
		(*C.char)(&mnf_char[0]), CHAR_SZ,
		(*C.char)(&desc_char[0]), CHAR_SZ,
		(*C.char)(&ser_char[0]), CHAR_SZ)
	if ret > 0 {
		return nil, getErr(ctx)
	}

	d := &DeviceDescriptor{}
	d.handle = uintptr(unsafe.Pointer(dev_list.dev))
	d.index = uint64(index)
	//d.manufacturer = C.GoString(&mnf_char[0])
	d.description = C.GoString(&desc_char[0])
	d.serial_number = C.GoString(&ser_char[0])

	return d, nil
}

//
func Open(index uint32) (handle uintptr, e error) {
}

func Close(handle uintptr) (e error) {
}

/*
func GetStatus(handle uintptr) (rx_queue, tx_queue, events int32, e error) {
}

func Read(handle uintptr, bytesToRead uint32) (b []byte, e error) {
}

func Write(handle uintptr, b []byte) (e error) {
}

func SetBaudRate(handle uintptr, baud uint32) (e error) {
}
*/

func bytesToString(b []byte) string {
	n := bytes.Index(b, []byte{0})
	return string(b[:n])
}

func getErr(ctx *C.struct_ftdi_context) error {
	return errors.New(C.GoString(C.ftdi_get_error_string(ctx)))
}

func main() {

	fmt.Printf("%X", GetLibraryVersion())
	n, err := CreateDeviceInfoList()
	fmt.Println("Num:", n, err)

	d, err := GetDeviceInfoDetail(0)
	fmt.Println("Info:", d, err)
}
