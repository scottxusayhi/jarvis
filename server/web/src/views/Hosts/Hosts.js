import React, {Component} from 'react'
import RegisteredHosts from './Registered'
import ConnectedHosts from './Connected'
import { Tabs } from 'antd';
const TabPane = Tabs.TabPane;


class Hosts extends Component {

    callback(key) {
      console.log(key);
    }

    render() {
        console.log("rendering");
        return (
            <Tabs onChange={this.callback} defaultActiveKey="2">
                <TabPane tab="已注册" key="1"><RegisteredHosts/></TabPane>
                <TabPane tab="已连接" key="2"><ConnectedHosts/></TabPane>
            </Tabs>
        )
    }
}

export default Hosts
