function getString(){
  console.log("Get String from Backend");
var request = new XMLHttpRequest();
console.log("url app: " + location.href )
request.open('GET', '/data', true);
request.onload = function () {

  // Begin accessing JSON data here
  var data = JSON.parse(this.response);
  console.log(data)
  if (request.status >= 200 && request.status < 400) {
    var h2 = document.getElementById('dataStr');
      h2.innerHTML = data;
  } else {
    const errorMessage = document.createElement('marquee');
    errorMessage.textContent = `Gah, it's not working!`;
    app.appendChild(errorMessage);
  }
}

request.send();
}