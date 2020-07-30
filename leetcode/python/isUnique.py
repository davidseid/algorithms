

def isUnique(str): 
  checker = 0

  for char in str: 
    code = ord(char)
    if checker & (1 << code):
      return False
    checker |= (1 << code)
  
  return True


print(isUnique('halbejugir'))