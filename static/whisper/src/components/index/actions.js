import Cookies from 'universal-cookie';

const cookies = new Cookies();

function isLogin() {
    return cookies.get('cookies') === null
}
