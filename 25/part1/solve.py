

def transformSubjectNumber(subjectNumber, loopSize):
    value = 1
    for _ in range(0,loopSize):
        value = value * subjectNumber
        value = value % 20201227
    return value


def findLoopSize(publicKey):
    for subjectNumber in range(0,1000):
        for loopSize in range(0,1000):
            if transformSubjectNumber(subjectNumber,loopSize) == publicKey:
                return subjectNumber,loopSize

# Test inputs
# cardPublicKey=5764801 
# doorPublicKey=17807724

# My input
cardPublicKey=19774466 
doorPublicKey=7290641

print("Card public key",cardPublicKey)
cardSubject, cardLoop = findLoopSize(cardPublicKey)
print("FOUND")
print("subject number: {}, loop size: {}".format(cardSubject,cardLoop))
print("")


print("Door public key",doorPublicKey)
doorSubject, doorLoop = findLoopSize(doorPublicKey)
print("FOUND")
print("subject number: {}, loop size: {}".format(doorSubject,doorLoop))
print("")


print("Encryption key",cardPublicKey,doorLoop)
print(transformSubjectNumber(cardPublicKey,doorLoop))

print("Encryption key",doorPublicKey,cardLoop)
print(transformSubjectNumber(doorPublicKey,cardLoop))
