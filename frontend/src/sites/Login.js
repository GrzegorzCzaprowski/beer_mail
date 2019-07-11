import React from 'react';
import logo from '../img/logo.jpg'


class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            login: '',
            password: '',
            err: '',
            showErr: false
        };
    }

    submitLogin = (e) => {
        e.preventDefault()
        const login = this.state.login
        const password = this.state.password;
        fetch('http://localhost:8000/user/login', {
            method: "POST",
            body: JSON.stringify({
                email : login,
                password : password
            })
        })
        .then(response => response.json())
        .then(data => {
            if(data.status === "succes"){
                localStorage.setItem('token', data.data)
                window.location.href = '/';
            }else {
                this.setState({
                    err: data.data,
                    showErr: true
                })
            }

        })
        .catch(function(error) {
            console.log('Request failed', error)
        });
    }

    handleChangeLogin = (event) => {
        this.setState({login: event.target.value});
    }

    handleChangePassword = (event) => {
        this.setState({password: event.target.value});
    }

    checkSorage = () => {
        const token = localStorage.getItem('token');
        if (token) {
           window.location.href = '/'; 
        }
    }

    render() {
        this.checkSorage(); 
        return (
            <div className="login-background">
                <div className="login-container">
                    <div className="login">
                        <div className="alert alert-danger text-center " htmlrole="alert" style={{display: this.state.showErr ? 'block': 'none'}}>
                            {this.state.err}
                        </div>
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