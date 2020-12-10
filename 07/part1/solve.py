import pprint

def containsTargetBag(targetName, insideBags):
    for child in insideBags:
        if child == targetName:
            return True
        if containsTargetBag(targetName, data[child]):
            return True
    return False

def countInsideBags(insideBags):
    if insideBags == {}:
        return 0
    total = 0
    for insideBagName in insideBags:
        total += ((countInsideBags(data[insideBagName]) + 1) * insideBags[insideBagName])
    return total

# input = "test.txt"
input = "input.txt"


# Get Input
# data will be stored as dictionary:
#   Key = Color Name
#   Value = Another Dictionary, with <color names:quantity> entries
data = {}
with open(input, 'r') as file:
    for line in file.readlines():
        words = line.split()
        key = words[0] + ' ' + words[1]
        data[key] = {}
        for i, word in enumerate(words):
            if words[i].isnumeric():
                subKey = words[i + 1] + ' ' + words[i + 2]
                data[key][subKey] = int(word)

pp = pprint.PrettyPrinter(indent=4)
pp.pprint(data)

answer = []
targetName = 'shiny gold'

for bag in data:
    if containsTargetBag(targetName, data[bag]):
        answer.append(bag)

print("SUM Parents of {}\t\t".format(targetName), len(answer))
print("TOTAL INSIDE bags of {}\t\t".format(targetName), countInsideBags(data[targetName]))
