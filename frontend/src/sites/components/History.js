import React from 'react';
import classnames from 'classnames';
import Search from './Search';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSyncAlt } from '@fortawesome/free-solid-svg-icons'

class History extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            data : [],
            copyData: [],
        }; 
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
    componentDidMount(){
        this.loadEvents()
    }

    loadEvents = () =>{
        fetch('http://localhost:8000/event/get', {
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
                    if(json.data != null){
                        const copyData = json.data.slice();
                        this.setState({
                            data: json.data,
                            copyData: copyData
                        })
                    }
                } else {
                    console.log(json.data)
                }
            })
        })
        .catch(function(error) {
            console.log('Request failed', error)
        });
    }

    handleSearch = (e) => {
        this.setState({ nameFilter: e.target.value }) 
    }

  render() {
   return (
     <div className={classnames("history", this.props.styleName) } style={this.props.style}>
        <div className="row justify-content-center mb-3">
            <Search filter={this.filterData} changeHandler={this.handleSearch} styleName="form-inline justify-content-center">
                <button className="btn btn-outline-success ml-2"  type="button" onClick={this.loadEvents}><FontAwesomeIcon icon={faSyncAlt} /></button>
            </Search>
        </div>
        <div className="row ml-4 mr-4 border-top">
            {
                this.state.copyData.map((event, i) => {
                    return (
                        <div className="col-12">
                            {event.name}
                        </div>
                    )
                })
            }
        </div>

     </div>
   );
 }

}


export default History;