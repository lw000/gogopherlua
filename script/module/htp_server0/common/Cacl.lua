local Cacl = {
    double = function (a)
        return a * 2
    end,
    
    maxmin = function ( a, b )
        if a > b then
            return a, b
        else
            return b, a
        end
    end,
    
    max = function (a, b)
        if a > b then
            return a
        else
            return b
        end
    end,
    
    min = function(a, b)
        if a > b then
            return b
        else
            return a
        end
    end
}

return Cacl