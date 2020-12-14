import copy

lines = []

with open('test.txt') as fp:
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

def rotateWaypoint(waypoint, angle):
    waypointResult = copy.deepcopy(waypoint)

    if angle>0:
        dir="clockwise"
    else:
        dir="counterclockwise"

    numRotations = angle/90
    for _ in range(numRotations):
        if dir=="clockwise":
            print("rotating 90 clockwise")
            # north to east
            waypointResult["east"] = waypoint["north"]
            # east to south
            waypointResult["south"] = waypoint["east"]
            # south to west
            waypointResult["west"] = waypoint["south"]
            # west to north
            waypointResult["north"] = waypoint["west"]
        if dir=="counterclockwise":
            print("rotating 90 counter-clockwise")
            # north to west
            waypointResult["west"] = waypoint["north"]
            # west to south
            waypointResult["south"] = waypoint["west"]
            # south to east
            waypointResult["east"] = waypoint["south"]
            # east to north
            waypointResult["north"] = waypoint["east"]

    return waypointResult

def manhattanDistance(pos):
    x = abs(pos["east"]-pos["west"])
    y = abs(pos["south"]-pos["north"])

    return x+y

initDirection = "east"
position = {"north": 0, "east": 0, "south": 0, "west": 0}
waypoint = {"north": 1, "east": 10, "south": 0, "west": 0}

for line in lines:
    # print(line)
    direction = line[:1]
    steps = int(line[1:])

    # rotate by angle
    if direction == "R":
        waypoint = rotateWaypoint(waypoint,steps)
    if direction == "L":
        waypoint = rotateWaypoint(waypoint,-steps)

    # move number of steps
    if direction == "F":
        for point in waypoint:
            # print(point, waypoint[point], waypoint[point]*steps)
            position[point] += waypoint[point]*steps
    if direction == "N":
        waypoint["north"]+=steps
    if direction == "E":
        position["east"] *= steps
    if direction == "S":
        position["south"] *= steps
    if direction == "W":
        position["west"] *= steps

    print("---")
    print(line)
    print("position",position)
    print("waypoint:",waypoint)

print(manhattanDistance(position))
