import React from 'react';
import classnames from 'classnames';


class Modal extends React.Component {

  constructor(props) {
    super(props);
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
            {this.props.children}
            
        </div>
        </div>
    </div>
   );
 }

}

class ModalForm extends React.Component {
  constructor(props) {
    super(props);
    this.submit = this.submit.bind(this)
  }

  submit = () => {
    if (this.props.submit) {
      this.props.submit()
    }
  }

  render() {
    return(
      <form onSubmit={this.submit}>
        <div className="modal-body">
            {this.props.children}
        </div>
        <div className="modal-footer">
          <button className="btn btn-primary" type="submit">{this.props.button}</button>
          <button type="button" className="btn btn-secondary" data-dismiss="modal">Close</button>
        </div> 
      </form>
    )
  }

}

class ModalBody extends React.Component {
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
    return(
      <div>
        <div className="modal-body">
              {this.props.children}
        </div>
        <div className="modal-footer">
          <button className="btn btn-primary" data-dismiss="modal" onClick={this.doParentFunc}>{this.props.button}</button>
          <button type="button" className="btn btn-secondary" data-dismiss="modal">Close</button>
        </div> 
      </div>
    )
  }

}


export default Modal;
export {
  ModalBody,
  ModalForm,
}