import React from 'react';
import ReactAudioPlayer from 'react-audio-player';
import boopSfx from './sounds/boop.mp3';
import Loader from 'react-loader-spinner';
import 'react-loader-spinner/dist/loader/css/react-spinner-loader.css';
import './App.css';

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
      </div>
    );
  }
}

export default App;
