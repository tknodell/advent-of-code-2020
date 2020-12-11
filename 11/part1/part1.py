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
    # print(currentSeatX,currentSeatY)
    # print(map[currentSeatX][currentSeatY])
    rows = len(map)
    seats = len(split(map[currentSeatX-1]))

    print("num of rows",rows)
    print("num of seats in row #",currentSeatY,seats)

    # up
    if currentSeatX==0:
        print("on top row, skipping up")
    else:
        print(map[currentSeatX-1][currentSeatY])

    # down
    if currentSeatY==rows:
        print("on bottom row, skipping down")
    else:
        print(map[currentSeatX+1][currentSeatY])

    # left
    if currentSeatX == 0:
        print("on leftmost row, skipping")
    else:
        print(map[currentSeatX][currentSeatY-1])

    # right
    if currentSeatX == seats:
        print("on rightmost row, skipping")
    else:
        print(map[currentSeatX][currentSeatY+1])


    # up-left
    if currentSeatX==0:
        print("on top row, skipping")
    else:
        print(map[currentSeatX-1][currentSeatY-1])

    # up-right
    if currentSeatX==0:
        print("on top row, skipping")
    else:
        print(map[currentSeatX-1][currentSeatY+1])

    # down-left
    if currentSeatX+1==rows:
        print("on bottom row, skipping")
    else:
        print(map[currentSeatX+1][currentSeatY-1])

    # down-right
    if currentSeatX+1==rows:
        print("on bottom row, skipping")
    else:
        print(map[currentSeatX+1][currentSeatY+1])



for seat in seat_map:
    print(seat)

print("---")
# countAdajcentSeats(seat_map,0,0)
# countAdajcentSeats(seat_map,0,10)
# countAdajcentSeats(seat_map,10,0)
countAdajcentSeats(seat_map,10,10)
