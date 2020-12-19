lines = []

with open('input.txt') as fp:
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
validCols=[]

myTicketValues = myTicket.split("\n")[1]
myTicketSplit = [int(x) for x in myTicketValues.split(",")]

def getValidLines(nums,tickets):
    lines = []
    invalid = False
    for i,v in enumerate(nearbyTickets.split("\n")):
        if i!=0: # skip first line
            invalid = False
            nearbyNums = v.split(',')
            for num in map(int, nearbyNums): # use map to convert num to int
                if num not in validNums:
                    invalid = True
            if invalid != True:
                lines.append(v)
    return lines

def findMatchColumn(validLines, min1,max1,min2,max2):
    position=0
    possiblePositions = []
    while position <20:
        validColumn = True
        for line in validLines:
            col = line.split(",")
            value = int(col[position])
            # print(value,position,min1,max1,min2,max2)
            if not ((min1 <= value <= max1) or (min2 <= value <= max2)):
                # print("num is NOT valid", value)
                validColumn = False
        if validColumn:
            possiblePositions.append(position)
        position=position+1
    return possiblePositions

# populate valid numbers
for rule in rules.split("\n"):
    ruleName = rule.split(":")[0]
    rulesSplit = rule.split(":")[1]
    firstRange = rulesSplit.split("or")[0].split(" ")[1]
    secondRange = rulesSplit.split("or")[1].split(" ")[1]

    # print(ruleName,firstRange, secondRange)
    for v in (firstRange,secondRange):
        nums = v.split('-')
        for num in range(int(nums[0]),int(nums[1])+1):
            if num not in validNums:
                validNums.append(num)
    validNums.sort()

validLines = getValidLines(validNums,nearbyTickets)

# part2

for rule in rules.split("\n"):
    ruleName = rule.split(":")[0]
    rulesSplit = rule.split(":")[1]
    firstRange = rulesSplit.split("or")[0].split(" ")[1]
    secondRange = rulesSplit.split("or")[1].split(" ")[1]

    if ruleName.startswith("departure"):
        firstRangeSplit = firstRange.split("-")
        min1,max1=firstRangeSplit[0],firstRangeSplit[1]
        
        secondRangeSplit = secondRange.split("-")
        min2,max2=secondRangeSplit[0],secondRangeSplit[1]

        possibleColumnPositions = findMatchColumn(validLines,int(min1),int(max1),int(min2),int(max2))
        print(ruleName,possibleColumnPositions)

# had to manually figure out the result
