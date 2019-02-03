import React, { Component } from 'react';
import { Row, Col } from 'react-bootstrap';

import StatsCard from '../../components/StatsCard/StatsCard.jsx';
import MoneyLabel from '../../components/MoneyLabel/MoneyLabel.jsx';

export class HeroStats extends Component {
  render() {
    return (
      <Row>
        <Col lg={3} sm={6}>
          <StatsCard
            bigIcon={<i className="pe-7s-wallet text-success" />}
            statsText="Account Balance"
            statsValue={<MoneyLabel amount={this.props.balance} />}
            statsIconText="Remaining cash balance"
          />
        </Col>
        <Col lg={3} sm={6}>
          <StatsCard
            bigIcon={<i className="pe-7s-portfolio text-danger" />}
            statsText="Net Porfolio"
            statsValue={<MoneyLabel amount={this.props.netAsset} />}
            statsIconText="Current total asset amount"
          />
        </Col>
      </Row>
    );
  }
}

export default HeroStats;
