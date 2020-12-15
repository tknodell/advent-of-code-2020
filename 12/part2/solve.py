import copy

lines = []

with open('input.txt') as fp:
    line = fp.readline()
    map_row = []
    while line:
        lines.append(line[:-1])
        line = fp.readline()


directions = {
    "north": 0,
    "east": 90,
    "south": 180,
    "west": 270,
}


def rotateWaypoint(waypoint, angle):
    waypointResult = copy.deepcopy(waypoint)

    if angle > 0:
        print("rotating clockwise", angle)
        if angle == 90:
            # north to east
            waypointResult["east"] = waypoint["north"]
            # east to south
            waypointResult["south"] = waypoint["east"]
            # south to west
            waypointResult["west"] = waypoint["south"]
            # west to north
            waypointResult["north"] = waypoint["west"]
        if angle == 180:
            # north to south
            waypointResult["south"] = waypoint["north"]
            # east to west
            waypointResult["west"] = waypoint["east"]
            # south to north
            waypointResult["north"] = waypoint["south"]
            # west to east
            waypointResult["east"] = waypoint["west"]
        if angle == 270:
            # north to west
            waypointResult["west"] = waypoint["north"]
            # east to north
            waypointResult["north"] = waypoint["east"]
            # south to east
            waypointResult["east"] = waypoint["south"]
            # west to south
            waypointResult["south"] = waypoint["west"]

    if angle < 0:
        print("rotating counter-clockwise", angle)
        if abs(angle) == 90:
            # north to west
            waypointResult["west"] = waypoint["north"]
            # east to north
            waypointResult["north"] = waypoint["east"]
            # south to east
            waypointResult["east"] = waypoint["south"]
            # west to south
            waypointResult["south"] = waypoint["west"]
        if abs(angle) == 180:
            # north to south
            waypointResult["south"] = waypoint["north"]
            # east to west
            waypointResult["west"] = waypoint["east"]
            # south to north
            waypointResult["north"] = waypoint["south"]
            # west to east
            waypointResult["east"] = waypoint["west"]
        if abs(angle) == 270:
            # north to east
            waypointResult["east"] = waypoint["north"]
            # east to south
            waypointResult["south"] = waypoint["east"]
            # south to west
            waypointResult["west"] = waypoint["south"]
            # west to north
            waypointResult["north"] = waypoint["west"]

    return waypointResult


def manhattanDistance(pos):
    x = abs(pos["east"]-pos["west"])
    y = abs(pos["south"]-pos["north"])

    return x+y


initDirection = "east"
position = {"north": 0, "east": 0, "south": 0, "west": 0}
waypoint = {"north": 1, "east": 10, "south": 0, "west": 0}

for line in lines:
    print(line)
    direction = line[:1]
    steps = int(line[1:])

    # rotate by angle
    if direction == "R":
        waypoint = rotateWaypoint(waypoint, steps)
    if direction == "L":
        waypoint = rotateWaypoint(waypoint, -steps)

    # move number of steps
    if direction == "F":
        for point in waypoint:
            # print(point, waypoint[point], waypoint[point]*steps)
            position[point] += waypoint[point]*steps
    if direction == "N":
        waypoint["north"] += steps
    if direction == "E":
        waypoint["east"] += steps
    if direction == "S":
        waypoint["south"] += steps
    if direction == "W":
        waypoint["west"] += steps

    print("position", position)
    print("waypoint:", waypoint)
    print("---")


print(manhattanDistance(position))

# 33756 too high
# 25580 not right
