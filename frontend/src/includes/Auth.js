import React from 'react';

class Auth extends React.Component {
    checkAuth() {
        const token = localStorage.getItem('token');
        if (!token) {
           window.location.href = '/login'; 
        }
    }
    
    render() {
        this.checkAuth();
        return null;
    }
}

export default Auth;