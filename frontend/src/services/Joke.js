import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;
const apiPort = process.env.REACT_APP_API_PORT;

export const jokeService = createApi({
    reducerPath: 'jokeAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}:${apiPort}`}),
    tagTypes: ['Jokes'],
    endpoints: (build) => ({
        getUserByID: build.query({
            query: (id) => ({url: `/user/${id}/`})
        }),
        getUserByName: build.query({
            query: (name) => ({url: `/user/${name}/data/`}),
        }),
        getJokesByAuthorName: build.query({
            query: ({name, ...params}) => {
                const sortArg = params.sortBy ? (params.sortBy === "top" ? (params.t ? params.t : 'day') : params.sortBy) : null;
                const pageArg = params.page;
                const args = (pageArg && sortArg) ? {sort: sortArg, page:pageArg} : (sortArg ? {sort: sortArg} : (pageArg ? {page: pageArg} : null));
                return {
                    url: `/user/${name}/jokes/`,
                    params: args,
                }
            },
            providesTags: ['Jokes']
        }),
        getFavoritesByID: build.query({
            query: (id) => ({url: `/user/favorites/${id}/`}),
            providesTags: ['Jokes']
        }),
        getSubscribedByID: build.query({
            query: ({id, ...params}) => {
                console.log(id, params);
                const sortArg = params.sortBy ? (params.sortBy === "top" ? (params.t ? params.t : 'day') : params.sortBy) : null;
                const pageArg = params.page;
                const args = (pageArg && sortArg) ? {id: id, sort: sortArg, page:pageArg} : (sortArg ? {id: id, sort: sortArg} : (pageArg ? {id: id, page: pageArg} : {id: id}));
                return {
                    url: `/user/subscribed/`,
                    params: args,
                }
            },
            providesTags: ['Jokes']
        }),
        getJokes: build.query({
            query: (...params) => {
                const parsedParams = params[0];
                const sortArg = parsedParams.sortBy ? (parsedParams.sortBy === "top" ? (parsedParams.t ? parsedParams.t : 'day') : parsedParams.sortBy) : null;
                const pageArg = parsedParams.page;
                const args = (pageArg && sortArg) ? {sort: sortArg, page:pageArg} : (sortArg ? {sort: sortArg} : (pageArg ? {page: pageArg} : null));
                return {
                    url: `/feed/`,
                    params: args,
                }
            },
            providesTags: ['Jokes']
        }),
        createJoke: build.mutation({
            query: (body) => {
                //const token = localStorage.getItem('userToken');
                return {
                    url: `/joke/create/`,
                    method: 'POST',
                    //headers: { authorization: `${token}`},
                    body: body
                };
            },
            invalidatesTags: ['Jokes']
        }),
        deleteJoke: build.mutation({
            query: (joke_id) => {
                const token = localStorage.getItem('userToken');
                return {
                    url: `/joke/delete/`,
                    method: 'DELETE',
                    headers: {authorization: `${token}`},
                    body: joke_id
                }
            },
            invalidatesTags: ['Jokes']
        }),
        addJokeToFavorites: build.mutation({
            query: (joke_id) => {
                // const token = localStorage.getItem('userToken');
                const userID = Number(localStorage.getItem('userID'));
                return {
                    url: `/joke/addToFavorites/`,
                    method: 'POST',
                    // headers: {authorization: `${token}`},
                    body: {user_id: userID, joke_id: joke_id},
                }
            },
            invalidatesTags: ['Jokes']
        }),
        removeJokeFromFavorites: build.mutation({
            query: (joke_id) => {
                // const token = localStorage.getItem('userToken');
                const userID = Number(localStorage.getItem('userID'));
                return {
                    url: `/joke/removeFromFavorites/`,
                    method: 'DELETE',
                    // headers: {authorization: `${token}`},
                    body: {user_id: userID, joke_id: joke_id},
                }
            },
            invalidatesTags: ['Jokes']
        }), 
        getTagsByJokeID: build.query({
            query: (joke_id) => {
                return {
                    url: `/joke/tags/${joke_id}/`,
                }
            },
            invalidatesTags: ['Tags']
        }),
        getTags: build.query({
            query: () => ({url: `/tag/`}),
            invalidatesTags: ['Tags']
        }),
        createTag: build.mutation({
            query: ({name, id}) => {
                return {
                    url: `/tag/create/`,
                    method: 'POST',
                    body: {
                        name: name, 
                        user_id: Number(id),
                    },
                }
            },
            providesTags: ['Tags'],
        }),
        deleteTag: build.mutation({
            query: ({name, id}) => {
                return {
                    url: `/tag/delete/`,
                    method: 'DELETE',
                    body: {
                        name: name, 
                        user_id: Number(id),
                    },
                }
            },
            providesTags: ['Tags'],
        }),
        addTagToJoke: build.mutation({
            query: ({tagID, jokeID}) => {
                return {
                    url: `/joke/addTag/`,
                    method: `POST`,
                    body: {
                        tag_id: tagID,
                        joke_id: jokeID,
                    }
                }
            },
            providesTags: ['Tags']
        }),
        removeTagFromJoke: build.mutation({
            query: ({tagID, jokeID}) => {
                return {
                    url: `/joke/removeTag/`,
                    method: `POST`,
                    body: {
                        tag_id: tagID,
                        joke_id: jokeID,
                    }
                }
            },
            providesTags: ['Tags']
        }),
        createReport: build.mutation({
            query: ({description, jokeID}) => {
                const userID = localStorage.getItem("userID");
                return {
                    url: `/report/create/`,
                    method: 'POST',
                    body: {
                        description: description,
                        receiver_joke_id: jokeID,
                        sender_id: userID,
                    }
                }
            }
        }),
        removeReport: build.mutation({
            query: ({id}) => {
                const userID = localStorage.getItem("userID");
                return {
                    url: `/report/delete/`,
                    method: 'DELETE',
                    body: {
                        report_id: id,
                        user_id: userID,
                    }
                }
            }
        })
    })
})


export const {
    useGetUserByIDQuery,
    useGetUserByNameQuery,
    useGetFavoritesByIDQuery,
    useGetSubscribedByIDQuery,
    useGetJokesByAuthorNameQuery,
    useGetJokesQuery,
    useCreateJokeMutation,
    useDeleteJokeMutation,
    useAddJokeToFavoritesMutation,
    useRemoveJokeFromFavoritesMutation,
    useGetTagsByJokeIDQuery,
    useGetTagsQuery,
    useCreateTagMutation,
    useDeleteTagMutation,
} = jokeService;