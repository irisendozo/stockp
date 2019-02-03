import React, { Component } from 'react';
import {
  FormGroup,
  Row,
  Col,
  ControlLabel,
  FormControl,
} from 'react-bootstrap';

import Card from '../Card/Card.jsx';
import Button from '../CustomButton/CustomButton.jsx';
import MoneyLabel from '../MoneyLabel/MoneyLabel.jsx';

export class PlaceOrderForm extends Component {
  constructor(props, context) {
    super(props, context);

    this.handleChangeFilter = this.handleChangeFilter.bind(this);
    this.handleChangeQuantity = this.handleChangeQuantity.bind(this);
    this.handleChangeMethod = this.handleChangeMethod.bind(this);

    this.state = {
      stockFilter: '',
      quantity: 0,
      method: 'Buy',
    };
  }

  handleChangeFilter(e) {
    this.setState({ stockFilter: e.target.value });
  }

  handleChangeQuantity(e) {
    this.setState({ quantity: e.target.value });
    this.setState({
      estimatedValue: e.target.value * this.props.searchedStock.Price,
    });
  }

  handleChangeMethod(e) {
    this.setState({ method: e.target.value });
  }

  render() {
    return (
      <Card
        title="Place Orders"
        content={
          <form>
            <Row>
              <Col md={12}>
                <FormGroup controlId="formControlsTextarea">
                  <ControlLabel>Search Stock</ControlLabel>
                  <FormControl
                    componentClass="input"
                    bsClass="form-control"
                    value={this.state.stockFilter}
                    onChange={this.handleChangeFilter}
                    placeholder="Stock symbol or code"
                  />
                </FormGroup>
                <FormGroup>
                  <ControlLabel>Current Price Per Unit</ControlLabel>
                  <FormControl.Static>
                    <MoneyLabel amount={this.props.searchedStock.Price} />
                  </FormControl.Static>
                </FormGroup>
                <Button
                  bsStyle="info"
                  pullRight
                  fill
                  onClick={() => {
                    this.props.searchFilter(this.state.stockFilter);
                  }}
                >
                  Search
                </Button>
              </Col>
            </Row>
            <Row>
              <Col md={3}>
                <FormGroup controlId="formControlsTextarea">
                  <ControlLabel>Quantity</ControlLabel>
                  <FormControl
                    componentClass="input"
                    bsClass="form-control"
                    placeholder="# of stocks"
                    value={this.state.quantity}
                    disabled={!this.props.searchedStock.Price}
                    onChange={this.handleChangeQuantity}
                  />
                </FormGroup>
              </Col>
              <Col md={9}>
                <FormGroup controlId="formControlsSelect">
                  <ControlLabel>Method</ControlLabel>
                  <FormControl
                    componentClass="select"
                    placeholder="Method"
                    value={this.state.method}
                    disabled={!this.props.searchedStock.Price}
                    onChange={this.handleChangeMethod}
                  >
                    <option value="buy">Buy</option>
                    <option value="sell">Sell</option>
                  </FormControl>
                </FormGroup>
              </Col>
            </Row>
            <Row>
              <Col md={12}>
                <FormGroup>
                  <ControlLabel>Estimated Purchase Value</ControlLabel>
                  <FormControl.Static>
                    <MoneyLabel amount={this.state.estimatedValue} />
                  </FormControl.Static>
                </FormGroup>
              </Col>
            </Row>
            <Button
              bsStyle="info"
              pullRight
              fill
              onClick={() => {
                this.state.method === 'Buy'
                  ? this.props.buy(
                      this.props.searchedStock.Symbol,
                      this.state.quantity
                    )
                  : this.props.sell(
                      this.props.searchedStock.Symbol,
                      this.state.quantity
                    );
              }}
            >
              Submit
            </Button>
            <div className="clearfix" />
          </form>
        }
      />
    );
  }
}

export default PlaceOrderForm;
