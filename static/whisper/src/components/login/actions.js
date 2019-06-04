export const doLogin = (state, setState) => {
    fetch();
    setState({
        username: state.username === "" && "Test",
        password: state.password === "" && "TestPass"
    });
    console.log(`${state.username}, ${state.password}`);
};
