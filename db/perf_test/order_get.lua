local wrk = require("wrk")

wrk.method = "GET"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["Cookie"] = "session_id=90ae8f3c-7f13-43fe-aed3-72ae9dc48105"

request = function()
    math.randomseed(os.time())
    local path = string.format("/api/v1/order/%d", math.random(100, 10000))
    return wrk.format(nil, path, nil, nil)
end