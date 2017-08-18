import React, { Component } from 'react';
import { Button, Popover, PopoverTitle, PopoverContent } from 'reactstrap';


class PopoverEditor extends Component {
  render() {
    return (
      <div>
        <Popover {...this.props}>
            {this.props.children}
        </Popover>
      </div>
    )
  }
}

export default PopoverEditor;
