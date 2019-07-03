import React from 'react';
import logo from '../img/logo.jpg';

class Header extends React.Component {



 render() {
   return (
       <nav class="navbar navbar-expand-lg navbar-light bg-light fixed-top">
           <a class="navbar-brand"><img src={logo} height="50" width="50"/>BeerMail</a>
           <button class="navbar-toggler" data-target="#my-nav" data-toggle="collapse" aria-controls="my-nav" aria-expanded="false" aria-label="Toggle navigation">
               <span class="navbar-toggler-icon"></span>
           </button>
           <div id="my-nav" class="collapse navbar-collapse">
               <ul class="navbar-nav mr-auto">
               <li class="nav-item active">
                   <a class="nav-link" href="#">S<span class="sr-only">(current)</span></a>
               </li>
               <li class="nav-item">
                   <a class="nav-link disabled" href="#" tabindex="-1" aria-disabled="true">Item 2</a>
               </li>
               </ul>
               <span class="navbar-text">
                   <ul class="navbar-nav mr-auto">
                   <li class="nav-item active">
                       <a class="nav-link" href="#">Wyloguj<span class="sr-only">(current)</span></a>
                   </li>
                   </ul>
               </span>
           </div>
       </nav>
   );
 }

}

export default Header;