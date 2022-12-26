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
            query: (joke_id) => {
                const token = localStorage.getItem('userToken');
                return {
                    url: `/delete`,
                    method: 'DELETE',
                    headers: {authorization: `${token}`},
                    body: joke_id
                }
            },
        }),
        addJokeToFavorites: build.mutation({
            query: (joke_id) => {
                const token = localStorage.getItem('userToken');
                const user_id = localStorage.getItem('userID');
                return {
                    url: `/addToFavorites`,
                    method: 'PUT',
                    headers: {authorization: `${token}`},
                    params: {user_id, joke_id},
                }
            }
        }),
        removeJokeFromFavorites: build.mutation({
            query: (joke_id) => {
                const token = localStorage.getItem('userToken');
                const user_id = localStorage.getItem('userID');
                return {
                    url: `/removeFromFavorites`,
                    method: 'PUT',
                    headers: {authorization: `${token}`},
                    params: {user_id, joke_id},
                }
            }
        }), 
        getTagsByJokeID: build.query({
            query: (joke_id) => {
                return {
                    url: `/tags`,
                    params: joke_id,
                }
            }
        }),
    })
})


export const {
    useCreateJokeMutation,
    useDeleteJokeMutation,
    useAddJokeToFavoritesMutation,
    useRemoveJokeFromFavoritesMutation,
    useGetTagsByJokeIDLazyQuery,
} = jokeService;