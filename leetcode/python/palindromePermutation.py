

testString = 'Tactt Coooa'

def get_counts(str):
  counts = {}
  for char in str.lower().replace(" ", ""):
    if char in counts:
      counts[char] += 1
    else:
      counts[char] = 1
  return counts

def is_palindrome_permutation(str):
  counts = get_counts(str)
  hasOdd = False

  for char in counts:
    if counts[char] % 2 == 1:
      if hasOdd:
        return False
      else:
        hasOdd = True
  
  return True


print(is_palindrome_permutation(testString))