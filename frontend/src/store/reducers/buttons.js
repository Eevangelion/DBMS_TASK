import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    sort: 'new',
    tag: '',
}

export const buttonsSlice = createSlice({
    name: 'buttons',
    initialState,
    reducers: {
        selectSort: (state, action) => {
            state.sort = action.payload
        },  
        selectTag: (state, action) => {
            state.tag = action.payload;
        },
    }
})

export const {
    selectSort,
    selectTag
} = buttonsSlice.actions;
export default buttonsSlice.reducer;