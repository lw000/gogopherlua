local web = {
    middleware = {}
}

local function testhttp()
    local response, err =
        http.request(
        "GET",
        "http://serv.acgcy.com/api/h5",
        -- "http://www.baidu.com",
        {
            query = "page=1",
            headers = {
                Accept = "*/*"
            }
        }
    )
    if err ~= nil then
        log.error(err)
        return
    end

    -- dump(response)
    -- dump(response.body)
    
    log.info(response.body)
end

function web.middleware.valid(data)
    log.debug(data)
    dump(data)
    -- return false, "禁止访问"
    return true, "允许访问"
end

function web.index(ctx)
    dump(ctx)
    --testhttp()
    return {c = 1, m = "", d = {method = "index"}, ctx = ctx}
end

function web.get(ctx)
    dump(ctx)
    --sleep(1000)
    return {c = 1, m = "", d = {method = "get"}, ctx = ctx}
end

function web.post(ctx)
    dump(ctx)
    return {c = 1, m = "", d = {method = "post"}, ctx = ctx}
end

return web
