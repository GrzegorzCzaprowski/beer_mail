import React from 'react';
import Header from './components/Header';
import Allusers from './components/Allusers';

class Admin extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
    };
  }

  render() {
   return (
      <div className= "main">
        <Header></Header>
        <div className="container">
          <div className="row">
            <Allusers styleName="col-12"></Allusers>
          </div>
        </div>
     </div>
   );
 }

}


export default Admin;