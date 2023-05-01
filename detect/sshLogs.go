package detect

import (
	"log"
	"os"
	"os/exec"
)

func AccessLog() string {
	if _, err := os.Stat("/var/log/auth.log"); err == nil {
		acLog, err := exec.Command("bash", "-c", "cat /var/log/auth.log | grep -i \"Accepted password\" ").Output()

		if err != nil {
			log.Fatal(err)
		}

		return string(acLog)
	} else {
		acLog, err := exec.Command("bash", "-c", "journalctl -u sshd | grep -i \"Accepted password\" ").Output()
		if err != nil {
			log.Fatal(err)
		}
		return string(acLog)
	}
}

func ErrorLog() string {
	if _, err := os.Stat("/var/log/auth.log"); err == nil {
		erLog, err := exec.Command("bash", "-c", "cat /var/log/auth.log | grep -i \"Failed password\" ").Output()

		if err != nil {
			log.Fatal(err)
		}
		return string(erLog)
	} else {
		erLog, err := exec.Command("bash", "-c", "journalctl -u sshd | grep -i \"Failed password\" ").Output()
		if err != nil {
			log.Fatal(err)
		}
		return string(erLog)
	}
}
