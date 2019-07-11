import React from 'react';
import Header from './components/Header';
import History from './components/History';


class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      location: {
        futureEvents: true,
        invitations: false,
        history: false,
      },
      navColor: "#F8F9FA"
    }; 
  }
  handleHistory = () => {
    this.setState({location: {
      futureEvents: false,
      invitations: false,
      history: true,
    }})
  }
  handleFuture = () => {
    this.setState({location: {
      futureEvents: true,
      invitations: false,
      history: false,
    }})
  }
  handleInvitations = () => {
    this.setState({location: {
      futureEvents: false,
      invitations: true,
      history: false,
    }})
  }

  render() {
   return (
     <div className="App">
        <Header></Header>
        <div className="container">
          <div className="row event-nav">
            <div className="col-4 event-nav-com text-center border-right border-left pt-2 pb-2" 
            style={{backgroundColor: this.state.location.futureEvents ? this.state.navColor: "transparent"}}
            onClick={this.handleFuture}>
              Future events
            </div>
            <div className="col-4 event-nav-com text-center border-right pt-2 pb-2"
            style={{backgroundColor: this.state.location.invitations ? this.state.navColor: "transparent"}}
            onClick={this.handleInvitations}>
              Invitations
            </div>
            <div className="col-4 event-nav-com text-center border-right pt-2 pb-2"
            style={{backgroundColor: this.state.location.history ? this.state.navColor: "transparent"}}
            onClick={this.handleHistory}>
              History
            </div>
          </div>
          <div className="row main border-left border-right ">
            <History style={{display: this.state.location.history ? "block": "none"}} styleName="col-12 mt-3"></History>
          </div>
        </div>
     </div>
   );
 }

}


export default App;
