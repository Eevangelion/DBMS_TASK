import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;


export const feedService = createApi({
    reducerPath: 'feedAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}/feed`}),
    endpoints: (build) => ({
        getJokes: build.query({
            query: (...params) => {
                const token = localStorage.getItem('userToken');
                const sortArg = params.sortBy ? (params.sortBy === "top" ? params.t : params.sortBy) : 'new';
                const pageArg = params.page;
                const args = params.page ? {sortArg, pageArg} : {sortArg}
                return {
                    url: '',
                    headers: {authorization: `${token}`},
                    params: args
                }
            },
        }),
    })
})