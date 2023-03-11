API規劃:
  GET /head: 取得第一頁key, input: title
  GET /page: 取得內容+下一頁key, input: key
  POST /head: 建立新列表
  PATCH /page: 修改page指向
  DELETE /head: 刪除列表