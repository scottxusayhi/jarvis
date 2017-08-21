import React, { Component } from 'react'
import { connect } from 'react-redux'
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import { Popover, PopoverTitle, PopoverContent } from 'reactstrap';
import classnames from 'classnames';
import HotTable from 'react-handsontable';
import Collapsible from 'react-collapsible';
import PopoverEditor from '../../../../components/PopoverEditor/PopoverEditor'

import {
    updateRegHost
} from '../../../../states/actions'

// subscribe
const mapStateToProps = state => {
    return {
        data: state.hostDetail.data
    }
}

// dispatch actions
const mapDispatchToProps = dispatch => {
    return {
        updateRegHost: (id, data) => {
            dispatch(updateRegHost(id, data))
        }
    }
}

let lastId = 0;
function newId (prefix='id') {
    lastId++;
    return `${prefix}${lastId}`;
}

class EditCell extends Component {

  constructor (props) {
      super(props);
      this.toggle = this.toggle.bind(this);
      this.state = {
        popoverOpen: false
    };
  }

  toggle() {
    this.setState({
      popoverOpen: !this.state.popoverOpen
    });
  }

  _onKeyPress(e) {
      if (e.key === 'Enter') {
          console.log('validate and save')
          this.toggle()
          this.props.onEnter()
      }
  }

  getInput() {
      console.log(this.me)
      return this.me.value
  }

  componentWillMount() {
      this.id = newId()
  }

  render() {
    return (
        <Row>
            <Col>{this.props.children}</Col>
            <Col>
                <button type="button" className="button btn-link" id={this.id} onClick={this.toggle}>
                    <i className="fa fa-pencil"/>
                </button>
                <Popover placement="bottom" isOpen={this.state.popoverOpen} target={this.id} toggle={this.toggle}>
                    <PopoverContent>
                        <input
                            className="form-control"
                            type="text"
                            placeholder="input..."
                            defaultValue={this.props.children}
                            ref={(me)=> {this.me = me}}
                            key={this.id}
                            id="example-text-input"
                            onChange={()=>{console.log(this.me.value)}}
                            // autoFocus={true}
                            onFocus={(event)=>{event.target.select()}}
                            onKeyPress={(e)=>{this._onKeyPress(e)}}
                        />
                    </PopoverContent>
                </Popover>
            </Col>
        </Row>
    )
  }



}

export default connect(
    mapStateToProps,
    mapDispatchToProps,
    null,
    {withRef: true}
)(EditCell)
