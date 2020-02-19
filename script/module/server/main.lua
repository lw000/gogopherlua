package.path = package.path .. ";./script/?.lua;./script/module/server/?.lua;"

local log = require("glog")
require("common.trackback")

function main()
    loop(1000)
    log.info("loop")
    sleep(5000)
    log.info("sleep")
    after(
        10000,
        function()
            log.info("after", os.date("*t"))
        end
    )
end

local status, msg = xpcall(main, __G__TRACKBACK__)
if not status then
    print(msg)
end
