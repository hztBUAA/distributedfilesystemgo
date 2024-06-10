package main

import (
    "io"
    "os"
    "strings"
)

func main() {
    // 创建一个字符串读取器
    r := strings.NewReader("Hello, World!")

    // 打开一个文件用于写入数据
    f, err := os.Create("test2/output.txt")
	// f,err := os.MkdirAll("test2/output.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // 使用 io.Copy 将数据从读取器复制到文件
    _, err = io.Copy(f, r)
    if err != nil {
        panic(err)
    }
}