package main

import (
    "fmt"
    "edge5134.com/notedge/internal/cache"
)

func main() {
    client := cache.New()
    err := client.SetURL("test", "https://example.com")
    if err != nil {
        fmt.Println("失败原因:", err)
    } else {
        fmt.Println("✅ 成功写入Redis!")
    }
}