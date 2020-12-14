lines = []

with open('input.txt') as fp:
    line = fp.readline()
    map_row = []
    while line:
        lines.append(line[:-1])
        line = fp.readline()


directions = {
    "north":0,
    "east":90,
    "south":180,
    "west":270,
}

def calcDirection(current, angle):
    directionAngle=directions[current]
    directionAngle+=angle
    directionAngle = directionAngle % 360

    for key, value in directions.items():
        if value == directionAngle:
            # print("New direction is", key)
            return key

def manhattanDistance(pos):
    x = abs(pos["east"]-pos["west"])
    y = abs(pos["south"]-pos["north"])

    return x+y

initDirection = "east"
position = {"north": 0, "east": 0, "south": 0, "west": 0}

currentDirection = initDirection

for line in lines:
    # print(line)
    direction = line[:1]
    steps = int(line[1:])

    # rotate by angle
    if direction == "R":
        currentDirection = calcDirection(currentDirection,steps)
    if direction == "L":
        currentDirection = calcDirection(currentDirection,-steps)

    # move number of steps
    if direction == "F":
        position[currentDirection] += steps
    if direction == "N":
        position["north"] += steps
    if direction == "E":
        position["east"] += steps
    if direction == "S":
        position["south"] += steps
    if direction == "W":
        position["west"] += steps

    print(position)

print(manhattanDistance(position))
