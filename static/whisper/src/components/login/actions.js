import * as Config from '../../config';
import axios from 'axios';

export const doLogin = (state) => {
    axios.post(`${Config.HOST}/login`, state)
        .then(res => {
            if(res.status === 200){
                //if success, then redirect to page
            }
        });
};
