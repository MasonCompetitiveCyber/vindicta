package ui

import "fyne.io/fyne/v2/widget"
import "os"
import "fmt"
import "os/exec"

// Used to get successful SSH Logins
func AccessLog(tab *widget.Label) {
	if _, err := os.Stat("/var/log/auth.log"); err == nil {
		acLog, err := exec.Command("bash", "-c", "cat /var/log/auth.log | grep -i \"Accepted password\" ").Output()

		if err != nil {
            sshErr := fmt.Sprintf("Cannot find ssh logs in /var/log/auth.log => %s", err)
            tab.SetText(sshErr)
		}
		tab.SetText(string(acLog))
	} else {
		acLog, err := exec.Command("bash", "-c", "journalctl -u sshd | grep -i \"Accepted password\" ").Output()
		if err != nil {
            sshErr := fmt.Sprintf("Cannot find ssh logs with journalctl  => %s", err)
            tab.SetText(sshErr)
		}
        tab.SetText(string(acLog))
	}
}

// Used to get the failed SSH Logs
func ErrorLog(tab *widget.Label) {
	if _, err := os.Stat("/var/log/auth.log"); err == nil {
		erLog, err := exec.Command("bash", "-c", "cat /var/log/auth.log | grep -i \"Failed password\" ").Output()

		if err != nil {	
            sshErr := fmt.Sprintf("Cannot find ssh logs in /var/log/auth.log => %s", err)
            tab.SetText(sshErr)
		}
		tab.SetText(string(erLog))
	} else {
		erLog, err := exec.Command("bash", "-c", "journalctl -u sshd | grep -i \"Failed password\" ").Output()
		if err != nil {
            sshErr := fmt.Sprintf("Cannot find ssh logs with journalctl  => %s", err)
            tab.SetText(sshErr)
		}
        tab.SetText(string(erLog))
	}
}

