import React from 'react';
import classnames from 'classnames';
import Search from './Search';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSyncAlt, faQuestion } from '@fortawesome/free-solid-svg-icons'
import Moment from 'react-moment';

class History extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            data : [],
            copyData: [],
            copyGuests: []
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
            <Search filter={this.filterData} changeHandler={this.handleSearch} styleName="form-inline justify-content-center" barClass="col-10"> 
                <button className="btn btn-outline-success "  type="button" onClick={this.loadEvents}><FontAwesomeIcon icon={faSyncAlt} /></button>
            </Search>
        </div>
        <div className="row ml-4 mr-4 border-top ">
            <div className="col-12 ">
                <div className="row border-bottom">
                    <div className="col-3 text-center pt-3 pb-3"><h4 className="border-left border-right vertical-align-center m-0">Name</h4></div>
                    <div className="col-3 text-center pt-3 pb-3"><h4 className="border-right vertical-align-center m-0">Host</h4></div>
                    <div className="col-3 text-center pt-3 pb-3"><h4 className="border-right vertical-align-center m-0">Place</h4></div>
                    <div className="col-3 text-center pt-3 pb-3"><h4 className="border-right vertical-align-center m-0">Date</h4></div>
                </div>
            </div>
            {
                this.state.copyData.map((event, i) => {
                    
                    return (
                        <div className="col-12" key={i}>
                            <div className="row hoverable" data-toggle="collapse" data-target={"#guests" + i}>
                                <div className="col-3 text-center pt-1 pb-1">
                                    {event.name}
                                </div>
                                <div className="col-3 text-center pt-1 pb-1">
                                    {event.idcreator}
                                </div>
                                <div className="col-3 text-center pt-1 pb-1">
                                    {event.place}
                                </div>
                                <div className="col-3 text-center pt-1 pb-1">
                                    <div className="mr-3 d-inline-block">
                                        <Moment format="H:m">
                                        {event.date}
                                        </Moment>
                                    </div>
                                    <div className="d-inline-block">
                                        <Moment format="DD/MM/Y">
                                        {event.date}
                                        </Moment>
                                    </div>
                                </div>
                            </div>
                            <div className="row collapse justify-content-center" id={"guests" + i}>
                                <div className="col-12 col-lg-6 mt-1 mb-3 " >
                                    <div className="row" style={{backgroundColor: "rgb(248, 249, 250)"}}>
                                        <div className="col-6 text-center pt-1 pb-1">Name</div>
                                        <div className="col-6 text-center pt-1 pb-1">Surname</div>
                                    </div>   
                                    {event.guests ? event.guests.map((guest,j) => {
                                        return (
                                            <div key={j}>
                                                {event.guests.confirm ? 
                                                    <div className="row">
                                                        <div className="col-6 text-center pt-1 pb-1">{guest.name}</div>
                                                        <div className="col-6 text-center pt-1 pb-1">{guest.surname}</div>
                                                    </div>
                                                : <div className="text-center">No guests</div>
                                                }
                                                
                                            </div>
                                        )
                                    }): <div>No guests</div>
                                    }
                                </div>
                            </div>
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