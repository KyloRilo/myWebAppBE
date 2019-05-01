var express = require('express');
var app = express();
const bcrypt = require('bcrypt');
const bodyParser = require('body-parser');
const saltRounds = 10;

app.use(bodyParser.json());

require('dotenv').config()

console.log("Backend Initializing");
const knex = require('knex')({
    client: 'mssql',
    connection: {
        host : process.env.DB_HOST,
        port : Number(process.env.DB_PORT),
        user : process.env.DB_USER,
        password : process.env.DB_PASS,
        database : 'SE41781_KRiley'
    }
})

app.post('/createuser', async (req, res) => {
    const hashedPass = await bcrypt.hash(req.body.Password, saltRounds);
    const user = {
        LoginName: req.body.LoginName,
        PasswordHash: hashedPass,
        FirstName: req.body.FirstName,
        LastName: req.body.LastName,
        Email: req.body.Email
    }
    const [createdUser] = await knex('usersTable').returning(['LoginName', 'FirstName', 'LastName', 'Email']).insert({...user});
    res.json(createdUser);
});
app.post('/login', async (req, res) =>{
    console.log(req.body);
    const login = {
        LoginName: req.body.LoginName
    }
    const [logReturn] = await knex('usersTable').select('*').where({...login});
    if (logReturn){
        if(await bcrypt.compare(req.body.Password, logReturn.PasswordHash)){
            res.json({
                LoggedIn: true,
            });
        }else{
            res.json({Error: "Incorrect Password"});
        }
    }else{
        res.json({Error: "User not found"});
    }
});
app.listen(3001, () => {
    console.log('Server running on http://localhost:3001')
});