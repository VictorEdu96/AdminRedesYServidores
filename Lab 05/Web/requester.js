//Function getting hour from the devide is executing from.
function getDeviceHour() {
  var fecha = new Date();
  let date;
  date = fecha.getHours() + ":" + fecha.getMinutes() + ":" + fecha.getSeconds();
  document.getElementById('deviceHour').innerHTML = date;
}

//Getting the hour via a request made to the server.
//Using fetch to get a response and convert it to JSON
function getServerHour() {
  fetch('http://localhost:8080/srvhour')
    .then(function(response) {
      return response.json();
    })
    .then(function(myJson) {
      var fecha = desJSON(myJson); //Using desJSON function to serve the date.
      document.getElementById('serverHour').innerHTML = fecha;
    });
}

//Take a JSON object and convert it to string format
function desJSON(jsonObj) {
  let date = "";
  let hour = jsonObj.Hour;
  let minutes = jsonObj.Minutes;
  let seconds = jsonObj.Seconds;
  date = String(hour) + ":" + String(minutes) + ":" + String(seconds);
  return date;
}

setInterval(getDeviceHour, 1000)
setInterval(getServerHour, 1000)
