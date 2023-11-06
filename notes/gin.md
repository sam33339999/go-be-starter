# GIN

## URL PARAMETER:
    - Query
    - Parameter in Path


### Query
> 例如 /list?id=2 這種的為 url query parameters ，送到後端時，去 parser id=2 這個資訊。
(c *gin.Context)

id := c.Query("id")
parseId, err := strconv.ParseUint(id, 10, 32)



example:
- /api/mock-login -> api/routes/auth.go
