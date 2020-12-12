import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import ReactAudioPlayer from 'react-audio-player';
import boopSfx from './sounds/boop.mp3';
import Loader from 'react-loader-spinner';
import TestComponent from './TestComponent';

import 'react-loader-spinner/dist/loader/css/react-spinner-loader.css';
import './App.css';

function ApiURl() {
  const DefaultPort = 5000;
  //const DefaultPortStr = ':'+DefaultPort.toString();
  const DefaultHost = 'http://localhost'; // 0.0.0.0'; 
  
  const url = process.env.REACT_APP_API_DOMAIN || DefaultHost; 
  const apiPort = process.env.PORT || DefaultPort;  
  if (typeof(process.env.PORT) === 'undefined')  {  
    const url = DefaultHost;// + DefaultPortStr;  
    console.log(url);
    return url;
  }
  console.log(url + ':' + apiPort);
  return url + ':' + apiPort;
}

const APIURL = ApiURl();
const TestDataUrl = APIURL + '/test'; 
console.log(TestDataUrl);

class App extends React.Component {
  constructor() {
    super();
    this.initialState = {  
      playSound: false,
    };
    this.state = this.initialState;

    this.handleClick = this.handleClick.bind(this);
  }

  handleClick = (play) => {
    console.log(this.state.playSound);
    this.setState({playSound: !this.state.playSound});
    <ReactAudioPlayer loop={this.state.playSound} />
  };

  render() {
    return (
      <div className='App'>
        <button className='btn-primary' style={{'backgroundColor':'#4ef18f'}} type='button' onClick={ this.handleClick }>Play</button>
        <Loader type='Rings' color='#00BFFF' height={100} width={100} timeout={2000} />
        { this.state.playSound && <ReactAudioPlayer src={boopSfx} autoPlay /> }

      <BrowserRouter>
        <Switch>
        <Route path='/test' component={TestComponent} />
        </Switch>
      </BrowserRouter>
      </div>
    );
  }
}

export default App;

