import React, { Component } from 'react';

class RightView extends Component {
  render() {
    return (
        <div style={{textAlign: 'right'}}>
            <div style={{display: "inline-block"}}>
                {this.props.children}
            </div>
        </div>
    )
  }
}

export default RightView;
