import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Layout } from 'antd';
import SideMenu from './components/SideMenu';
import Dashboard from './pages/Dashboard';
import MdsEntry from './pages/MdsEntry';
import MdsListing from './pages/MdsListing';
import './App.css';

const { Header, Sider, Content } = Layout;

function App() {
  return (
    <Router>
      <Layout style={{ minHeight: '100vh' }}>
        <Sider width={200} className="site-layout-background">
          <div className="logo" />
          <SideMenu />
        </Sider>
        <Layout>
          <Header className="header">
            <div className="logo" />
            <h1 style={{ color: 'white' }}>Beitler Logistics POC</h1>
          </Header>
          <Content style={{ margin: '24px 16px', padding: 24, background: '#fff' }}>
            <Switch>
              <Route exact path="/" component={Dashboard} />
              <Route path="/dashboard" component={Dashboard} />
              <Route path="/mds-entry" component={MdsEntry} />
              <Route path="/mds-listing" component={MdsListing} />
            </Switch>
          </Content>
        </Layout>
      </Layout>
    </Router>
  );
}

export default App;
