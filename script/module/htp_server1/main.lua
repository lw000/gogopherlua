package.path = package.path .. ";./script/?.lua;./script/module/htp_server1/?.lua"

require("global")
local web = require("web")

function main()
    log.info(gin:middleware(web.middleware.valid))
    log.info("/", gin:get("/", web.index))
    log.info("/get", gin:get("/get", web.get))
    log.info("/post", gin:post("/post", web.post))
    gin:run(8081)
end

local status, msg = xpcall(main, __G__TRACKBACK__)
if not status then
    print(msg)
end
