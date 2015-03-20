
package main

import "unsafe"
import "syscall"
import "fmt"
import "bytes"

func bytesToString(b []byte) string {
    n := bytes.Index(b, []byte{0})
    return string(b[:n])
}

var d2xx = syscall.MustLoadDLL("ftd2xx.dll")

var (
    createDeviceInfoList = d2xx.MustFindProc("FT_CreateDeviceInfoList")
    getDeviceInfoDetail = d2xx.MustFindProc("FT_GetDeviceInfoDetail")
    ft_open = d2xx.MustFindProc("FT_Open")
    ft_close = d2xx.MustFindProc("FT_Close")
    ft_read = d2xx.MustFindProc("FT_Read")
    ft_write = d2xx.MustFindProc("FT_Write")
    ft_getStatus = d2xx.MustFindProc("FT_GetStatus")
    setBaudRate = d2xx.MustFindProc("FT_SetBaudRate")
    )

func CreateDeviceInfoList() (n int64, e error) {
    r, _, err := createDeviceInfoList.Call(uintptr(unsafe.Pointer(&n)))
    if r == 0 {
        return n, nil
    }
    return n, err
}

type DeviceDescriptor struct {
    index uint64
    flags uint64
    dtype uint64
    id    uint64
    location uint64
    serial_number string
    description string
    handle uintptr
}

func GetDeviceInfoDetail(index uint32) (descriptor *DeviceDescriptor, e error) {
    var d DeviceDescriptor
    var sn [16]byte
    var description [64]byte
    d.index = uint64(index)
    r, _, e := getDeviceInfoDetail.Call(uintptr(index), uintptr(unsafe.Pointer(&(d.flags))), uintptr(unsafe.Pointer(&d.dtype)), uintptr(unsafe.Pointer(&d.id)), uintptr(unsafe.Pointer(&d.location)), uintptr(unsafe.Pointer(&sn)), uintptr(unsafe.Pointer(&description)), uintptr(unsafe.Pointer(&d.handle)))
    if r == 0 {
        d.serial_number = bytesToString(sn[:])
        d.description = bytesToString(description[:])
        return &d, nil
    }
    return &d, e
}

func Open(index uint32) (handle uintptr, e error) {
    r, _, e := ft_open.Call(uintptr(index), uintptr(unsafe.Pointer(&handle)))
    if r == 0 {
        return handle, nil
    }
    return handle, e
}

func Close(handle uintptr) (e error) {
    r, _, e := ft_close.Call(handle)
    if r == 0 {
        return  nil
    }
    return e
}

func GetStatus(handle uintptr) (rx_queue, tx_queue, events int32, e error) {
    r, _, e := ft_getStatus.Call(handle,
        uintptr(unsafe.Pointer(&rx_queue)),
        uintptr(unsafe.Pointer(&tx_queue)),
        uintptr(unsafe.Pointer(&events)))
    if r == 0 {
        return rx_queue, tx_queue, events, nil
    }
    return rx_queue, tx_queue, events, e
}

func Read(handle uintptr, bytesToRead uint32) (b []byte, e error) {
    var bytesWritten uint32
    b = make([]byte, bytesToRead)
    ptr := &b[0] //A reference to the first element of the underlying "array"
    r, _, e := ft_read.Call(handle,
            uintptr(unsafe.Pointer(ptr)),
            uintptr(bytesToRead),
            uintptr(unsafe.Pointer(&bytesWritten)))
    if r == 0 {
        return b[:bytesWritten], nil
    }
    return b, e
}

func Write(handle uintptr, b []byte) (e error) {
    var bytesWritten uint32
    bytesToWrite := len(b)
    ptr := &b[0] //A reference to the first element of the underlying "array"
    r, _, e := ft_write.Call(handle,
            uintptr(unsafe.Pointer(ptr)),
            uintptr(bytesToWrite),
            uintptr(unsafe.Pointer(&bytesWritten)))
    if r == 0 {
        return nil
    }
    return e
}

func SetBaudRate(handle uintptr, baud uint32) (e error) {
    r, _, e := setBaudRate.Call(handle, uintptr(baud))
    if r == 0 {
        return nil
    }
    return e
}


func main() {
    var getLibVer = d2xx.MustFindProc("FT_GetLibraryVersion")
    var version uint64
    r, _, _ := getLibVer.Call( uintptr(unsafe.Pointer(&version)))
    fmt.Printf("R: %d, VERSION %X",r, version)
    n, _ := CreateDeviceInfoList()
    fmt.Println("DEVS", n)
    d, _ := GetDeviceInfoDetail(0)
    fmt.Println(d)
    d, _ = GetDeviceInfoDetail(1)
    fmt.Println(d)
    d, _ = GetDeviceInfoDetail(2)
    fmt.Println(d)
    d, _ = GetDeviceInfoDetail(3)
    fmt.Println(d)

    dev, _ := Open(2)
    rx, tx, ev, _ := GetStatus(dev)
    fmt.Println(rx,tx,ev,dev)
    SetBaudRate(dev, 912600)
    msg := make([]byte, 1000,1000)
    Write(dev, msg)
    rx, tx, ev, _ = GetStatus(dev)
    fmt.Println(rx,tx,ev,dev)
    Close(dev)
}
