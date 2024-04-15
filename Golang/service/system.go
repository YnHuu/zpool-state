package service

func (s *service) Reboot() {
	shell(`sleep 1;reboot`)
}
