
# test numbers
# player1Hand=[9,2,6,3,1]
# player2Hand=[5,8,4,7,10]

# puzzle
player1Hand=[6,25,8,24,30,46,42,32,27,48,5,2,14,28,37,17,9,22,40,33,3,50,47,19,41]
player2Hand=[1,18,31,39,16,10,35,29,26,44,21,7,45,4,20,38,15,11,34,36,49,13,23,43,12]

roundCount=0

def calcScore(hand):
    score=0
    hand.reverse()
    for index,card in enumerate(hand):
        score += (card * (index+1))
    return score

while len(player1Hand)!=0 and len(player2Hand)!=0:
    playerOneCard = player1Hand[0]
    playerTwoCard = player2Hand[0]

    print(roundCount, playerOneCard,playerTwoCard)
    if (playerOneCard>playerTwoCard):
        print("player1 wins with", playerOneCard)
        player1Hand.append(playerOneCard)
        player1Hand.append(playerTwoCard)
        player1Hand.pop(0)
        player2Hand.pop(0)
    else:
        print("player2 wins with", playerTwoCard)
        player2Hand.append(playerTwoCard)
        player2Hand.append(playerOneCard)
        player1Hand.pop(0)
        player2Hand.pop(0)

    # print(player1Hand,player2Hand)
    roundCount=roundCount+1

print("\nresults")
if len(player1Hand)==0:
    finalScore=calcScore(player2Hand)
else:
    finalScore=calcScore(player1Hand)

print(player1Hand, player2Hand)
print(finalScore)
