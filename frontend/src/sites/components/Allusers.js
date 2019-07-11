import React from 'react';
import classnames from 'classnames';
import Modal from './Modal';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSyncAlt } from '@fortawesome/free-solid-svg-icons'
import { faPlus } from '@fortawesome/free-solid-svg-icons'
import { faSearch } from '@fortawesome/free-solid-svg-icons'

class Allusers extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            data : [],
            copyData: [],
            userToDelete: undefined,
            showInfo: false,
            showErr: false,
            err: "",
            nameFilter: '',
            userName: "",
            userSurname: "",
            userEmail: "",
            userPassword: "",
            userAdmin: false,
            info: ""
        };

        this.deleteUser = this.deleteUser.bind(this)
    }
    componentDidMount() {
        this.loadUsers()
    }

    loadUsers = () =>{
        fetch('http://localhost:8000/user/get', {
            method: "GET",
            headers: new Headers({
                'Authorization' : "Bearer " + localStorage.getItem('token')
            })
        })
        .then(response => {
            response.json().then(json => {
                if(response.status === 401) {
                    localStorage.clear();
                    window.location.href = '/login';
                }
                if(json.status === "succes"){
                    const copyData = json.data.slice();
                    this.setState({
                        data: json.data,
                        copyData: copyData
                    })
                } else {
                    console.log(json.data)
                }
            })
        })
        .catch(function(error) {
            console.log('Request failed', error)
        });
    }

    deleteUser = (id) => {
        fetch('http://localhost:8000/user/delete/' + id, {
            method: "DELETE",
            headers: new Headers({
                'Authorization' : "Bearer " + localStorage.getItem('token')
            })
        })
        .then(response => {
            response.json().then(json => {
                if(response.status === 401) {
                    localStorage.clear();
                    window.location.href = '/login';
                }
                if(json.status === "succes"){
                    this.setState({
                        showInfo: true,
                        showErr:false,
                        info: " User "+json.data.name+" "+json.data.surname +" succesfully deleted."
                    })
                    this.loadUsers()
                } else {
                    this.setState({
                        showInfo: false,
                        showErr:true,
                        err: json.data
                    })
                }
            })
        })
    }

    addUser = () => {
        fetch('http://localhost:8000/user/post/', {
            method: "POST",
            headers: new Headers({
                'Authorization' : "Bearer " + localStorage.getItem('token')
            }),
            body: JSON.stringify({
                email : this.state.userEmail,
                name : this.state.userName,
                surname : this.state.userSurname,
                password : this.state.userPassword,
                isadmin: this.state.userAdmin
            })
        })
        .then(response => {
            response.json().then(json => {
                if(response.status === 401) {
                    localStorage.clear();
                    window.location.href = '/login';
                }
                if(json.status === "succes"){
                    this.setState({
                        showInfo: true,
                        showErr:false,
                        info: "User added"
                    })
                    this.loadUsers()
                } else {
                    this.setState({
                        showInfo: false,
                        showErr:true,
                        err: json.data
                    })
                }
            })
        })
    }

    filterData = (e) => {
       e.preventDefault()
       const data = this.state.data.filter(d => {
           return d.name.toLowerCase().includes(this.state.nameFilter.toLowerCase());
       });

       this.setState({
           copyData: data
       });
    }

    handleUserName = (event) => {
        this.setState({userName: event.target.value});
    }

    handleUserSurname = (event) => {
        this.setState({userSurname: event.target.value});
    }

    handleUserEmail = (event) => {
        this.setState({userEmail: event.target.value});
    }

    handleUserPassword = (event) => {
        this.setState({userPassword: event.target.value});
    }
    handleUserAdmin = (event) => {
        this.setState({userAdmin: event.target.value});
    }

    render() {
        return (
            <div className={classnames('allusers', this.props.styleName)}>
                <div className="alert alert-success text-center mb-2" htmlrole="alert" style={{display: this.state.showInfo ? 'block': 'none'}}>
                    {this.state.info}
                   
                </div>
                <div className="alert alert-danger text-center mb-2" htmlrole="alert" style={{display: this.state.showErr ? 'block': 'none'}}>
                    {this.state.err}
                </div>
                <Modal styleName="fade" title="Are you sure?" button="Delete" id="usermodal" parentFunc={this.deleteUser} item={this.state.userToDelete}>
                    <p>Do you want to delete this user. The change is irreversible.</p>
                </Modal>
                <Modal styleName="fade" title="Add user" button="Add" id="addusermodal" parentFunc={this.addUser} >
                    <form>
                        <label htmlFor="name">Name</label>
                        <input id="name" type="text" placeholder="Enter name" className="form-control" onChange={this.handleUserName}/>
                        <label htmlFor="surname">Surname</label>
                        <input id="surname" type="text" placeholder="Enter Surname" className="form-control" onChange={this.handleUserSurname}/>
                        <label htmlFor="email">Email</label>
                        <input id="email" type="email" placeholder="Enter email" className="form-control" onChange={this.handleUserEmail}/>
                        <label htmlFor="password">Password</label>
                        <input id="password" type="password" placeholder="Enter password" className="form-control" onChange={this.handleUserPassword}/>
                        <div className="form-check">
                            <input className="form-check-input" type="checkbox" value="" id="isadmin" onChange={this.userAdmin}/>
                            <label className="form-check-label" htmlFor="isadmin">
                                Admin
                            </label> 
                        </div>
                    </form>
                </Modal>
                <form className="form-inline d-flex d-lg-none justify-content-center mb-2" onSubmit={this.filterData}>
                    <div className="input-group">
                        <input className="form-control " type="search" placeholder="Search" 
                            aria-label="Search " onChange={(e) => { this.setState({ nameFilter: e.target.value }) }}/>
                        <div className="input-group-append">
                            <button type="submit" className="btn btn-outline-success" ><FontAwesomeIcon icon={faSearch}/></button> 
                        </div>
                        
                    </div>
                    <button className="btn btn-outline-success ml-2"  type="button" onClick={this.loadUsers}><FontAwesomeIcon icon={faSyncAlt} /></button>
                    <button className="btn btn-outline-success ml-2"  type="button" data-toggle="modal" data-target="#addusermodal"><FontAwesomeIcon icon={faPlus} /></button>
                </form>
                <table className="table">
                    <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Name</th>
                            <th scope="col">Surname</th>
                            <th scope="col">Email</th>
                            <th  className="d-flex justify-content-center ">
                                <form className="form-inline d-none d-lg-flex" onSubmit={this.filterData}>
                                    <div className="input-group">
                                        <input className="form-control " type="search" placeholder="Search" aria-label="Search "
                                            onChange={(e) => { this.setState({ nameFilter: e.target.value }) }}/>
                                        <div className="input-group-append">
                                            <button className="btn btn-outline-success" type="submit"><FontAwesomeIcon icon={faSearch} /></button> 
                                        </div>
                                        
                                        
                                    </div>    
                                    <button className="btn btn-outline-success ml-2"  type="button" onClick={this.loadUsers}><FontAwesomeIcon icon={faSyncAlt} /></button>
                                    <button className="btn btn-outline-success ml-2"  type="button" data-toggle="modal" data-target="#addusermodal"><FontAwesomeIcon icon={faPlus} /></button>
                                </form>
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            this.state.copyData.map((user, i) => {
                                return (
                                    <tr key={user.id.toString()}>
                                        <th scope="row">{i + 1}</th>
                                        <td>{ user.name }</td>
                                        <td>{ user.surname }</td>
                                        <td>{ user.email }</td>
                                        <td className="text-center">
                                            <button type="button" onClick={()=> { this.setState({userToDelete: user.id}) }} className="btn btn-danger" data-toggle="modal" data-target="#usermodal">Delete</button>
                                        </td>
                                    </tr>
                                )
                            })
                        }
                    </tbody>
                </table>
            </div>
    );
 }

}


export default Allusers;
