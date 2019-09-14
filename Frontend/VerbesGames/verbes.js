const verbes = [{ Task: [] }];
let phrases = [];
const correction = document.querySelector("#correction");
let currentUser = localStorage.getItem("cureentUser");
if (currentUser === null) {
  currentUser = "Hero1";
}

sendAnswer = (user, answer) => {
  let dataAnswer = {
    "VKID": user,
    "Check": answer
  };

  let answerAsString = JSON.stringify(dataAnswer);
  console.log(answerAsString);

  var settings = {
    async: true,
    crossDomain: true,
    url: "http://10.0.31.45:8080/api/input/check",
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    data: answerAsString
  };

  $.ajax(settings).done(function(response) {
    console.log(response);
  });
};

function getVerbes() {
  for (let i = 1; i <= 10; i++) {
    let inputValue = document.querySelector(`#verbe${i}`).value;
    verbes[0].Task.push(inputValue);
  }
}

correction.addEventListener("click", event => {
  getVerbes();
  sendAnswer(currentUser, verbes[0].Task);
  verbes[0].Task = [];
  console.log(verbes);
});

phrases = [
  "1. I ___ (go) to Italy three years ago.",
  "2. That is the best drink I have ever ___  (drink).",
  "I ___ (think) of the best idea.",
  "4. My grandmother ___ (sell) chocolate when she was young.",
  "5. In Egypt, I ___ (grow) so fast that we had to trim it.",
  "6. My tree  ___ (grow) so fast that we had to trim it.",
  "7. I ___ (buy) lots of DVDs last weekend.",
  "8. A man ___  (shoot) at the man, but he missed him.",
  "9. My brother ___ (fall) down the stairs and cracked his head open.",
  "10. I ___ (wear) my best clothes yesterday."
];
