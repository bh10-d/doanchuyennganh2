const express = require('express');
const cors = require('cors');
const mysql = require('mysql2');

const PORT = 9090;
// const HOST = 'localhost';
const HOST = '172.26.0.4';
// const HOST_MYSQL = 'localhost';
const HOST_MYSQL = '172.26.0.2';

var connection = mysql.createConnection({
    host     : HOST_MYSQL,
    user     : 'root',
    password : 'root',
    database : 'dacn2'
});

connection.connect(function(err) {
    if (err) throw err;
    console.log("Connected!!!")
});

const app = express();
app.use(cors());
app.get('/api/data',(req, res) => {
    res.send({
        text: "This api running on nodejs server"
    })
})

app.get('/api/test', (req, res) => {
    var sql = "SELECT * FROM todo_items";
    connection.query(sql, function(err, results) {
      if (err) throw err;
      res.send(results);
    });
})

app.delete('/api/test/:id', (req, res) => {
    var sql = `DELETE FROM todo_items WHERE id=${req.params.id}`
    connection.query(sql, function(err, results) {
        if (err) throw err;
        console.log({status: "success"});
        res.send({status: "success"});
      });
    // res.send({status: "success"});
    console.log(req.params.id)
    
})

app.listen(PORT, HOST);


console.log(`Server is running on: http://${HOST}:${PORT}`);