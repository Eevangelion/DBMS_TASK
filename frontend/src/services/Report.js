import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;


export const reportService = createApi({
    reducerPath: 'reportAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}/report`}),
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