lines = []

with open('test.txt') as fp:
    line = fp.readline()
    map_row = []
    while line:
        lines.append(line[:-1])
        line = fp.readline()
 
departTime = lines[0]
schedules = lines[1].split(",")
validSched = []

for s in schedules:
    if s != "x":
        validSched.append(s)




print(departTime)
print(validSched)
