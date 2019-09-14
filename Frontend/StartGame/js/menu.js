var links = document.querySelectorAll('.link')
console.log(links)
for(var i = 0; i < links.length; i++){
   var link = links[i]
   addClass(link, i) 
}

function addClass(link, o){
   console.log(link, o)
   setTimeout(function(){
      link.classList.add('demo')
      removeClass(link)
   }, o*750)
}

function removeClass(link){
   setTimeout(function(){
      link.classList.remove('demo')
   }, 750)
}