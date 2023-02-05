import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;
const apiPort = process.env.REACT_APP_API_PORT;

export const userService = createApi({
    reducerPath: 'userAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}:${apiPort}/user`}),
    endpoints: (build) => ({
        getUserByID: build.query({
            query: (id) => ({url: `/${id}/`})
        }),
        getUserByName: build.query({
            query: (name) => ({url: `/${name}/data/`}),
        }),
        getJokesByAuthorName: build.query({
            query: (name, ...params) => {
                const sortArg = params.sortBy ? (params.sortBy === "top" ? (params.t ? params.t : 'day') : params.sortBy) : null;
                const pageArg = params.page;
                const args = (params.page && sortArg) ? {sortArg: sortArg, pageArg:pageArg} : (sortArg ? {sortArg: sortArg} : (params.page ? {pageArg: params.page} : null));
                return {
                    url: `/${name}/jokes/`,
                    params: args,
                }
            },
        }),
        getFavoritesByName: build.query({
            query: (name) => ({url: `/${name}/favorites/`}),
        }),
    })
})

export const {
    useGetUserByIDQuery,
    useGetUserByNameQuery,
    useGetFavoritesByNameQuery,
    useGetJokesByAuthorNameQuery,
} = userService;