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
    })
})


export const {
    useGetGitQuery
} = authService;