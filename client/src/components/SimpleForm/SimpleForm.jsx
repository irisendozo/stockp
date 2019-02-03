import React, { Component } from 'react';
import {
  FormGroup,
  Row,
  Col,
  ControlLabel,
  FormControl,
  HelpBlock,
} from 'react-bootstrap';

import Card from '../Card/Card.jsx';
import Button from '../../components/CustomButton/CustomButton.jsx';

export class SimpleForm extends Component {
  constructor(props, context) {
    super(props, context);

    this.handleChange = this.handleChange.bind(this);

    this.state = {
      value: 0,
    };
  }

  handleChange(e) {
    this.setState({ value: e.target.value });
  }

  render() {
    return (
      <Card
        title={this.props.title}
        content={
          <form
            onSubmit={() => {
              this.props.submitForm(this.state.value);
              this.setState({ value: 0 });
            }}
          >
            <Row>
              <Col md={12}>
                <FormGroup controlId="formControlsNumber">
                  <ControlLabel>Amount</ControlLabel>
                  <FormControl
                    componentClass="input"
                    bsClass="form-control"
                    placeholder="$ USD"
                    type="number"
                    value={this.state.value}
                    onChange={this.handleChange}
                  />
                </FormGroup>
                {this.props.feedbackMessage &&
                  this.props.feedbackMessage.length > 0 && (
                    <HelpBlock>{this.props.feedbackMessage}</HelpBlock>
                  )}
              </Col>
            </Row>
            <Button bsStyle={this.props.status} pullRight fill type="submit">
              Submit
            </Button>
            <div className="clearfix" />
          </form>
        }
      />
    );
  }
}

export default SimpleForm;
