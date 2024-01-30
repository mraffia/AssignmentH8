function sum(a, b) {
  return a + b;
}

function div(a, b) {
  if (b !== 0) {
    return a/b;
  }
  return 0;
}

// console.log("Hasil Penjumlahan " + sum(10, 15));
// console.log("Hasil Bagi Valid " + div(3, 2));
// console.log("Hasil Bagi Invalid " + div(3, 0));

function segitiga(num) {
  for (let i = 1; i < num; i++) {
    let str = "";

    str += " ".repeat(num - 1 - i);
    str += "*".repeat(i);

    console.log(str);
  }
}

function segitigaPagar(num) {
  for (let i = 1; i < num; i++) {
    let str = "";

    str += " ".repeat(num - 1 - i);
    for (let j = 0; j < i; j++) {
      if (str.length % 2 === 0) {
        str += "*";
      } else {
        str += "#";
      }
    }

    console.log(str);
  }
}

function segitigaSamaSisi(num) {
  for (let i = 1; i < num; i++) {
    let str = "";

    str += " ".repeat(num - 1 - i);
    str += "*".repeat(i);
    str += "*".repeat(i-1);
    str += " ".repeat(num - 1 - i);
    
    console.log(str);
  }
}

segitiga(10);
segitigaPagar(10);
segitigaSamaSisi(10);

let umur = [30, 32, 24, 26, 19, 17, 81];

function sortArray(arr) {
  let sorted = [];
  let pickedIndex = [];

  while (sorted.length < arr.length) {
    let min = 99999;
    let minIndex = 0;

    for (let i = 0; i < arr.length; i++) {
      if (arr[i] < min && !pickedIndex.includes(i)) {
        min = arr[i];
        minIndex = i;
      }
    }
    sorted.push(min);
    pickedIndex.push(minIndex);
  }

  return sorted;
}

function printSortedArray(sorted) {
  for (let i = 0; i < sorted.length; i++) {
    console.log(sorted[i]);
  }
}

printSortedArray(sortArray(umur));