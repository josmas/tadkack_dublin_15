var tropowebapi = require('tropo-webapi') ,
    express = require('express');

// Create an express application
var app = express();

app.get('/', function(req, res) {
  res.render('index.ejs', {});
});

app.post('/', function(req, res){
  // Create a new instance of the TropoWebAPI object.
  var tropo = new tropowebapi.TropoWebAPI();
  // Use the say method https://www.tropo.com/docs/webapi/say.htm
  tropo.say("Hello World!");

  res.send(tropowebapi.TropoJSON(tropo));
});

app.listen(1337);
console.log('Visit http://localhost:1337/ to accept inbound calls!');
