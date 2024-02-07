// ------- JAVASCRIPT LOGIC EXERCISE -------

// Exercise 1
function compareNumbers(firstNumber, secondNumber) {
  if (firstNumber === secondNumber) return -1
  else if (firstNumber > secondNumber) return false
  else return true;
}
// console.log(compareNumbers(5, 8));
// console.log(compareNumbers(5, 3));
// console.log(compareNumbers(4, 4));
// console.log(compareNumbers(3, 3));
// console.log(compareNumbers(17, 2));

// Exercise 2
function reverseString(text) {
  let reversed = '';
  for (let i = text.length; i >= 0; i--) {
    reversed += text.charAt(i);
  }
  return reversed;
}
// console.log(reverseString('Hello World and Coders'));
// console.log(reverseString('John Doe'));
// console.log(reverseString('I am a bookworm'));
// console.log(reverseString('Coding is my hobby'));
// console.log(reverseString('Super'));

// Exercise 3
function urutHuruf(text) {
  let arr = text.split('');

  for (var i = 0; i < arr.length; i++) {
    // Last i elements are already in place  
    for (var j = 0; j < (arr.length - i - 1); j++) {
      // Checking if the item at present iteration is greater than the next iteration
      if (arr[j] > arr[j + 1]) {
        // If the condition is true then swap them
        var temp = arr[j]
        arr[j] = arr[j + 1]
        arr[j + 1] = temp
      }
    }
  }

  return arr.join('');
}
// console.log(urutHuruf('halo'));
// console.log(urutHuruf('qwerty'));
// console.log(urutHuruf('qwertyuiopasdfghjklzxcvbnm'));

// Exercise 4
function isArithmeticProgression(numbers) {
  let difference = null;
  let isArith = true;

  for (let i = 0; i < numbers.length - 1; i++) {
    if (difference === null) {
      difference = numbers[i + 1] - numbers[i];
    } 

    if (difference !== numbers[i + 1] - numbers[i]) {
      isArith = false;
    }
  }

  return isArith;
}
// console.log(isArithmeticProgression([1, 2, 3, 4, 5, 6]));
// console.log(isArithmeticProgression([2, 4, 6, 12, 24]));
// console.log(isArithmeticProgression([2, 4, 6, 8]));
// console.log(isArithmeticProgression([2, 6, 18, 54]));
// console.log(isArithmeticProgression([1, 2, 3, 4, 7, 9]));

// Exercise 5
function threeStepsAB(text) {
  for (let i = 0; i < text.length - 4; i++) {
    if (text.charAt(i).toLowerCase() === "a") {
      if (text.charAt(i + 4).toLowerCase() === "b") {
        return true;
      }
    } else if (text.charAt(i).toLowerCase() === "b") {
      if (text.charAt(i + 4).toLowerCase() === "a") {
        return true;
      }
    }
  }

  return false;
}
// console.log(threeStepsAB('lane borrowed'));
// console.log(threeStepsAB('i am sick'));
// console.log(threeStepsAB('you are boring'));
// console.log(threeStepsAB('barbarian'));
// console.log(threeStepsAB('bacon and meat'));

// Exercise 6
function gcd(firstNumber, secondNumber) {
  let divisor = 1;

  for (let i = 1; i < firstNumber; i++) {
    if (firstNumber % i === 0 && secondNumber % i === 0) {
      divisor = i;
    }
  }

  return divisor;
}
// console.log(gcd(12, 16))
// console.log(gcd(50, 40))
// console.log(gcd(22, 99))
// console.log(gcd(24, 36))
// console.log(gcd(17, 23))

// Exercise 7
function isPrime(number) {
  let prime = true;

  for (let i = 2; i < number; i++) {
    if (number % i === 0) {
      prime = false;
    }
  }

  return prime;
}
// console.log(isPrime(3));
// console.log(isPrime(7));
// console.log(isPrime(6));
// console.log(isPrime(23));
// console.log(isPrime(33));

// Exercise 8
function listPrima(angkaPertama, angkaKedua) {
  let listPrime = [];

  for (let i = angkaPertama; i <= angkaKedua; i++) {
    if (i === 1) {
      continue;
    }

    let prime = true;
    for (let j = 2; j < i; j++) {
      if (i % j === 0) {
        prime = false;
      }
    }

    if (prime) {
      listPrime.push(i);
    }
  }

  return listPrime;
}
// console.log(listPrima(1, 5));
// console.log(listPrima(5, 10));
// console.log(listPrima(10, 20));