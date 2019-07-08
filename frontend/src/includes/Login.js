import React from 'react';
import logo from '../img/logo.jpg'

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
      
        e.preventDefault()
        fetch('http://localhost:8000/user/login', {
            method: 'POST',
            credentials: 'include',
            body: JSON.stringify({
                email : login,
                password: password,
            })
        })
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
          
        <div className="login">
            <div class="row justify-content-center mt-3">
                <img src={logo} height="100" width="100"/>
            </div>
            <div class="row justify-content-md-center mt-3">
                <form class="col-sm-12 col-md-4" onSubmit={this.submitLogin}>
                    <div class="form-group">
                        <label for="exampleInputEmail1">Email</label>
                        <input onChange={this.handleChangeLogin} type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email" />
                    </div>
                    <div class="form-group">
                        <label for="exampleInputPassword1">Hasło</label>
                        <input onChange={this.handleChangePassword} type="password" class="form-control" id="exampleInputPassword1" placeholder="Password" />
                    </div>
                    <button type="submit" class="btn btn-primary">Zaloguj</button>
                </form>
            </div>
        </div>
      );
    }
   
   }
   
   
   export default Login;