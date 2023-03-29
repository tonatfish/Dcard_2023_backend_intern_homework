API規劃:
  GET /head/:id: 取得第一頁key, input: title
  GET /page/:id: 取得內容+下一頁key, input: key
  POST /list/:id: 建立新列表(head為id)

需求分析:
  1. 資料是key-value pair
  2. 資料會快速更新
  3. 資料每天須清除

假設:
  因為是根據使用者推薦 假設題目的"some-list-key"是user_id
  articles假設為一串article id的陣列形式儲存

解決方案:
  使用redis可供快速的查詢和簡單的key-value型態 並且可以設定一天後自動清除
  內部function會產出隨機長度為15的英數混和id 幾乎保證不會重複
  POST預計回傳為202及簡單的status:OK
  GetHead預計回傳200及題目所示輸出 或者404表示沒找到
  GetPage預計回傳200及題目所示輸出 或者404表示沒找到

難點:
  文中說到每小時更新列表 提到使用量可能很大 這點在redis中可能會有問題 redis在ram上的設計讓他天生快但小

改良方案:
  使用PostgreSQL + Redis 分層設計 將每個使用者的前幾頁 (通常會最常使用) 儲存在redis中 往後的頁數則儲存在PostgreSQL上
  這個方案可以實現使用者的快速讀取 並且也能符合大使用量的情境
  然而這個方案由於PostgreSQL沒有自帶的expire time 因此怎麼讓他自動清除會是一個問題

資料庫規劃: simple_article: article的資訊(id, title, description), page: page資訊(key, article_ids), user: user_id for head

使用框架: gin

please install and build redis first following the instructions here: https://www.runoob.com/docker/docker-install-redis.html

set想法:
 input可能: 二維陣列 一維陣列