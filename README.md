# Micro Demo

---

## 開啟 Micro 微服務

1. 啟動容器
```shell
$ docker-compose up -d
```

2. 查看 Consul (服務發現)
http://127.0.0.1:8500

3. 查看請求監控
http://127.0.0.1:8501/stats

4. 開啟導覽頁
http://127.0.0.1:8502/registry (已註冊服務列表)
http://127.0.0.1:8502/query (微服務Demo頁)


---

## 相關文件

- 簡介網址：[連結](https://hackmd.io/MwFgbAhgZgjFBMBaEAGEATZ8CsVEQGMoVExsQB2SqAUwQmyA?both#2-%E8%A6%81%E7%AB%99%E5%9C%A8%E5%B7%A8%E4%BA%BA%E7%9A%84%E8%82%A9%E8%86%80%E4%B8%8A-go-micro)
- [Micro官方文件](https://micro.mu/docs)
