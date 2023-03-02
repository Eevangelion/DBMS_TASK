import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;
const apiPort = process.env.REACT_APP_API_PORT;

export const authService = createApi({
    reducerPath: 'authAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}:${apiPort}`}),
    endpoints: (build) => ({
        getGit: build.query({
            query: (code) => ({url: `/user/oauth/${code}`}),
        }),
        getToken: build.mutation({
            query: () => {
                const refreshToken = localStorage.getItem("refresh_token");
                const accessToken = localStorage.getItem("access_token");
                const userID = localStorage.getItem("userID");
                const userName = localStorage.getItem("userName");
                const userRole = localStorage.getItem("userRole"); 
                return {
                    url: `/user/refresh/`,
                    method: `POST`,
                    headers: {authorization: accessToken},
                    body: {
                        refresh_token: refreshToken,
                        id: userID, 
                        name: userName, 
                        role: userRole
                    }
                }
            }
        }),
        loginUser: build.mutation({
            query: (data) => {
                return {
                    url: `/user/login`,
                    method: 'POST',
                    body: data,
                }
            }
        }),
        registerUser: build.mutation({
            query: (data) => {
                return {
                    url: `/user/register`,
                    method: 'POST',
                    body: data,
                }
            }
        })
    })
})


export const {
    useGetGitQuery,
    useGetTokenMutation,
} = authService;