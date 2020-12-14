from collections import Counter

seat_map = []

with open('test.txt') as fp:
   line = fp.readline()
   map_row = []
   while line:
     seat_map.append(line[:-1])
     line = fp.readline()

def countAdajcentSeats(grid, x, y):
    neighbors = []
    rows, cols = len(grid[0]), len(grid)

    if x > 0: # not top row
        neighbors.append(grid[x-1][y]) # up
        if y > 0:
            neighbors.append(grid[x-1][y-1]) # up left
        if y < rows - 1:
            neighbors.append(grid[x-1][y+1]) # up right
    if x < cols-1: # not bottom row
        neighbors.append(grid[x+1][y]) # down
        if y > 0:
            neighbors.append(grid[x+1][y-1]) # down left
        if y < rows - 1:
            neighbors.append(grid[x+1][y+1]) # down right
    if y > 0:
        neighbors.append(grid[x][y-1]) # left
    if y < rows - 1:
        neighbors.append(grid[x][y+1]) # right

    print(neighbors)
    print(Counter(neighbors))

for seat in seat_map:
    print(seat)

print("---")
countAdajcentSeats(seat_map,0,0)
countAdajcentSeats(seat_map,0,9)
countAdajcentSeats(seat_map,9,0)
countAdajcentSeats(seat_map,9,9)
