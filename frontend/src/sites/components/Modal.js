import React from 'react';
import classnames from 'classnames';


class Modal extends React.Component {

  constructor(props) {
    super(props);
    this.doParentFunc = this.doParentFunc.bind(this)
  }

  doParentFunc(){
    if(this.props.parentFunc){
        this.props.parentFunc(this.props.item)
    } 
  }

  render() {
   return (
    <div className={classnames("modal", this.props.styleName)} id={this.props.id} tabIndex="-1" role="dialog">
        <div className="modal-dialog" role="document">
        <div className="modal-content">
            <div className="modal-header">
            <h5 className="modal-title">{this.props.title}</h5>
            <button type="button" className="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
            </div>
            <div className="modal-body">
                {this.props.children}
            </div>
            <div className="modal-footer">
            <button type="button" className="btn btn-primary" data-dismiss="modal" onClick={this.doParentFunc}>{this.props.button}</button>
            <button type="button" className="btn btn-secondary" data-dismiss="modal">Close</button>
            </div>
        </div>
        </div>
    </div>
   );
 }

}


export default Modal;