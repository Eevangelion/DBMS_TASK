const clientId = process.env.REACT_APP_CLIENT_ID;

export const getAuthorizeCodeHref = () => {
    return `https://github.com/login/oauth/authorize?client_id=${clientId}`;
};

export const getCode = () => {
    const urlParams = new URLSearchParams(window.location.search);
    const code = urlParams.get("code");
    return code;
}
