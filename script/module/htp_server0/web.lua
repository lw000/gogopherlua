local crypto = require("crypto")

local web = {}

function web.index(ctx)
    dump(ctx)
    log.debug(ctx)

    return {c = 1, m = "index", d = {module = "module0"}, ctx = ctx}
end

function web.get(ctx)
    dump(ctx)
    log.debug(ctx)

    local ext1 = {}
    ext1[0] = 11111
    ext1[1] = 11111
    ext1[2] = 11111
    ext1[3] = 11111

    return {
        c = 1,
        m = "get",
        d = {module = "module0"},
        ctx = ctx,
        ext = {0, 1, 3, 4, 5, 6},
        ext1 = ext1,
        md5 = crypto.md5("111111111111111111111")
    }, 1, 2, 3
end

function web.post(ctx)
    dump(ctx)
    log.debug(ctx)

    return {
        c = 1,
        m = "post",
        d = {module = "module0"},
        ctx = ctx,
        test = {0, 1, 3, 4, 5, 6},
        md5 = crypto.md5("111111111111111111111")
    }
end

return web
