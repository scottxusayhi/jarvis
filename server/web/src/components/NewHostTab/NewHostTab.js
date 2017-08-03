import React from 'react';
import { TabContent, TabPane, Nav, NavItem, NavLink, Card, Button, CardTitle, CardText, Row, Col } from 'reactstrap';
import { InputGroup, InputGroupAddon, Input } from 'reactstrap';
import classnames from 'classnames';

export default class NewHostTab extends React.Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.state = {
      activeTab: '2'
    };
  }

  toggle(tab) {
    if (this.state.activeTab !== tab) {
      this.setState({
        activeTab: tab
      });
    }
  }
  render() {
    return (
      <div>
        <Nav tabs>
          <NavItem>
            <NavLink
              className={classnames({ active: this.state.activeTab === '1' })}
              onClick={() => { this.toggle('1'); }}
            >
              表格
            </NavLink>
          </NavItem>
          <NavItem>
            <NavLink
              className={classnames({ active: this.state.activeTab === '2' })}
              onClick={() => { this.toggle('2'); }}
            >
              JSON
            </NavLink>
          </NavItem>
        </Nav>
        <TabContent activeTab={this.state.activeTab}>
          <TabPane tabId="1">
            <InputGroup>
              <InputGroupAddon>@</InputGroupAddon>
              <Input placeholder="username" />
            </InputGroup>
            <br />
            <InputGroup>
              <InputGroupAddon>
                <Input addon type="checkbox" aria-label="Checkbox for following text input" />
              </InputGroupAddon>
              <Input placeholder="Check it out" />
            </InputGroup>
            <br />
            <InputGroup>
              <Input placeholder="username" />
              <InputGroupAddon>@example.com</InputGroupAddon>
            </InputGroup>
            <br />
            <InputGroup>
              <InputGroupAddon>$</InputGroupAddon>
              <InputGroupAddon>$</InputGroupAddon>
              <Input placeholder="Dolla dolla billz yo!" />
              <InputGroupAddon>$</InputGroupAddon>
              <InputGroupAddon>$</InputGroupAddon>
            </InputGroup>
            <br />
            <InputGroup>
              <InputGroupAddon>$</InputGroupAddon>
              <Input placeholder="Amount" type="number" step="1" />
              <InputGroupAddon>.00</InputGroupAddon>
            </InputGroup>
          </TabPane>
          <TabPane tabId="2">
            <Input type="textarea" name="text" id="exampleText" rows="20"/>
          </TabPane>
        </TabContent>
      </div>
    );
  }
}