package.path = package.path .. ";./script/?.lua;./script/module/cron_test/?.lua"
---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2019/11/6 11:40
---

require("common.trackback")
local log = require("glog")
local gcron = require("gcron")

cron = gcron.new()

function execStat()
    log.info(os.date("%Y-%m-%d %H:%M:%S", os.time()))
end

function main()
    cron:start()
    for i = 1, 1000 do
        cron:add("*/1 * * * * ?", execStat)
    end
end

local status, msg = xpcall(main, __G__TRACKBACK__)
if not status then
    print(msg)
end