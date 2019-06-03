import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import CameraIcon from '@material-ui/icons/PhotoCamera';
import CssBaseline from '@material-ui/core/CssBaseline';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import {makeStyles} from '@material-ui/core/styles';
import Login from "./components/login/Login";

const useStyles = makeStyles(theme => ({
    icon: {
        marginRight: theme.spacing(2),
    },
    heroContent: {
        backgroundColor: theme.palette.background.paper,
        padding: theme.spacing(8, 0, 6),
    },
    heroButtons: {
        marginTop: theme.spacing(4),
    },
    cardGrid: {
        paddingTop: theme.spacing(8),
        paddingBottom: theme.spacing(8),
    },
    card: {
        height: '100%',
        display: 'flex',
        flexDirection: 'column',
    },
    cardMedia: {
        paddingTop: '56.25%', // 16:9
    },
    cardContent: {
        flexGrow: 1,
    },
    footer: {
        backgroundColor: theme.palette.background.paper,
        padding: theme.spacing(1),
    },
}));


export default function App() {
    const classes = useStyles();

    return (
        <React.Fragment>
            <CssBaseline/>
            <Box display="flex" flexDirection="column" justifyContent={"space-between"} height={'100vh'}>
                <Box>
                    <AppBar position="relative">
                        <Toolbar>
                            <CameraIcon className={classes.icon}/>
                            <Typography variant="h6" color="inherit" noWrap>
                                Whisper
                            </Typography>
                        </Toolbar>
                    </AppBar>
                </Box>
                <Box>
                    <Login/>
                </Box>
                {/* Footer */}
                <Box>
                    <footer className={classes.footer}>
                        <Typography variant="h6" align="center" gutterBottom>
                            Footer
                        </Typography>
                        <Typography variant="subtitle1" align="center" color="textSecondary" component="p">
                            Something here to give the footer a purpose!
                        </Typography>
                    </footer>
                </Box>
                {/* End footer */}
            </Box>
        </React.Fragment>
    );
}
