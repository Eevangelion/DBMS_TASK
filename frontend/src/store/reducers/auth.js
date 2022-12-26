import { createSlice } from '@reduxjs/toolkit';
import { getAccessToken, getCode, getJWToken, getUserInfo } from '../actions/authActions';


const initialState = {
    code: null,
    loading: false
};

export const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        setCode: {
            reducer: (state, action) => {
                state.code = action.payload;
            },
            prepare: getCode
        },
        logout: (state) => {
            localStorage.removeItem('userToken');
            state.code = null;
            state.accessToken = undefined;
            state.user = undefined;
        }
    },
    extraReducers: (builder) => {
        builder
            .addCase(getAccessToken.fulfilled, (state, action) => {
                state.code = null;
                state.accessToken = action.payload.token;
                state.loading = false;
                state.error = undefined;
            })
            .addCase(getAccessToken.pending, (state) => {
                state.loading = true;
            })
            .addCase(getAccessToken.rejected, (state, action) => {
                state.loading = false;
                state.error = action.payload;
            })

            .addCase(getJWToken.pending, (state) => {
                state.loading = true;
                state.error = undefined;
            })
            .addCase(getJWToken.fulfilled, (state, action) => {
                localStorage.setItem('userToken', action.payload.AuthToken);
                state.accessToken = undefined;
                state.loading = false;
            })
            .addCase(getJWToken.rejected, (state, action) => {
                state.loading = false;
                state.error = action.payload;
            })
            .addCase(getUserInfo.pending, (state) => {
                state.loading = true;
            })
            .addCase(getUserInfo.fulfilled, (state, action) => {
                state.user = action.payload;
                state.loading = false;
            })
            .addCase(getUserInfo.rejected, (state, action) => {
                state.loading = false;
                state.error = action.payload?.error_message;
            });
    }
});

export const { setCode, logout } = authSlice.actions;
export default authSlice.reducer;