import React from 'react';
import logo from './logo.svg';
import './App.css';
import Header from './components/header/Header'
import {createMuiTheme, withStyles} from '@material-ui/core/styles';
import {ThemeProvider} from '@material-ui/styles';
import Bottom from "./components/bottom/Bottom";
import blue from '@material-ui/core/colors/blue';

const theme = createMuiTheme({
    palette: {
        primary: blue
    }
});


function App() {
    return (
        <div className="App">
            <ThemeProvider theme={theme}>
                <Header/>
                <Bottom/>
            </ThemeProvider>

        </div>
    );
}

export default App;
