local M = {}

function M.upgrade_callback(c)
    if c.Method == "GET" then
        return true
    end
    return false
end

return M
