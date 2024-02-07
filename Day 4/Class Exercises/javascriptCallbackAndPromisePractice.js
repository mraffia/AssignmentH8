// ------- JAVASCRIPT CALLBACK AND PROMISE PRACTICE -------

// Soal 1: Fetch Data dari API (Callback)
function fetchUserData(username, cb) {
  fetch(`https://api.github.com/users/${username}`)
  .then(response => response.json())
  .then(data => cb(data))
  .catch(error => console.log(error));
}

function retrieveName(data) {
  console.log(data.name);
}
//console.log(fetchUserData('mraffia', retrieveName));

// Soal 2: Fetch Data dari API (Promise)
function fetchUserDataPromise(username, cb) {
  fetch(`https://api.github.com/users/${username}`)
  .then(response => response.json())
  .then(data => cb(data))
  .catch(error => console.log(error));
}