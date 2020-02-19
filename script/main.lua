package.path = package.path .. ";./script/?.lua;./script/common/?.lua"
require("init")
local log = require("glog")
--local pb = require("script.luapb.pb")

-- 请求加载proto文件
--sp_person = pb.require("script.protos.Person")

local Game = class("Game")
function Game:ctor()
    self.__value = 10
    log.info("ctor", self.__value)
end

function test()
    log.info({a = 1, b = 2, c = {a = 1, b = 2, c = {a = 1, b = 2}}, d = {0, 1, 2, 3, 4, 5, 6, 7}})
    log.info(string.trim("   12313231 123132131 1231231231   "))

    for i = 1, 100 do
        local v = i / 3.5
        print("round_four", v, round_four(v))
        print("math.round", v, math.round(v))
    end

    for i = -100, 0 do
        local v = i / 2.5
        print("round_four", v, round_four(v))
        print("math.round", v, math.round(v))
    end

    local tstart = os.time()
    local now = os.time()
    log.info("os.time()", os.time())
    log.info('os.date("*t")', os.date("*t", 1565326338))
    log.info('os.date("!")', os.date("!"))
    log.info("os.date()", os.date())
    log.info(os.date("!", os.time()))
    log.info(os.date("*t"), os.date("*t"), os.date("*t"))

    local localTime = os.date("*t")
    dump(localTime)
    log.info("localTime", localTime)
    local nowDate = os.date("*t", now)
    log.info("nowDate", nowDate)
    local weekth = nowDate.wday
    log.info("weekth", weekth)
    log.info("weekth", os.date("%w", now))
    if 0 == weekth then
        weekth = 7
    end
    log.info("weekth", weekth)
    local diff = weekth - 1
    log.info("diff", diff)
    local monday = os.date("*t", now)
    log.info("monday", monday)
    monday.day = monday.day - diff
    log.info("monday", monday)

    -- dump(now)
    -- dump(os.time(monday))
    -- dump(os.time(localTime))
    -- dump(os.date("%Y-%m-%d %H:%M:%S", os.time(monday)))
    -- dump(os.date("%Y-%m-%d %H:%M:%S", os.time(localTime)))
    local bbb = true
    if not bbb then
        log.info(bbb)
    end

    local aaaaaa = nil
    if cccccc == nil then
        log.waring(aaaaaa)
    -- aaaaaa[1] = "131132"
    end

    local cccccc = {a = 111}
    if cccccc ~= nil then
        log.info(cccccc)
        cccccc[1] = {1, 2, 3, 4, 5, 6, 6, 7}
        cccccc[3] = {1, 2, 3, 4, 5, 6, 6, 7}
        dump(cccccc)

        for k, v in pairs(cccccc) do
            log.info(k, v)
        end
    end

    local tend = os.time()
    log.info(tend - tstart)
end

function main()
    --log.info(
    --    "主服务模块启动",
    --    os.date("%Y-%m-%d %H:%M:%S", os.time()),
    --    {
    --        a = 111,
    --        b = 2222,
    --        c = "asdfdfsf",
    --        d = true,
    --        e = false,
    --        f = {
    --            a = 111,
    --            b = 2222,
    --            c = "asdfdfsf",
    --            d = true,
    --            e = false,
    --            f = {a = 111, b = 2222, c = "asdfdfsf", d = true, e = false}
    --        }
    --    }
    --)

    log.info("统计开始", {now = os.date("%Y-%m-%d %H:%M:%S", os.time()), msg = "stat new round"})

    dump(os.date("*t", os.time()))

    game = Game:new()
    dump(game)

    --game1 = clone(game)
    --
    --print(io.readfile("script/common/core.lua"))

    --test()
end

local status, msg = xpcall(main, __G__TRACKBACK__)
if not status then
    print(msg)
end
