seat_map = []

with open('test.txt') as fp:
   line = fp.readline()
   map_row = []
   while line:
     seat_map.append(line[:-1])
     line = fp.readline()

def traverse(dx, dy):
  x, y, tree_count = 0, 0, 0
  while y < len(seat_map):
    if seat_map[y][x % len(seat_map[0])] == '#':
      tree_count = tree_count + 1
    x = x + dx 
    y = y + dy
  return tree_count

def split(line): 
    return [char for char in line]  


def countAdajcentSeats(map, currentSeatX, currentSeatY):
    print(map[currentSeatX][currentSeatY])
    rows = len(map)
    seats = len(split(map[currentSeatX]))

    print("num of rows",rows)
    print("num of seats",seats)

    # cols = len(map[currentSeatX].split())
    # print("num of cols",cols)

    # up
    # down
    # left
    # right

    # up-left
    # up-right
    # down-left
    # down-right

for seat in seat_map:
    print(seat)

print("---")
countAdajcentSeats(seat_map,0,0)
