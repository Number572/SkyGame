
let enemies, heroes;
let heroPart = document.querySelector(".heroPart");
let enemyPart = document.querySelector(".enemyPart");
let step = 0;
let skillsDiv = document.querySelector(".skills");
let spellChosen = false;
let currentPlayer = localStorage.getItem("cureentUser");
  console.log(currentPlayer);
 

 if (currentPlayer === null) {
    currentPlayer = "Hero1";
  }

let dataWithPlayers = {"Heroes":[currentPlayer, "Hero2", "Hero3"]};

console.log(dataWithPlayers.Heroes);


createFight = () => {
  
  dataWithPlayers = JSON.stringify(dataWithPlayers);
  var settings = {
  "async": true,
  "crossDomain": true,
  "url": "http://10.0.31.45:8080/api/battle/create",
  "method": "POST",
  "headers": {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
  },
  "processData": false,
  "data": dataWithPlayers
}


$.ajax(settings).done(function (response) {
  jsonData = response;
  heroes = jsonData["Heroes"];
  enemies = jsonData["Enemies"];
  console.log(heroes)
  for(let i = 0; i < 3; i++) {
      let hero = heroes[i];
      let enemy = enemies[i];
      HeroAdd(hero);
      EnemyAdd(enemy);
    }
    startFight();
    let checkInterval = setInterval(getUpdatedData, 1500);
});
}

let updatedataWithPlayers = JSON.stringify(dataWithPlayers);
  var settings = {
  "async": true,
  "crossDomain": true,
  "url": "http://10.0.31.45:8080/api/battle/create",
  "method": "POST",
  "headers": {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
  },
  "processData": false,
  "data": dataWithPlayers
}

FightDataUpdate = (attacker, defender) => {
  let heroes = JSON.parse(dataWithPlayers);
  let jsonData = {
    "Heroes": heroes.Heroes,
    "Attacker": attacker,
    "Defender": defender
  };
  console.log(jsonData);

  let dataAsString = JSON.stringify(jsonData);

  var settings = {
  "async": true,
  "crossDomain": true,
  "url": "http://10.0.31.45:8080/api/battle/update",
  "method": "POST",
  "headers": {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
  },
  "processData": false,
  "data": dataAsString
}

  $.ajax(settings).done(function (response) {
  console.log(response);
});
}

getUpdatedData = () => {
  var settings = {
  "async": true,
  "crossDomain": true,
  "url": "http://10.0.31.45:8080/api/battle/data",
  "method": "POST",
  "headers": {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
  },
  "processData": false,
  "data": dataWithPlayers
}


$.ajax(settings).done(function (response) {
    updateFightField(response);
});
}

closeFight = () => {
  var settings = {
  "async": true,
  "url": "http://10.0.31.45:8080/api/battle/close",
  "method": "POST",
  "headers": {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
  },
  "data": dataWithPlayers
}

$.ajax(settings).done(function (response) {
  console.log(response);
});
}

updateFightField = (newData) => {
  console.log(newData);
  let newDataForHeroes = newData["Heroes"];
  let newDataForEnemies = newData["Enemies"];
  console.log(newDataForHeroes);
  console.log(newDataForEnemies);
  for(let i = 0; i < 3; i++) {
    currentHP = enemies[i].querySelector("p:last-child");
    currentHP.innerHTML = "HP" + newDataForEnemies[i].HP;
    if (newDataForEnemies[i].HP <= 0) {
        death(enemies[i]);
      }
    console.log(newDataForEnemies[i].HP)

    currentHP = heroes[i].querySelector("p:last-child");
    currentHP.innerHTML = "HP" + newDataForHeroes[i].HP;

    if (newDataForHeroes[i].HP <= 0) {
        death(heroes[i]);
      }
  }
}

createFight()



  HeroAdd = (hero) => {
    let charClass;
    if(hero.Class === 0) {
      charClass = "warrior"
    } else {
      charClass = "heal"
    }
    let heroDiv = document.createElement("div");
      heroDiv.classList.add("hero");
      heroDiv.classList.add(charClass);
      heroDiv.setAttribute("who", hero.VKID);
      let lvl = document.createElement("p");
      let hp = document.createElement("p");
      lvl.innerHTML = "LVL" + hero.LVL;
      hp.innerHTML = "HP" + hero.HP;
      let sprite = document.createElement("div");
      sprite.classList.add("sprite");
      heroDiv.appendChild(sprite);
      heroDiv.appendChild(lvl);
      heroDiv.appendChild(hp);
      heroPart.appendChild(heroDiv);
  }


  EnemyAdd = (enemy) => {
    let charClass;
    if(enemy.Class === 0) {
      charClass = "warrior"
    } else {
      charClass = "heal"
    }
    let enemyDiv = document.createElement("div");
      enemyDiv.classList.add("enemy");
      enemyDiv.classList.add(charClass);
      enemyDiv.setAttribute("who", enemy.VKID);
      let lvl = document.createElement("p");
      let hp = document.createElement("p");
      lvl.innerHTML = "LVL" + enemy.LVL;
      hp.innerHTML = "HP" + enemy.HP;
      let sprite = document.createElement("div");
      sprite.classList.add("sprite");
      enemyDiv.appendChild(sprite);
      enemyDiv.appendChild(lvl);
      enemyDiv.appendChild(hp);
      enemyPart.appendChild(enemyDiv);
  }

  startFight = () => {
    heroes = document.querySelectorAll(".hero")
    enemies = document.querySelectorAll(".enemy")
    if(step == 0 ) {
      showSkills();
    }
    if(step >= 0 && step < 3) {
      let skills = document.querySelectorAll(".skill");
      for(let i = 0; i < skills.length; i++) {
        skills[i].onclick = () => {
          if (spellChosen == false) {
              skills[i].style.border = "3px solid yellow";
              let random = Math.floor(Math.random()*3);
              attack(heroes[step], 10, enemies[random]);
                console.log(step);
                FightDataUpdate(heroes[step].getAttribute("who"), enemies[random].getAttribute("who"));
                skills[i].style.border = "3px solid rgba(0,0,0,0)";
                nextStep();
                step++;
              spellChosen = true;
              

          }
        }
      }
    }
  }

