import { createAsyncThunk } from '@reduxjs/toolkit';

const authEndpoint = process.env.REACT_APP_VKAUTH_URI;
const clientId = process.env.REACT_APP_CLIENT_ID;
const redirectUri = process.env.REACT_APP_REDIRECT_URI;
const apiHost = process.env.REACT_APP_API_HOST;

export const getAuthorizeCodeHref = () => {
    return `https://github.com/login/oauth/authorize?client_id=${clientId}`;
};

export const getCode = () => {
    const urlParams = new URLSearchParams(window.location.search);
    const code = urlParams.get("code");
    return code;
}
