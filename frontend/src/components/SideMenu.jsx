import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { Menu } from 'antd';
import {
  DashboardOutlined,
  CarOutlined,
  FileTextOutlined
} from '@ant-design/icons';

const { SubMenu } = Menu;

const SideMenu = () => {
  const [collapsed, setCollapsed] = useState(false);

  const onCollapse = collapsed => {
    setCollapsed(collapsed);
  };

  return (
    <Menu
      theme="dark"
      defaultSelectedKeys={['dashboard']}
      mode="inline"
      inlineCollapsed={collapsed}
    >
      <Menu.Item key="dashboard" icon={<DashboardOutlined />}>
        <Link to="/dashboard">Dashboard</Link>
      </Menu.Item>

      <SubMenu key="logistics" icon={<CarOutlined />} title="Logistics">
        <Menu.Item key="mds-entry">
          <Link to="/mds-entry">MDS Entry</Link>
        </Menu.Item>
        <Menu.Item key="mds-listing">
          <Link to="/mds-listing">MDS Listing</Link>
        </Menu.Item>
      </SubMenu>
      
      <Menu.Item key="reports" icon={<FileTextOutlined />}>
        <Link to="/reports">Reports</Link>
      </Menu.Item>
    </Menu>
  );
};

export default SideMenu;
