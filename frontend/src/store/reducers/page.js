import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    feedIsActive: true,
    userPageIsActive: true,
    searchPageIsActive: true,
    subscribesIsActive: true,
    reportListIsActive: true,
}

export const pagesSlice = createSlice({
    name: 'pages',
    initialState,
    reducers: {
        selectPage: (state, action) => {
            switch (action.payload.page) {
            case 'feed': state.feedIsActive = action.payload.state;break;
            case 'userPage': state.userPageIsActive = action.payload.state;break;
            case 'searchPage': state.searchPageIsActive = action.payload.state;break;
            case 'reportList': state.reportListIsActive = action.payload.state;break;
            default: state.subscribesIsActive = action.payload.state;break;
            }
        },  
    }
})

export const {
    selectPage
} = pagesSlice.actions;
export default pagesSlice.reducer;