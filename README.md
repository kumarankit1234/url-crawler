# Design
![image info](./design.png)

# How to use
```
urlCrawler := crawler.New(crawler.Options{})
urlCrawler.Start("https://monzo.com")
for !urlCrawler.IsDone() {
    time.Sleep(1 * time.Second)
}
urlCrawler.Stop()

```