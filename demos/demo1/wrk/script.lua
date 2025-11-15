response = function(status, headers, body)
	if body ~= "Banana" and body ~= "Apple" then
		io.stderr:write("!!! INVALID RESPONSE: " .. body .. "\n")
	end
end