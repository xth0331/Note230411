package homework

import "fmt"

// 注册流程：
// 1. 要求用户从命令行输入 用户名 ，密码，确认密码。
// 2. 如果，输入不为空，且两次密码相同，则打印 注册成功，并退出程序，否则根据情况：提示：输入不得为空， 两次密码不一致。

func Regiester() {
	for {
		var username, pwd, cpwd string
		fmt.Println("欢迎注册！！！")
		fmt.Print("用户名:")
		fmt.Scanf("%s\n", &username)
		fmt.Print("密码:")
		fmt.Scanf("%s\n", &pwd)
		fmt.Print("确认密码:")
		fmt.Scanf("%s\n", &cpwd)
		if username == "" || pwd == "" || cpwd == "" {
			fmt.Println("输入不得为空")
			continue
		}
		if pwd != cpwd {
			fmt.Println("两次密码不一致")
			continue
		}
		fmt.Println("注册成功")
		break
	}

}
