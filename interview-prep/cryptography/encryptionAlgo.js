

/*
i: text msg
o: text msg (encrypted)

c: output should look nothing like input
  -  should be very hard to figure out the input

ideas:
  - rotate letters
  - get char code, if index is even half, if odd double
  - sum each number
  - move into range
  - extensions (shift up to other languages, or perhaps use pattern to choose diff languages)

  - david
  - 01234

  - 100, 97, 118, 105, 100
  -->
  - d   a   v    i    d
  - 50, 194, 59, 210, 50
  - 50, 244, 303, 513, 563
  - 50, 58, 117, 48, 98
    2   :     u   0    b

    david -> 2:u0b -> 


    50        58       117       48        98
    50        
    100       
  range 33 - 126 (93)
*/

let string = 'david';
let i = 244; 

while (i > 126) {
  i -= 93;
}

console.log(i)



