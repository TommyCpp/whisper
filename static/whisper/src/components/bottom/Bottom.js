import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';

const styles = makeStyles(theme => ({
    footer: {
        backgroundColor: theme.palette.primary.main,
    }
}));


export default function Bottom() {
    const classes = styles();
    return (
        <footer className={classes.footer}>
            <Container maxWidth="sm">
                <Typography variant="body1">My sticky footer can be found here.</Typography>
            </Container>
        </footer>
    )
}
