package LoginBFF

import (
	"LaunchService/LaunchUser/client"
)
// 注册服务
func LoginBFF(UserName string, Password string) {
	
	client.LaunchLoginClient(UserName, Password)
}
