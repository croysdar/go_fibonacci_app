import React, {Component} from 'react';
import {Card, InputGroup, FormGroup, Button, Text} from "@blueprintjs/core"

import './App.css';

class App extends Component{
  state = {
    amount: '',
    sequence: "",
  }

  getSequence = async (e) => {
    e.preventDefault();
    
    fetch('api/fibonacci/' + this.state.amount)
    .then(response => {
      return response.text()
    })
    .then(text => {
      console.log(text)
      this.setState({sequence: text})

    })
  }

  render(){
    return (
      <div className="App">
        <div className="App-body">
          <Card className="App-card">
            <h3>Fibonacci Sequence Generator</h3>
            <h4>Enter a number to see a fibonacci sequence</h4>
            <FormGroup >
              <InputGroup 
                id="text-input" 
                value={this.state.amount}
                onChange = {e =>  this.setState({amount: e.target.value}) }
              />
              <Button
                type = "submit"
                id = "submit-button"
                text = "Get sequence"
                onClick = { (e) => this.getSequence(e) }
              />
            </FormGroup>
            <Text> <h5> {this.state.sequence} </h5> </Text>
          </Card>
        </div>
      </div>
    );
  }
}

export default App;
