local wrk = require("wrk")

wrk.method = "GET"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["Cookie"] = "session_id=96a51ab1-181b-4ef0-94b6-eda4db5c9f16"

local orderId = 0
request = function()
    orderId = orderId + 1
    local path = string.format('/api/v1/order/%d', orderId)
    return wrk.format(nil, body, nil, nil)
end