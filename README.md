API規劃:
  GET /head/:id: 取得第一頁key, input: title
  GET /page/:id: 取得內容+下一頁key, input: key
  POST /list/:id: 建立新列表(head為id), input: {"data": string[][]}
  AddList預計回傳為202及簡單的status:OK
  GetHead預計回傳200及題目所示輸出 或者404表示沒找到
  GetPage預計回傳200及題目所示輸出 或者404表示沒找到

需求分析:
  1. 資料接近key-value pair
  2. 資料會頻繁更新
  3. 資料每天須清除

設計:
  因為是根據使用者推薦 假設題目的"some-list-key"是user_id
  articles假設為一串article id的陣列形式儲存
  內部function會產出隨機長度為15的英數混和id 幾乎保證不會重複
  由於是一天暫存形式 page是不斷新增而不是更新舊資料 (空間換取跑linked list的時間)

初始解決方案:
  使用redis可供快速的查詢和簡單的key-value型態 並且可以設定一天後自動清除
  Redis使用自動的存在一天即過期
  難點:
    文中說到每小時更新列表 提到使用量可能很大 這點在redis中可能會有問題 redis的設計讓他天生快但小

最終解決方案:
  使用PostgreSQL + Redis 分層設計 將每個使用者的前幾頁 (通常會最常使用) 儲存在redis中 往後的頁數則儲存在PostgreSQL上
  head也儲存在Redis上
  這個方案可以實現快速讀取 並且也能符合大使用量的情境
  PostgreSQL使用內建的pg_cron設定清除存在超過一天的資料
  Redis使用自動的存在一天即過期
  難點:
    規劃上PostgreSQL採用每日定時清除 而Redis則是自動一天後清除
    如此一來資料可能會有對不上的空窗期
    但這種情況可以回傳404並讓使用者從head重新取得即可

資料庫規劃:
  Redis:
    head:
      key: 'head_' + user id
      value: head資訊(nextPageKey)
    page:
      key: 'page_' + page id
      value: page資訊(article_ids, nextPageKey)
  PostgreSQL:
    page:
      id: page id
      data: page資訊(article_ids, nextPageKey)(json)
      created_at: timestamp for clean

使用框架: gin

使用docker-compose將使用到的資料庫包裝起來
內容包含:
  PostgreSQL: 利用sqls/init.sql設定資料庫
  Redis: 無須預先設定
  pgadmin: 用以檢查PostgreSQL的使用狀況

Before project, run docker with:
  docker-compose up

To run the project:
  go run .

To run the test:
  go test .\test\