showSkills = () => {
  for(let i = 0; i < 3; i++) {
    let skill = document.createElement("div");
    skill.classList.add("skill");
    skillsDiv.appendChild(skill)
  }
}

nextStep = () => {
  heroes = document.querySelectorAll(".hero")
  enemies = document.querySelectorAll(".enemy")

  aliveEnemies = checkDead(enemies);

  if (aliveEnemies.length === 0) {
    alert("ПОБЕДА!")
    closeFight();
  }
  let aliveHeroes = checkDead(heroes);
      if (aliveHeroes.length === 0) {
      alert("Проигрыш");
       closeFight();
  }
    
    if (step >= 3)  {
      $(".skills").empty();
      enemiesSteps();
    }

    if(step >= 0 && step < 3) {
      let skills = document.querySelectorAll(".skill");
      for(let i = 0; i < skills.length; i++) {
        skills[i].onclick = () => {
              clearSkillsBorder(skills);
              skills[i].style.border = "3px solid yellow";
              let random = Math.floor(Math.random()* aliveEnemies.length);
              console.log(aliveEnemies);
              if (heroes[step].getAttribute("dead") != "true" && aliveEnemies[random].getAttribute("dead") != "true") {
                attack(heroes[step], 10, aliveEnemies[random]);
                skills[i].style.border = "3px solid rgba(0,0,0,0)";
                console.log(step);
                FightDataUpdate(heroes[step].getAttribute("who"), aliveEnemies[random].getAttribute("who"));
                
                console.log(step);
              }
              step++;
              nextStep();
              
          
        }
      }
    }
    
}


enemiesSteps = () => {

  for(let i = 0; i < enemies.length; i++) {
    setTimeout(() => {
      let aliveHeroes = checkDead(heroes);
      if (aliveHeroes.length === 0) {
        alert("Проигрыш");
        closeFight();
      }
      let random = Math.floor(Math.random()*aliveHeroes.length);
      if(enemies[i].getAttribute("dead") != "true" && aliveHeroes[random].getAttribute("dead") != "true") {
        attack(enemies[i], 10, aliveHeroes[random]);
          FightDataUpdate(enemies[i].getAttribute("who"), aliveHeroes[random].getAttribute("who"));
      }
      step++;
      if (step >=5) {
        setTimeout( () => {
          step = 0;
          nextStep();
          showSkills();
        }, 3000)
        
      }
    }, 2000 * (i+1))
  }
}

clearSkillsBorder = (skills) => {
  for(let i = 0; i < skills.length; i++) {
    skills[i].style.border = "3px solid rgba(0,0,0,0)"
  }
}
  



attack = (hero, damage, enemy) => {
      
      spellChosen = false;
      let currPos = hero.offsetLeft;


      hero.style.left = "0px";
      let newPos = 0;

      if (step == 0) {
        newPos = enemy.offsetLeft - 65;
      }

      if (step == 1) {
        newPos = enemy.offsetLeft - 160;
      }
      if (step == 2) {
        newPos = enemy.offsetLeft - 265;
      }
      if (step == 3) {
        newPos = -(hero.offsetLeft - enemy.offsetLeft) + 30 ;
      }

      if (step == 4) {
        newPos = -(hero.offsetLeft - enemy.offsetLeft) + 50;
      }
      if (step == 5) {
        newPos = -(hero.offsetLeft - enemy.offsetLeft)+53;
      }

      hero.style.left =  newPos + "px";


      let sprite = hero.querySelector(".sprite");
      

      setTimeout(() => {
        sprite.style.backgroundPosition = "-408px -288px";
        sprite.style.height = "65%";
      }, 1000)

      setTimeout(() => {
        sprite.style.backgroundPosition = "-15px 15px";
        sprite.style.height = "70%";
      }, 1100)

      setTimeout(() => {
        hero.style.left = "0px";
      }, 1500)

}

death = (deadPerson) => {
  let sprite = deadPerson.querySelector(".sprite");
  setTimeout(() => {
        sprite.style.backgroundPosition = "-603px -487";
        sprite.style.height = "65%";
      }, 800)

  setTimeout(() => {
        sprite.style.backgroundPosition = "-588px -487px";
        sprite.style.height = "65%";
      }, 1100)
  deadPerson.setAttribute("dead", true);
}


checkDead = (characters) => {
  let returnCharacters = [];
  for(let i = 0; i < characters.length; i++) {
    if (characters[i].getAttribute("dead") != "true") {
      returnCharacters.push(characters[i]);
    }
  }
  return returnCharacters;
}




