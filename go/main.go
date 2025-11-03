package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
	"runtime"
	"time"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"strings"
	"strconv"
)

type PingResult struct {
	Size    int           `json:"size"`
	Host    string        `json:"host"`
	Elapsed float64		  `json:"elapsed"`
	Payload string        `json:"payload"`
}

func main() {

	host := flag.String("host", "8.8.8.8", "host to ping")
	text := flag.String("text", "hello", "custom text")
	timeout := flag.Int("timeout", 1000, "timeout per ping in milliseconds")
	flag.Parse()

	switch runtime.GOOS {
	case "darwin":		
		runPingDarwin(*text, *host, *timeout)

	case "windows":
		runPingWindows(*text, *host, *timeout)

	default:
		fmt.Println("Unsupported OS")
	}
}

func runPingWindows(text string, host string, timeout int) {

	conn, err := icmp.ListenPacket("ip4:icmp", "")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	
	dstIP := net.ParseIP(host)
	if dstIP == nil {
		addrs, err := net.LookupIP(host)
		if err != nil || len(addrs) == 0 {
			panic(fmt.Sprintf("cannot resolve host: %v", host))
		}
		dstIP = addrs[0]
	}
	
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   1,
			Seq:  1,
			Data: []byte(text),
		},
	}
	
	data, err := msg.Marshal(nil)
	if err != nil {
		panic(err)
	}
	
	start := time.Now()
	if _, err := conn.WriteTo(data, &net.IPAddr{IP: dstIP}); err != nil {
		panic(err)
	}
	
	buf := make([]byte, 1500)
	conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		panic(err)
	}
	
	elapsed := time.Since(start)
		
	result := PingResult{
		Size:    n-8,
		Host:    addr.String(),
		Elapsed: float64(elapsed.Microseconds()) / 1000,
		Payload: string(buf[8:n]),
	}
	
	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonBytes))		
}

func runPingDarwin(text string, host string, timeout int) {

	// macOS: use /sbin/ping with custom payload
	payloadHex := hex.EncodeToString([]byte(text))
	payloadSize := len(text)
	
	start := time.Now()

	cmd := exec.Command(
		"ping",
		"-c", fmt.Sprintf("%d", 1),
		"-p", payloadHex,
		"-s", fmt.Sprintf("%d", payloadSize),
		"-W", fmt.Sprintf("%d", timeout/1000),
		host,
	)

	//fmt.Printf("Running: %v\n", cmd.Args)

	timer := time.AfterFunc(time.Duration(timeout)*time.Millisecond*time.Duration(2), func() {
		cmd.Process.Kill()
		fmt.Println("Ping timed out")
	})
	defer timer.Stop()

	out, err := cmd.CombinedOutput()
	elapsed := time.Since(start)
	if err != nil {
		fmt.Println("Ping error:", err)
		return
	}
	
	var size int
	lines := string(out)
	bytesStr := extractBytes(lines)	
	size, _ = strconv.Atoi(bytesStr)
		
	result := PingResult{
		Size:    size,
		Host:    extractAddr(lines),
		Elapsed: float64(elapsed.Microseconds()) / 1000,
		Payload: text,
	}
	
	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonBytes))	
}

func extractBytes(output string) string {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "data bytes") {
			parts := strings.Split(line, ":")
			if len(parts) < 2 {
				continue
			}
			sizePart := strings.Fields(strings.TrimSpace(parts[1]))[0]
			if _, err := strconv.Atoi(sizePart); err == nil {
				return sizePart
			}
		}
	}
	return "?"
}

func extractAddr(output string) string {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PING") {
			start := strings.Index(line, "(")
			end := strings.Index(line, ")")
			if start != -1 && end != -1 && start < end {
				return line[start+1 : end]
			}
		}
	}
	return ""
}