local GameLogic = {
    PlayerCount = 0
}

function GameLogic:new(o, playerCount)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    self.PlayerCount = playerCount
    return o
end

function GameLogic:start()
    print("GameLogic:start " .. self.PlayerCount)
end

function GameLogic:stop()
    print("GameLogic:stop")
end

function GameLogic:update()
    print("GameLogic:update")
end

return GameLogic
