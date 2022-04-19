function OnBeforeCompilation()
    print("-----------------------  Attached Protection  -----------------------\n")
         local file = vmprotect.core():outputArchitecture()
	 for i = 1, file:mapFunctions():count() do
	    local fn = file:mapFunctions():item(i)
            local addType
	    if fn:name():find("Cfunc_VMProtect") then
	       addType = CompilationType.Ultra
	       file:functions():addByAddress(fn:address(), addType)
	       print(fn:name() .. " : " .. tostring(fn:address()) .. "  =>  Add To Protection\n")
	    end
         end
    print("-----------------------  Attached Protection  -----------------------\n")
end

function GetRandomName()
	local res = ""
	for i = 1, 8 do
		res = res .. string.char(32 + math.random(string.byte("z") - 32))
	end
	return res
end

function OnAfterSaveFile()
    local file = vmprotect.core():outputArchitecture()
    if vmprotect.core():inputFile():format() == "PE" then
    print("-----------------------  Obfuscate Segment Name  -----------------------\n")
	for i = 1, file:segments():count() do
		segment = file:segments():item(i)
		name = GetRandomName()
		segment:setName(name)
		print(string.format("Segment \"%s\" is renamed to \"%s\"", segment:name(), name))
	end
    print("-----------------------  Obfuscate Segment Name  -----------------------\n")
    end
end