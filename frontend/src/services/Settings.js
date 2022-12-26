import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/dist/query/react'

const apiHost = process.env.REACT_APP_API_HOST;


export const settingsService = createApi({
    reducerPath: 'settingsAPI',
    baseQuery: fetchBaseQuery({ baseUrl: `http://${apiHost}/settings`}),
    endpoints: (build) => ({
        getSettings: build.query({
            query: () => {
                const token = localStorage.getItem('userToken');
                return {
                    url: '/profile',
                    headers: {authorization: `${token}`},
                }
            },
        }),
        getDeveloperSettings: build.query({
            query: () => {
                const token = localStorage.getItem('userToken');
                return {
                    url: '/develop',
                    headers: {authorization: `${token}`},
                }
            }
        }),
        applyReport: build.mutation({
            query: (report_id) => {
                const token = localStorage.getItem('userToken');
                return {
                    url: `/develop/apply_report`,
                    method: 'PUT',
                    headers: {authorization: `${token}`},
                    params: report_id
                }
            },
        }),
        denyReport: build.mutation({
            query: (report_id) => {
                const token = localStorage.getItem('userToken');
                return {
                    url: `/develop/deny_report`,
                    method: 'PUT',
                    headers: {authorization: `${token}`},
                    params: report_id
                }
            },
        }),
    })
})

export const {
    useGetSettingsQuery,
    useApplyReportMutation,
    useDenyReportMutation
} = settingsService;