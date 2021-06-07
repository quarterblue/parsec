import React from 'react';
import './App.css';
import 'antd/dist/antd.css';
import { Layout, Menu, Breadcrumb } from 'antd';
import Block from './components/Block';

const { Header, Content, Footer } = Layout;

const App = () => {
  return (
    <Layout className="layout">
      <Header>
        <div className="logo" />
        <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
          <Menu.Item key="0" style={{ marginRight: '2em', fontSize: '1.2em' }}>
            PARSEC
          </Menu.Item>
          <Menu.Item key="1">SHA256</Menu.Item>
          <Menu.Item key="2">Block</Menu.Item>
          <Menu.Item key="3">Blockchain</Menu.Item>
          <Menu.Item key="4">Distributed Network</Menu.Item>
          <Menu.Item key="5">Transactions</Menu.Item>
          <Menu.Item key="6">Wallet</Menu.Item>
          <Menu.Item key="7">Signed</Menu.Item>
          <Menu.Item key="8">Wallet-Transactions</Menu.Item>
        </Menu>
      </Header>
      <Content style={{ padding: '0 50px' }}>
        <Breadcrumb style={{ margin: '16px 0' }}>
          <Breadcrumb.Item>Home</Breadcrumb.Item>
          <Breadcrumb.Item>List</Breadcrumb.Item>
          <Breadcrumb.Item>App</Breadcrumb.Item>
        </Breadcrumb>
        <div className="site-layout-content">Content</div>
        <Block />
      </Content>
      <Footer style={{ textAlign: 'center' }}>
        Created by quarterblue 2021
      </Footer>
    </Layout>
  );
};

export default App;
