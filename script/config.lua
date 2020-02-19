package.path = package.path .. ";./script/?.lua;./script/common/?.lua"
require("init")

__TY_CONFIG__ = {
    debug = 1 -- 0=release模式
}
dump(__TY_CONFIG__, "__TY_CONFIG__")

__TY_CHILD_MODULES_PATH__ = {
    "script/module/htp_server0/main.lua",
    --"script/module/htp_server1/main.lua",
    --"script/module/schedule_server/main.lua",
    --"script/module/server/main.lua",
    --"script/module/redis_test/main.lua",
    "script/module/cron_test/main.lua",
    --"script/main.lua",
}
dump(__TY_CHILD_MODULES_PATH__, "__TY_CHILD_MODULES_PATH__")
