// a ssid package to get the SSID of the Wi-Fi you are coonected to.
package ssid

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
)

// String creates the ssid.
// Support only drawin（macOS）
func String() string {
	var ssid string
	switch runtime.GOOS {
	case "darwin":
		ssid = forDarwin()
		// TODO: Support for Linux.
		// TODO: Support for Windows.
	}
	return ssid
}

func forDarwin() string {
	airportBin := "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/A/Resources/airport"
	airportOption := "-I"
	cmd := exec.Command(airportBin, airportOption)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}
	defer cmd.Wait()
	var str string

	if b, err := ioutil.ReadAll(stdout); err == nil {
		str = string(b)
	}
	r := regexp.MustCompile(`[^B]SSID: (.+\n)`)
	// FindStringSubmatch returns []string{"SSID: your-ssid", "your-ssid"}
	name := r.FindStringSubmatch(str)
	return name[1]
}
