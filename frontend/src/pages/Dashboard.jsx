import React from 'react';
import { Card, Row, Col, Statistic } from 'antd';

const Dashboard = () => {
  return (
    <div>
      <h1>Dashboard</h1>
      <Row gutter={16}>
        <Col span={8}>
          <Card>
            <Statistic title="Total MDS Entries" value={42} />
          </Card>
        </Col>
        <Col span={8}>
          <Card>
            <Statistic title="Active MDS" value={38} />
          </Card>
        </Col>
        <Col span={8}>
          <Card>
            <Statistic title="Recent Updates" value={12} />
          </Card>
        </Col>
      </Row>
    </div>
  );
};

export default Dashboard;
