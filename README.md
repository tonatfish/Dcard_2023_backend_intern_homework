API規劃:
  GET /head: 取得第一頁key, input: title
  GET /page: 取得內容+下一頁key, input: key
  POST /list/:id: 建立新列表(head為id)
  PATCH /page: 修改page指向
  DELETE /head: 刪除列表

關鍵需求:
每日清除 -> expire
GO
unit test
restful api
大使用量 -> 列表快速多次更新
key-value
postgresql?
linked list

資料庫規劃: simple_article: article的資訊(id, title, description), page: page資訊(key, article_ids), user: user_id for head

使用框架: gin

please install and build redis first following the instructions here: https://www.runoob.com/docker/docker-install-redis.html

set想法:
 input可能: 二維陣列 一維陣列