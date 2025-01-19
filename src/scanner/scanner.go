package scanner

import (
    "fmt"
    "net"
    "time"
)

func ScanPort(ip string, port int) {
    address := fmt.Sprintf("%s:%d", ip, port)
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)
    if err == nil {
        fmt.Printf("Port %d is open on %s\n", port, ip)
        conn.Close()
    }
}