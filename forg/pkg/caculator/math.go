package caculator

// 首字母大写 可以跨包访问
// 如果小写只能在当前包下访问
// 同包下不同文件的全局资源可以随意访问
func Add(a, b int) int {
	return a + b
}
