lines = []

with open('test.txt') as fp:
    line = fp.readline()
    map_row = []
    while line:
        lines.append(line[:-1])
        line = fp.readline()

parts = "\n".join(lines).split("\n\n")
rules = parts[0]
myTicket = parts[1]
nearbyTickets = parts[2]

validNums=[]

for rule in rules.split("\n"):
    ruleName = rule.split(":")[0]
    firstRange = rule.split("or")[0].split(" ")[1]
    secondRange = rule.split("or")[1].split(" ")[1]

    # print(ruleName,firstRange, secondRange)
    for v in (firstRange,secondRange):
        nums = v.split('-')
        for num in range(int(nums[0]),int(nums[1])+1):
            if num not in validNums:
                validNums.append(num)
    validNums.sort()

print(validNums)
