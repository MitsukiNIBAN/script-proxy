package main

import "os/exec"

func startV2ray() error {
	cmd := exec.Command("v2ray", "--config=")
	return cmd.Run()
}

func stopV2ray() error {
	cmd := exec.Command("sudo", "kill", "-9", "$(pidof v2ray)")
	return cmd.Run()
}

func isV2rayRunning() bool {
	cmd := exec.Command("sudo", "pidof", "v2ray")
	cmd.Run()
	return false
}
