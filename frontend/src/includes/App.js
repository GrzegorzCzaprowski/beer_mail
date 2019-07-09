import React from 'react';
import Header from './Header';
import Auth from './Auth';

class App extends React.Component {

 render() {
   return (
     <div className="App">
       <Auth></Auth>
       <Header></Header>
     </div>
   );
 }

}


export default App;
