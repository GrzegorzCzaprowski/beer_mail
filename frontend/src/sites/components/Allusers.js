import React from 'react';
import classnames from 'classnames';

class Allusers extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
        data : []
    };
  }
  componentDidMount() {
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
                this.setState({
                    data: json.data
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

  render() {
    return (
        <div className={classnames('allusers', this.props.styleName)}>
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">Name</th>
                        <th scope="col">Surname</th>
                        <th scope="col">Email</th>
                        <th className="d-flex justify-content-center">
                            <form className="form-inline">
                                <div className="input-group">
                                    <input className="form-control " type="search" placeholder="Search" aria-label="Search "/>
                                    <div className="input-group-append">
                                        <button className="btn btn-outline-success" type="button">Search</button> 
                                    </div>
                                </div>
                            </form>
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {
                        this.state.data.map((user, i) => {
                            return (
                                <tr>
                                    <th scope="row">{i + 1}</th>
                                    <td>{ user.name }</td>
                                    <td>{ user.surname }</td>
                                    <td>{ user.email }</td>
                                    <td className="text-center">
                                        <button type="button" class="btn btn-danger">Delete</button>
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
