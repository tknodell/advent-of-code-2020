lines = []

with open('test.txt') as fp:
    line = fp.readline()
    map_row = []
    while line:
        lines.append(line[:-1])
        line = fp.readline()
 
departTime = 1068781
schedules = lines[1].split(",")
validSched = []

for s in schedules:
    if s != "x":
        validSched.append(int(s))

startTime = 0
maxTime = int(departTime)+10
i=0

print(validSched)
for i in range(departTime-20,maxTime+20):
    print(i)
 
    for sched in validSched:
        if i%sched==0:
            print("bus {} departing {} after the hour, have to wait {} minutes. id is {}".format(sched,i%60,i-departTime, (i-departTime)*sched))
            # exit(0)
    i = i+1
