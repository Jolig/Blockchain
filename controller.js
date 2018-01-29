
var http = require('http');
var express = require('express');
var express_handlebars = require('express-handlebars');
var routes = require('./routes/welcome');
var bodyParser = require('body-parser');



// Init App
var app = express();



//Setting Port
app.set('port', 443);



//View Engine
app.set('views', __dirname + '/views');
var hndlbars = express_handlebars.create({defaultLayout:'colors'});
app.engine('handlebars', hndlbars.engine);
app.set('view engine', 'handlebars');



// Set Static Folder
app.use(express.static(__dirname + '/public'));



// Application Home page has to be fetched from routes = ./routes/welcome
app.use('/', routes);



//Creating Server
http.createServer(app).listen(app.get('port'), function(){
  console.log('Express server listening on port ' + app.get('port'));
});



// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: false }))
// parse application/json
app.use(bodyParser.json());



//Control comes here after submit button on CreateLog Form is clicked
app.post('/submitLog', function(req, res) {

	/*var body = req.body;
	console.log(body);*/
	var username = req.body.username;
	var role 	 = req.body.role;
	var action 	 = req.body.action;
    console.log("Entered Details - \n Username : "+username+"   Role: "+role+"   Action: "+action)
    res.render('logSubmit_Success');

});




/*const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer((req, res) => {
	res.statusCode = 200;
	res.setHeader('Content-type', 'text/plain');
	res.end("Hello World!!");
});

server.listen(port, hostname, () => {
	console.log("Server started on port" + port)
});*/

