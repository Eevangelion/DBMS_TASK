import { createAsyncThunk } from '@reduxjs/toolkit';

const authEndpoint = process.env.REACT_APP_VKAUTH_URI;
const clientId = process.env.REACT_APP_CLIENT_ID;
const redirectUri = process.env.REACT_APP_REDIRECT_URI;
const littleBoyPort = process.env.REACT_APP_LITTLE_BOY_PORT;
const littleBoyHost = process.env.REACT_APP_LITTLE_BOY_HOST;
const apiHost = process.env.REACT_APP_API_HOST;

export const getAuthorizeCodeHref = (redirect) => {
    const display = 'page';
    const response_type = 'code';
    const state = 4194308;
    // eslint-disable-next-line max-len
    return `${authEndpoint}/authorize?client_id=${clientId}&display=${display}&redirect_uri=${redirectUri}&response_type=${response_type}&v=5.120&state=${state}`;
};



export const getAccessToken = createAsyncThunk(
    'auth/access_token',
    async (_, { getState, rejectWithValue }) => {
        try {
            const code = getState().VkAuth.code;
            const response = await fetch(`http://${littleBoyHost}:${littleBoyPort}/access_token?code=${code}`);
            return (await response.json());
        } catch (error) {
            return rejectWithValue('Не удалось получить Access Token');
        }
    }
);


export const getJWToken = createAsyncThunk(
    'auth/JWToken',
    async function (_, { getState, rejectWithValue }) {
        try {
            const accessToken = getState().VkAuth.accessToken;
            const settings = {
                method: 'POST',
                body: JSON.stringify({
                    AuthToken: accessToken
                })
            };

            const response = await fetch(`http://${apiHost}/auth/vk`, settings);
            return (await response.json());
        } catch (error) {
            return rejectWithValue('Не удалось получить JWT');
        }
    }
);


export const getUserInfo = createAsyncThunk('auth/userInfo', async function (_, { rejectWithValue }) {
    const token = localStorage.getItem('userToken');
    const settings = {
        method: 'GET',
        headers: { authorization: `${token}` }
    };

    const response = await fetch(`http://${apiHost}/users/current`, settings).then((response) => response.json());

    if (response.error) {
        localStorage.removeItem('userToken');
        return rejectWithValue(response);
    }

    return (await response);
});