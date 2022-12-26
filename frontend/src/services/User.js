import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;


export const userService = createApi({
    reducerPath: 'userAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}/user`}),
    endpoints: (build) => ({
        getUserByName: build.query({
            query: (name) => ({url: `/${name}`}),
        }),
        getJokesByAuthorName: build.query({
            query: (name, ...params) => {
                const sortArg = params.sortBy ? (params.sortBy === "top" ? (params.t ? params.t : 'day') : params.sortBy) : 'new';
                const pageArg = params.page;
                const args = params.page ? {sortArg, pageArg} : {sortArg};
                return {
                    url: `/${name}/`,
                    params: args
                }
            },
        }),
    })
})

export const {
    useGetUserByNameQuery,
    useGetJokesByAuthorNameQuery
} = userService;