package version

// 只有首字母大写的全局变量才能在其他的包中被使用
const Version string = "1.1"

// 小写字母不能被其他包所调用
const createDate = "20240127"
