package.path = package.path .. ";./script/?.lua;./script/module/htp_server0/?.lua"

require("global")

local web = require("web")
--local GameLogic = require("game.game")
--local Shape = require("common.Shape")
--local Cacl = require("common.Cacl")

function update(d)
    --print(d)
    -- local aa
    -- aa = nil
    -- aa["aaaa"] = "111111"
    --aaaaa = 1
    --print(aaaaa/0)
    --
    --for a= 1,10 do
    --    print(a)
    --end
end

function testGameLogic()
    logic = GameLogic:new({}, 10)
    logic:start()
    logic:stop()
    logic:update()

    shape = Shape:new({}, 100)
    shape:printArea()
end

function main()
    -- print("max: " .. Cacl.max(100, 200))
    -- print("min: " .. Cacl.min(100, 200))
    -- a, b = Cacl.maxmin(100, 200)
    -- print("maxmin: " .. a .. ", " .. b)

    --testGameLogic()

    log.info("/", gin:get("/", web.index))
    log.info("/get", gin:get("/get", web.get))
    log.info("/post", gin:post("/post", web.post))
    gin:run(8080)
end

local status, msg = xpcall(main, __G__TRACKBACK__)
if not status then
    print(msg)
end
