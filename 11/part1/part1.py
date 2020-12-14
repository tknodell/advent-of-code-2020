from collections import Counter
import copy

seat_map = []

with open('test.txt') as fp:
   line = fp.readline()
   map_row = []
   while line:
     seat_map.append(list(line[:-1]))
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

    # print(neighbors)
    # print(Counter(neighbors))
    return Counter(neighbors)

# copy seat map so we can apply simultaneously
seat_map_original = copy.deepcopy(seat_map)

for i=0;i<1;i++:
    for row in range(len(seat_map_original)):
        for col in range(len(seat_map_original[0])):
            currentSeat = seat_map_original[row][col]
            neighborCount = countAdajcentSeats(seat_map_original, row, col)

            # print(currentSeat,neighborCount)
            if currentSeat=="L" and neighborCount["#"]==0:
                # print("occuping", row,col)
                seat_map[row][col]= "#"

            # if currentSeat=="#" and neighborCount["#"]>=4:
            #     print("clearing", row,col)

for seat in seat_map:
    print(seat)
