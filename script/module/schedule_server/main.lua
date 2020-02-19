package.path = package.path .. ";./script/?.lua;./script/module/schedule_server/?.lua;"

require("global")

function testcrypto()
    log.info("md5", crypto.md5("111111111111111"))
    log.info("crc32", crypto.crc32("111111111111111"))
    log.info("sha1", crypto.sha1("111111111111111"))
    log.info("sha256", crypto.sha256("111111111111111"))
    log.info("sha512", crypto.sha512("111111111111111"))
    log.info("hmac", crypto.hmac("md5", "111111111111111", "1111111111"))
end

function testlog()
    log.info(1111111)
    log.info(123.123)
    log.info(true)
    log.info({a = 100, b = 200, c = 300})
    log.info({a = 1, b = 2, c = 3}, "InfoInfoInfoInfoInfoInfo")
    log.waring("WaringWaringWaringWaringWaringWaringWaringWaring")
    log.error("ErrorErrorErrorErrorErrorErrorErrorErrorError")
end

function testurl()
    parsed_url = url.parse("http://example.com/")
    dump(parsed_url)
    print(parsed_url.host)
end

function testhttp()
    response, error_message =
        http.request(
        "GET",
        "http://serv.acgcy.com/api/h5",
        {
            query = "page=1",
            headers = {
                Accept = "*/*"
            }
        }
    )
    dump(response)
    log.info(response.body)
end

function testJson()
    s = json.encode({a = 1111, b = 222222, c = 333333333})
    dump(s)
    ss = json.decode(s)
    dump(ss)
end

local index = 1
function scheduleTest(data)
    index = index + 1
    log.info("schedule", index, os.date("*t", os.time()))
    -- dump(data)
end

function execStat()
    index = index + 1
    log.info(index, os.date("*t", os.time()))
end

function main()
    testlog()
    testcrypto()
    testurl()
    testhttp()
    testJson()

    global.Cron:start()
    for i = 0, 1000 do
        global.Cron:add("*/1 * * * * ?", execStat)
    end

    global.Cron1:start()
    for i = 0, 1000 do
        global.Cron1:add("*/1 * * * * ?", execStat)
    end

    global.Schedule:start()
    for i = 0, 10000 do
        global.Schedule:add(1, scheduleTest, {a = "aaaaaaaaaa", b = "bbbbbbbbbb", c = "dddddddddd"})
    end

    -- global.Loop:loop()
end

local status, msg = xpcall(main, __G__TRACKBACK__)
if not status then
    print(msg)
end
