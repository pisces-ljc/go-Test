package processes

import "fmt"

//UserMgr 实例在服务器端有且只有一个
// 在很多地方需要使用，因此设置为全局变量

//维护在线用户列表
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// AddOnlineUser 完成对 onlineProcess添加
func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserID] = up
}

// DeleteOnlineProcess 删除
func (um *UserMgr) DeleteOnlineProcess(userID int) {
	delete(um.onlineUsers, userID)
}

// GetAllOnlineUser 查询
func (um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return um.onlineUsers
}

// GetOnlineUserById 根据id返回对应的值
func (um *UserMgr) GetOnlineUserById(userID int) (up *UserProcess, err error) {
	up, ok := um.onlineUsers[userID]
	if !ok {
		err = fmt.Errorf("用户%d 不在线", userID)
		return
	}

	return
}
