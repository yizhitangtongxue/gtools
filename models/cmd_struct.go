package models

type CmdStruct struct {
	// COMMAND     PID         USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
	Command    string
	Pid        string
	User       string
	IpProtocol string
	Connection string
}
