import React from 'react';
import logo from '../img/logo.jpg'
import axios from 'axios'

class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            login: '',
            password: ''
        };
    }

    submitLogin = (e) => {
        const login = this.state.login
        const password = this.state.password;
        var http = axios.create({
            baseURL: 'http://localhost:8000',
            body: JSON.stringify({
                email : login,
                password: password,
            }),
            headers: {
                Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaXNhZG1pbiI6dHJ1ZSwiZXhwIjoxNTYyNTg5NTk4fQ.i_I1-kqdbYy90qcBNGePQhhM1029jopLsO4DtWDOIZc" 
            }
            
        })
        e.preventDefault()
        http.post("/user/login")
        .then(response => response.json())
        .then(data => {
            console.log(data)
        })
        .catch(function(error) {
            console.log('Request failed', error)
        });

    }

    handleChangeLogin = (event) => {
        this.login = event.target.value;
        this.setState({login: event.target.value});
    }

    handleChangePassword = (event) => {
        this.setState({password: event.target.value});
    }

    render() {
      return (
         
        <div className="login-background">
            <div className="container">
                <div className="login">
                    <div className="row justify-content-center mt-3">
                        <img src={logo} height="100" width="100"/>
                    </div>
                    <div className="row justify-content-center mt-3 ">
                        <form className="col-sm-12 col-md-5" onSubmit={this.submitLogin}>
                            <div className="form-group">
                                <label htmlFor="exampleInputEmail1">Email</label>
                                <input onChange={this.handleChangeLogin} type="email" className="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email" />
                            </div>
                            <div className="form-group">
                                <label htmlFor="exampleInputPassword1">Hasło</label>
                                <input onChange={this.handleChangePassword} type="password" className="form-control" id="exampleInputPassword1" placeholder="Password" />
                            </div>
                            <div className="text-center">
                                <button type="submit" className="btn btn-primary col-12">Zaloguj</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>        
      );
    }
   
   }
   
   
   export default Login;