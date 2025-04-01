package init

// init函数会在main函数执行前执行
// 初始化函数，进行一些初始化操作
// 每一个源文件都可以包含一个init函数，该函数会在main函数执行前执行
// 全局变量定义->init函数 -> main函数
// 多源文件都有init函数的时候执行顺序
// 按照导入包的顺序执行，先执行其他包的init，在执行main函数包中的init -> main
