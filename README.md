## Nano DB SDK for Go

### 使用方法
```go
import (
    ndb "git.lolli.tech/lollipopkit/nano-db-sdk-go"
)
var (
    // 创建一个数据库对象
    db = ndb.NewDB("DB Url", "DB Cookie")
)

type Chapter struct {
    Id int64
    Cid int64
    Content string
    Title string
}

func main() {
    // 写入数据
    chapter1 := Chapter{
        title: "1",
        Id: 1,
        Cid: 1,
        Content: "test",
    }
    err = db.Write("novel/bookname/1.json", chapter1)
    if err != nil {
        panic(err)
    }

    // 读取数据
    err := db.Read("novel/bookname/1.json", &chapter1)
    if err != nil {
        panic(err)
    }
    println(chapter1.Title)

    // 删除数据
    err = db.Delete("novel/bookname/1.json")
    if err != nil {
        panic(err)
    }

    items, err := db.SearchDB("novel")
    if err != nil {
        panic(err)
    }
    println(items)
}
```

**剩余方法**请查看[测试](db_test.go)或源码