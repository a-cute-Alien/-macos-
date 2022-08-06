package main

import (
	"fmt"
	"os/exec"
)

func main() {

	//mac 管理员密码
	passwd := ""
	//需要替换mac地址的网卡
	net := "en0"
	//要替换的mac地址
	mac := ""
	//锐捷登陆用户名
	username := ""
	//锐捷登陆密码
	userpass := ""

	//卸载网卡
	cmd, err := exec.Command("bash", "-c", `echo `+passwd+` | sudo -S ifconfig `+net+` down`).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	fmt.Println(output)
	//替换mac地址
	cmd, err = exec.Command("bash", "-c", `echo `+passwd+` | sudo -S ifconfig `+net+` up`).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output = string(cmd)
	fmt.Println(output)
	//加载网卡
	cmd, err = exec.Command("bash", "-c", `echo `+passwd+`| sudo -S ifconfig `+net+` ether `+mac).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output = string(cmd)
	fmt.Println(output)
	//通过mentohust连接锐捷服务器
	cmd, err = exec.Command("bash", "-c", `echo `+passwd+`| sudo -S mentohust -u`+username+` -p`+userpass+` -n`+net+` -dhcp-type0`).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output = string(cmd)
	fmt.Println(output)
	defer exec.Command("exit")

}
func readPropety() {

}
