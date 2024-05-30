local wrk = require("wrk")
local json = require("json")

wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["Cookie"] = "YouNoteJWT=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY3NTAwOTYsImlkIjoiYzM2YjQzODAtNmNjMS00MmU4LWFkZmMtY2RmNTQ1ZWUxZmUzIiwidXNyIjoiZWxhc3RpYyJ9.eTSdPGHG8Oqx75wqwqG6-VskliosipOBqoCaFG5Iu9w; Path=/; Secure; HttpOnly; Expires=Sun, 26 May 2024 19:01:36 GMT;"

request = function()
    math.randomseed(os.time())
    local order = { items = {} }
    local numItems = math.random(1, 5)
    for i = 1, numItems do
        table.insert(order.items, {
            productID = math.random(1, 10),
            count = math.random(1, 10)
        })
    end
    local body = json.encode(order)
    return wrk.format(nil, "/api/v1/order/create", nil, body)
end

