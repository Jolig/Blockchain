
var express = require('express');
var router = express.Router();


// Get Homepage
router.get('/', function(req, res){
	res.render('index');
});


router.get('/log', function(req, res){
	res.render('createLogForm');
});


module.exports = router;
