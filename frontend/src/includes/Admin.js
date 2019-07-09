import React from 'react';
import Header from './Header';

class Admin extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
    };
  }

  render() {
   return (
     <div className="conainer">
       
       <Header></Header>
     </div>
   );
 }

}


export default Admin;