--[[
    dump是一个用于调试输出数据的函数，能够打印出nil,boolean,number,string,table类型的数据，以及table类型值的元表
    参数data表示要输出的数据
    参数showMetatable表示是否要输出元表
    参数lastCount用于格式控制，用户请勿使用该变量
]]
--
--function dump(data, showMetatable, lastCount)
--    if type(data) ~= "table" then
--        --Value
--        if type(data) == "string" then
--            io.write('"', data, '"')
--        else
--            io.write(tostring(data))
--        end
--    else
--        --Format
--        local count = lastCount or 0
--        count = count + 1
--        io.write("{\n")
--        --Metatable
--        if showMetatable then
--            for i = 1, count do
--                io.write("\t")
--            end
--            local mt = getmetatable(data)
--            io.write('"__metatable" = ')
--            dump(mt, showMetatable, count) -- 如果不想看到元表的元表，可将showMetatable处填nil
--            io.write(",\n") --如果不想在元表后加逗号，可以删除这里的逗号
--        end
--        --Key
--        for key, value in pairs(data) do
--            for i = 1, count do
--                io.write("\t")
--            end
--            if type(key) == "string" then
--                io.write('"', key, '" = ')
--            elseif type(key) == "number" then
--                io.write("[", key, "] = ")
--            else
--                io.write(tostring(key))
--            end
--            dump(value, showMetatable, count) -- 如果不想看到子table的元表，可将showMetatable处填nil
--            io.write(",\n") --如果不想在table的每一个item后加逗号，可以删除这里的逗号
--        end
--        --Format
--        for i = 1, lastCount or 0 do
--            io.write("\t")
--        end
--        io.write("}")
--    end
--    --Format
--    if not lastCount then
--        io.write("\n")
--    end
--end

 local function dump_value_(v)
     if type(v) == "string" then
         v = "\"" .. v .. "\""
     end
     return tostring(v)
 end

 function dump(value, description, nesting)
     if type(nesting) ~= "number" then nesting = 3 end

     local lookupTable = {}
     local result = {}

     local traceback = string.split(debug.traceback("", 2), "\n")
     print("dump from: " .. string.trim(traceback[3]))

     local function dump_(value, description, indent, nest, keylen)
         description = description or "<var>"
         local spc = ""
         if type(keylen) == "number" then
             spc = string.rep(" ", keylen - string.len(dump_value_(description)))
         end
         if type(value) ~= "table" then
             result[#result +1 ] = string.format("%s%s%s = %s", indent, dump_value_(description), spc, dump_value_(value))
         elseif lookupTable[tostring(value)] then
             result[#result +1 ] = string.format("%s%s%s = *REF*", indent, dump_value_(description), spc)
         else
             lookupTable[tostring(value)] = true
             if nest > nesting then
                 result[#result +1 ] = string.format("%s%s = *MAX NESTING*", indent, dump_value_(description))
             else
                 result[#result +1 ] = string.format("%s%s = {", indent, dump_value_(description))
                 local indent2 = indent.."    "
                 local keys = {}
                 local keylen = 0
                 local values = {}
                 for k, v in pairs(value) do
                     keys[#keys + 1] = k
                     local vk = dump_value_(k)
                     local vkl = string.len(vk)
                     if vkl > keylen then keylen = vkl end
                     values[k] = v
                 end
                 table.sort(keys, function(a, b)
                     if type(a) == "number" and type(b) == "number" then
                         return a < b
                     else
                         return tostring(a) < tostring(b)
                     end
                 end)
                 for i, k in ipairs(keys) do
                     dump_(values[k], k, indent2, nest + 1, keylen)
                 end
                 result[#result +1] = string.format("%s}", indent)
             end
         end
     end
     dump_(value, description, "- ", 1)

     for i, line in ipairs(result) do
         print(line)
     end
 end