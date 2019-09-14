window.addEventListener("load", init);
// Globals

// Available Levels
const levels = {
  easy: 10,
  medium: 3,
  hard: 1
};

// To change level
const currentLevel = levels.easy;
document.querySelector("#seconds").innerHTML = currentLevel;

let time = currentLevel;
let score = 0;
let isPlaying;
let randIndex = 0;

// DOM Elements
const wordInput = document.querySelector("#word-input");
const currentWord = document.querySelector("#current-word");
const scoreDisplay = document.querySelector("#score");
const timeDisplay = document.querySelector("#time");
const message = document.querySelector("#message");
const seconds = document.querySelector("#seconds");
const speakerBtn = document.querySelector("#speaker");
let currentUser = localStorage.getItem("cureentUser");
if (currentUser === null) {
  currentUser = "Hero1";
}

sendAnswer = (user, task, answer) => {
  let answerData = {
    VKID: user,
    Task: task,
    Check: answer
  };

  let answerString = JSON.stringify(answerData);
  console.log(answerString);
  var settings = {
    async: true,
    crossDomain: true,
    url: "http://10.0.31.45:8080/api/task/check",
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    data: answerString
  };

  $.ajax(settings).done(function(response) {
    console.log(response);
  });
};

// Init the SpeechSynthesis API
const synth = window.speechSynthesis;
let voices = [];

const words = [
  { english: "hat", translation: ["шляпа", "шапка"] },
  { english: "river", translation: ["река", "поток"] },
  {
    english: "lucky",
    translation: ["везучий", "удачливый", "удачный"]
  },
  {
    english: "communication",
    translation: ["коммуникация", "общение", "сообщение"]
  },
  {
    english: "reward",
    translation: ["награда", "вознаграждение"]
  },
  {
    english: "development",
    translation: ["развитие", "разработка", "создание"]
  },
  { english: "application", translation: ["приложение"] }
];

// Initialize Game
function init() {
  // Show number of seconds in UI
  seconds.innerHTML = currentLevel;
  // Load word from array
  showWord(words);
  // Start matching on word input
  wordInput.addEventListener("input", startMatch);
  // Call countdown every second
  setInterval(countdown, 1000);
  // Check game status
  setInterval(checkStatus, 50);
}

// Start match
function startMatch() {
  if (matchWords()) {
    isPlaying = true;
    time = currentLevel + 1;
    showWord(words);
    wordInput.value = "";
    score++;
  }

  // If score is -1, display 0
  if (score === -1) {
    scoreDisplay.innerHTML = 0;
  } else {
    scoreDisplay.innerHTML = score;
  }
}

// Match currentWord to wordInput
function matchWords() {
  if (words[randIndex].translation.includes(wordInput.value)) {
    let questionWord = document.querySelector("#current-word").innerHTML;
    message.innerHTML = "Correct!!!";
    let question = sendAnswer(currentUser, questionWord, wordInput.value);
    return true;
  } else {
    message.innerHTML = "";
    return false;
  }
}

// Pick & show random word
function showWord(words) {
  // Generate random array index
  randIndex = Math.floor(Math.random() * words.length);
  // Output random word
  currentWord.innerHTML = words[randIndex].english;
}

// Countdown timer
function countdown() {
  // Make sure time is not run out
  if (time > 0) {
    // Decrement
    time--;
  } else if (time === 0) {
    // Game is over
    isPlaying = false;
  }
  // Show time
  timeDisplay.innerHTML = time;
}

// Check game status
function checkStatus() {
  if (!isPlaying && time === 0) {
    message.innerHTML = "Game Over!!!";
    score = -1;
  }
}

getVoices = () => {
  voices = synth.getVoices();
};

// Fill the voice array
getVoices();
if (synth.onvoiceschanged !== undefined) {
  synth.onvoiceschanged = getVoices;
}

const speak = () => {
  // Check if Already speaking
  if (synth.speaking) {
    console.error("Уже говорит...");
    return;
  }

  // Sending a Speech Request to the API
  const speakRequest = new SpeechSynthesisUtterance(currentWord.innerHTML);

  // Run when the speacking will be done:
  speakRequest.onend = e => {
    console.log("Done Speaking!");
  };

  // Error
  speakRequest.onerror = e => {
    console.log("Speacking Error!");
  };

  // Choose the Accent
  speakRequest.voice = voices[6];

  speakRequest.rate = 1;
  speakRequest.pitch = 1;

  // Speak
  synth.speak(speakRequest);
};

// OnClick speakerBtn
speakerBtn.addEventListener("click", e => speak());
