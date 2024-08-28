local pi_leibniz = function(n)
    if (n == 0) then
        return n
    end
  
    local pi = 0
    local sign = 1
    local i = 1
    
    while i <= (n * 2) do
        pi = pi + sign * (4 / i)
        sign = -sign
        i = i + 2
    end
    
    return pi
end

local pi_wallis = function(n)
    if (n == 0) then
        return n
    end
    
    local pi = 4.
    local i = 3
    
    while i <= (n + 2) do
        pi = pi * ((i - 1) / i) * ((i + 1) / i)    
        i = i + 2
    end
    
    return pi
end

local pi_nilakantha = function(n)
    if (n == 0) then
        return n
    end
    
    local pi = 3
    local sign = 1
    local i = 2
    
    while i <= (n * 2) do
        pi = pi + sign * (4 / (i * (i + 1) * (i + 2)))
        sign = -sign
        i = i + 2
    end
    
    return pi
end