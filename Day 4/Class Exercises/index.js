function findMost(arr) {
  let most;
  let mostCount = 0;

  let mostTemp;
  let mostTempCount = 0;

  for (let i = 0; i < arr.length; i++) {
    mostTemp = arr[i];

    for (let j = 0; j < arr.length; j++) {
      if (arr[j] === mostTemp) {
        mostTempCount++;
      }
    }

    if (mostTempCount > mostCount) {
      most = mostTemp;
      mostCount = mostTempCount;
    }

    mostTempCount = 0;
  }

  console.log(most + ' ' + mostCount);
  return most;
}

// findMost([1,2,2,3,3,3,4,2]);
// findMost([5,5,7,'a','b','a','a','a','c',7]);

let janjian = new Promise(
  (resolve, reject) => {
    resolve('bruh');
    reject('oops');
  }
);

//janjian.then((result) => { console.log(result) }).catch((error) => { console.log(error) })