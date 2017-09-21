## 安装方法

	go get github.com/aurthur-go/aconfig

## 使用方法

>ini配置文件格式样列
	
	[admin]
	username = root
	password = password
	

>初始化

	conf := aconfig.SetIni("./conf/conf.ini") //aconfig.SetIni(filepath) 其中filepath是你ini 配置文件的所在位置

>获取一组配置信息

	admin := conf.GetSection("admin") //admin是你的section
	fmt.Println(admin["username"]) //root

>获取单个配置信息

	username := conf.GetValue("admin", "username") //admin是你的[section]，username是你要获取值的key名称
	fmt.Println(username) //root

>删除一个配置信息

	conf.DeleteValue("admin", "username")	//username 是你删除的key
	username = conf.GetValue("admin", "username")
	if len(username) == 0 {
		fmt.Println("username is not exists") //this stdout username is not exists
	}

>添加一个配置信息

	conf.SetValue("admin", "username", "aurthur")
	username = conf.GetValue("admin", "username")
	fmt.Println(username) // 添加配置信息如果存在[section]则添加或者修改对应的值，如果不存在则添加section

>获取所有配置信息

	conf.ReadList() //返回[]map[string]map[string]string的格式 即setion=>key->value
