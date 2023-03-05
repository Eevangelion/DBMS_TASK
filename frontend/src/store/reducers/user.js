import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    userID: undefined,
    userName: undefined,
    userRole: undefined,
}

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        selectUser: (state, action) => {
            switch (action.payload.data) {
            case 'userID': state.userID = action.payload.state;break;
            case 'userName': state.userName = action.payload.state;break;
            case 'userRole': state.userRole = action.payload.state;break;
            default:break;
            }
        },  
    }
})


export const {
    selectUser
} = userSlice.actions;
export default userSlice.reducer;