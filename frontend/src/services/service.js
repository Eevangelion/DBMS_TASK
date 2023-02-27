import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;
const apiPort = process.env.REACT_APP_API_PORT;

export const jokeService = createApi({
    reducerPath: 'jokeAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}:${apiPort}`}),
    tagTypes: ['Jokes', 'Users', 'Tags', 'Reports'],
    endpoints: (build) => ({
        getUserByID: build.query({
            query: (id) => {
                const token = localStorage.getItem('access_token');
                return {
                    url: `/user/${id}/`,
                    headers: {authorization: `${token}`},
                }
            },
            providesTags: ['Users']
        }),
        getUserByName: build.query({
            query: (name) => {
                const token = localStorage.getItem('access_token');
                return {
                    url: `/user/${name}/data/`,
                    headers: {authorization: `${token}`},
                }
            },
            providesTags: ['Users']
        }),
        getJokesByAuthorName: build.query({
            query: ({name, ...params}) => {
                const token = localStorage.getItem('access_token');
                const sortArg = params.sortBy ? (params.sortBy === "top" ? (params.t ? params.t : 'day') : params.sortBy) : null;
                const pageArg = params.page;
                const args = (pageArg && sortArg) ? {sort: sortArg, page:pageArg} : (sortArg ? {sort: sortArg} : (pageArg ? {page: pageArg} : null));
                return {
                    url: `/user/${name}/jokes/`,
                    headers: {authorization: `${token}`},
                    params: args,
                }
            },
            providesTags: ['Jokes', 'Tags']
        }),
        getFavoritesByID: build.query({
            query: (id) => ({url: `/user/favorites/${id}/`}),
            providesTags: ['Jokes', 'Tags']
        }),
        getSubscribedByID: build.query({
            query: ({id, ...params}) => {
                const token = localStorage.getItem('access_token');
                const sortArg = params.sortBy ? (params.sortBy === "top" ? (params.t ? params.t : 'day') : params.sortBy) : null;
                const pageArg = params.page;
                const args = (pageArg && sortArg) ? {id: id, sort: sortArg, page:pageArg} : (sortArg ? {id: id, sort: sortArg} : (pageArg ? {id: id, page: pageArg} : {id: id}));
                return {
                    url: `/user/subscribed/`,
                    headers: {authorization: `${token}`},
                    params: args,
                }
            },
            providesTags: ['Jokes', 'Tags']
        }),
        getJokes: build.query({
            query: (...params) => {
                const token = localStorage.getItem('access_token');
                const parsedParams = params[0];
                const sortArg = parsedParams.sortBy ? (parsedParams.sortBy === "top" ? (parsedParams.t ? parsedParams.t : 'day') : parsedParams.sortBy) : null;
                const pageArg = parsedParams.page;
                const args = (pageArg && sortArg) ? {sort: sortArg, page:pageArg} : (sortArg ? {sort: sortArg} : (pageArg ? {page: pageArg} : null));
                return {
                    url: `/feed/`,
                    headers: {authorization: `${token}`},
                    params: args,
                }
            },
            providesTags: ['Jokes', 'Tags']
        }),
        getJokeByID: build.query({
            query: (id) => {
                const token = localStorage.getItem("access_token");
                return {
                    url: `/joke/${id}/`,
                    header: {authorization: `${token}`},
                }
            }
        }),
        getSearchResult: build.query({
            query: ({q, t, ...params}) => {
                const token = localStorage.getItem('access_token');
                const query = q;
                const type = t;
                const sortArg = params.sortBy ? (params.sortBy === "top" ? (params.t ? params.t : 'day') : params.sortBy) : null;
                const pageArg = params.page;
                const args = (pageArg && sortArg) ? {sort: sortArg, page:pageArg} : (sortArg ? {sort: sortArg} : (pageArg ? {page: pageArg} : null));
                return {
                    url: `/search/${type}/${query}/`,
                    headers: {authorization: `${token}`},
                    params: args,
                }
            },
            providesTags: ['Jokes', 'Tags', 'Users']
        }),
        createJoke: build.mutation({
            query: (body) => {
                const token = localStorage.getItem('access_token');
                return {
                    url: `/joke/create/`,
                    method: 'POST',
                    headers: { authorization: `${token}`},
                    body: body
                };
            },
            invalidatesTags: ['Jokes']
        }),
        deleteJoke: build.mutation({
            query: (joke_id) => {
                const token = localStorage.getItem('access_token');
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
                const token = localStorage.getItem('access_token');
                const userID = Number(localStorage.getItem('userID'));
                return {
                    url: `/joke/addToFavorites/`,
                    method: 'POST',
                    headers: {authorization: `${token}`},
                    body: {user_id: userID, joke_id: joke_id},
                }
            },
            invalidatesTags: ['Jokes']
        }),
        removeJokeFromFavorites: build.mutation({
            query: (joke_id) => {
                const token = localStorage.getItem('access_token');
                const userID = Number(localStorage.getItem('userID'));
                return {
                    url: `/joke/removeFromFavorites/`,
                    method: 'DELETE',
                    headers: {authorization: `${token}`},
                    body: {user_id: userID, joke_id: joke_id},
                }
            },
            invalidatesTags: ['Jokes']
        }), 
        getTagsByJokeID: build.query({
            query: (joke_id) => {
                const token = localStorage.getItem('access_token');
                return {
                    url: `/joke/tags/${joke_id}/`,
                    headers: {authorization: `${token}`},
                }
            },
            invalidatesTags: ['Tags']
        }),
        getTags: build.query({
            query: () => {
                const token = localStorage.getItem('access_token');
                return {
                    url: `/tag/`,
                    headers: {authorization: `${token}`},
                }
            },
            invalidatesTags: ['Tags']
        }),
        createTag: build.mutation({
            query: ({name, id}) => {
                const token = localStorage.getItem('access_token');
                return {
                    url: `/tag/create/`,
                    headers: {authorization: `${token}`},
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
                const token = localStorage.getItem('access_token');
                return {
                    url: `/tag/delete/`,
                    method: 'DELETE',
                    headers: {authorization: `${token}`},
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
                const token = localStorage.getItem('access_token');
                return {
                    url: `/joke/addTag/`,
                    method: `POST`,
                    headers: {authorization: `${token}`},
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
                const token = localStorage.getItem('access_token');
                return {
                    url: `/joke/removeTag/`,
                    method: `POST`,
                    headers: {authorization: `${token}`},
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
                const token = localStorage.getItem('access_token');
                const userID = Number(localStorage.getItem("userID"));
                return {
                    url: `/report/create/`,
                    method: 'POST',
                    headers: {authorization: `${token}`},
                    body: {
                        description: description,
                        receiver_joke_id: jokeID,
                        sender_id: userID,
                    }
                }
            },
            invalidatesTags: ['Users', 'Reports']
        }),
        removeReport: build.mutation({
            query: (id) => {
                const token = localStorage.getItem('access_token');
                return {
                    url: `/report/delete/`,
                    method: 'DELETE',
                    headers: {authorization: `${token}`},
                    body: {
                        report_id: id,
                    }
                }
            },
            invalidatesTags: ['Reports']
        }),
        getReports: build.query({
            query: () => {
                const token = localStorage.getItem('access_token');
                return {
                    url: `/report/`,
                    method: 'GET',
                    headers: {authorization: `${token}`},
                }
            },
            providesTags: ['Reports']
        }),
        applyReport: build.mutation({
            query: (reportID) => {
                const token = localStorage.getItem('access_token');
                return {
                    url: `/report/apply/`,
                    method: 'POST',
                    headers: {authorization: `${token}`},
                    body: {
                        report_id: reportID,
                    }
                }
            },
            invalidatesTags: ['Users', 'Reports']
        }),
        subscribeToUser: build.mutation({
            query: (receiverID) => {
                const token = localStorage.getItem('access_token');
                const userID = localStorage.getItem("userID");
                return {
                    url: `/user/subscribe/`,
                    method: 'POST',
                    headers: {authorization: `${token}`},
                    body: {
                        receiver_id: Number(receiverID),
                        sender_id: Number(userID),
                    }
                }
            },
            invalidatesTags: ['Users']
        }),
        changeUserName: build.mutation({
            query: (name) => {
                const token = localStorage.getItem('access_token');
                const userID = localStorage.getItem("userID");
                return {
                    url: `/user/change_name/`,
                    method: 'PUT',
                    headers: {authorization: `${token}`},
                    body: {
                        name: name,
                        user_id: userID,
                    }
                }
            },
            invalidatesTags: ['Users']
        }),
        changePassword: build.mutation({
            query: (password) => {
                const token = localStorage.getItem('access_token');
                const userID = localStorage.getItem("userID");
                return {
                    url: `/user/change_password/`,
                    method: 'PUT',
                    headers: {authorization: `${token}`},
                    body: {
                        transformed_password: password,
                        user_id: userID,
                    }
                }
            }
        }),
    })
})


export const {
    useGetUserByIDQuery,
    useGetUserByNameQuery,
    useGetFavoritesByIDQuery,
    useGetSubscribedByIDQuery,
    useGetJokesByAuthorNameQuery,
    useGetJokesQuery,
    useGetJokeByIDQuery,
    useGetSearchResultQuery,
    useCreateJokeMutation,
    useDeleteJokeMutation,
    useAddJokeToFavoritesMutation,
    useRemoveJokeFromFavoritesMutation,
    useGetTagsByJokeIDQuery,
    useGetTagsQuery,
    useCreateTagMutation,
    useDeleteTagMutation,
    useGetReportsQuery,
    useCreateReportMutation,
    useRemoveReportMutation,
    useSubscribeToUserMutation,
    useChangeUserNameMutation,
    useChangePasswordMutation,
    useApplyReportMutation,
} = jokeService;