import React from 'react';
import classnames from 'classnames';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSearch } from '@fortawesome/free-solid-svg-icons'

class Search extends React.Component {

  constructor(props) {
    super(props);

  }
  

  render() {
   return (
        <form className={classnames("history", this.props.styleName) } style={this.props.style} onSubmit={this.props.filter}>
            <div className={classnames("input-group", this.props.barClass)}>
                <input className="form-control" type="search" placeholder="Search" 
                aria-label="Search " onChange={this.props.changeHandler}/>
                <div className="input-group-append">
                        <button type="submit" className="btn btn-outline-success" ><FontAwesomeIcon icon={faSearch}/></button> 
                </div>
            </div>
            {this.props.children}
        </form>
   );
 }

}


export default Search;