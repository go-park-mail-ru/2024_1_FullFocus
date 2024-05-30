local wrk = require("wrk")

wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["Cookie"] = "session_id=96a51ab1-181b-4ef0-94b6-eda4db5c9f16"

request = function()
    math.randomseed(os.time())
    local body = string.format('{"items": [{"productID": %d, "count": %d} ]}', math.random(1, 10), math.random(1, 10))
    return wrk.format(nil, "/api/v1/order/create", nil, body)
end

