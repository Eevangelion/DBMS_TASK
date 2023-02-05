import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;
const apiPort = process.env.REACT_APP_API_PORT;

export const reportService = createApi({
    reducerPath: 'reportAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}:${apiPort}/report`}),
    endpoints: (build) => ({
        createReport: build.mutation({
            query: (joke_id, report) => {
                const token = localStorage.getItem('userToken');
                return {
                    url: `/create`,
                    method: 'POST',
                    headers: {authorization: `${token}`},
                    body: report,
                    params: joke_id,
                }
            },
        })
    })
})

export const {
    useCreateReportMutation
} = reportService;