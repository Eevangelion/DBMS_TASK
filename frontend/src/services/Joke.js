import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;


export const jokeService = createApi({
    reducerPath: 'jokeAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}/joke`}),
    endpoints: (build) => ({
        createJoke: build.mutation({
            query: (joke) => {
                const token = localStorage.getItem('userToken');
                return {
                    url: `/create`,
                    method: 'POST',
                    headers: { authorization: `${token}`},
                    body: joke
                };
            },
        }),
        deleteJoke: build.mutation({
            query: (id) => {
                const token = localStorage.getItem('userToken');
                return {
                    url: `/delete`,
                    method: 'DELETE',
                    headers: {authorization: `${token}`},
                    body: id
                }
            },
        }),
        addJokeToFavorites: build.mutation({
            query: (name, id) => {
                const token = localStorage.getItem('userToken');
                return {
                    url: `/addToFavorites`,
                    method: 'PUT',
                    headers: {authorization: `${token}`},
                }
            }
        }),
        removeJokeFromFavorites: build.mutation({
            query: (name, id) => {
                const token = localStorage.getItem('userToken');
                return {
                    url: `/removeFromFavorites`,
                    method: 'PUT',
                    headers: {authorization: `${token}`},
                }
            }
        })
    })
})


export const {
    useCreateJokeMutation,
    useDeleteJokeMutation,
    useAddJokeToFavoritesMutation,
    useRemoveJokeFromFavoritesMutation,
} = jokeService;