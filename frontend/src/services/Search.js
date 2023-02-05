import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;
const apiPort = process.env.REACT_APP_API_PORT;

export const searchService = createApi({
    reducerPath: 'searchAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}:${apiPort}/search`}),
    endpoints: (build) => ({
        getJokes: build.query({
            query: (q, t, ...params) => {
                const queryArg = q;
                const typeArg = t;
                const pageArg = params.page;
                const args = params.page ? {queryArg, typeArg, pageArg} : {queryArg, typeArg};
                return {
                    url: '',
                    params: args
                }
            },
        }),
    })
})

export const {
    useGetJokesQuery
} = searchService;