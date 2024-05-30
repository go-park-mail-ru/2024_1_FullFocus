local wrk = require("wrk")

wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["Cookie"] = "session_id=7c8c0db2-838d-4a3d-9506-745cb8dd42f9"

request = function()
    math.randomseed(os.time())
    local body = string.format('{"items": [{"productID": %d, "count": %d} ]}', math.random(1, 10), math.random(1, 10))
    return wrk.format(nil, "/api/v1/order/create", nil, body)
end

