/*
	ORM模型定义
*/
package model

// 用户模型
type User struct {
	Name     string
	Password string
	Option   map[string]interface{} // 扩展字段，记录用户的一些额外属性
}

func CheckUser() (bool, error) {
	...
	if err != nil {
		...
	}
	...
	if err != nil{
		...
	}
	...
}

func DoSomething()  {
	...
	lock := resource.Lock()
	...
	...
	lock.UnLock()
}
