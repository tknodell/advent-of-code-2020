lines = []

with open('test.txt') as fp:
    line = fp.readline()
    map_row = []
    while line:
        lines.append(line[:-1])
        line = fp.readline()
 
departTime = int(lines[0])
schedules = lines[1].split(",")
validSched = []

for s in schedules:
    if s != "x":
        validSched.append(int(s))

startTime = 0
maxTime = int(departTime)+10
i=0

print(validSched)
for i in range(departTime-10,maxTime+10):
    print(i)
    minsAfter=i%60
 
    for sched in validSched:
        if i%sched==0:
            print("bus {} departing {} after the hour, have to wait {} minutes".format(sched,i%60,i-departTime))

    # print(i,minsAfter)
    # if minsAfter in validSched:
    #     print("bus {} departing at {}").format(minsAfter,i)
    i = i+1
