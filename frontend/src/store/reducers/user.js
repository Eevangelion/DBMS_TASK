// import { createSlice } from '@reduxjs/toolkit';
// import { getUserInfo } from '../actions/auth';
// import { useAddJokeToFavoritesMutation,
//     useRemoveJokeFromFavoritesMutation
// } from '../../services/joke';
// import {
//     useGetUserByNameQuery
// } from '../../services/user';

// const initialState = {
//     loading: false
// };

// export const userSlice = createSlice({
//     name: 'user',
//     initialState,
//     reducers: {
//         clean: () => initialState,
//     },
//     extraReducers: (builder) => {
//         builder
//             .addCase(getUserInfo.fulfilled, (state, action) => {
//                 const payload = action.payload;
//                 state.favoriteJokes = payload.favoriteJokes;
//                 state.username = payload.username;
//                 state.loading = false;
//             })
//             .addCase(useAddJokeToFavoritesMutation.rejected, (state, action) => {
//                 state.error = action.payload;
//                 state.loading = false;
//             })     
//             .addCase(useRemoveJokeFromFavoritesMutation.rejected, (state, action) => {
//                 state.error = action.payload;
//                 state.loading = false;
//             })   
//             .addCase(useGetUserByNameQuery.rejected, (state, action) => {
//                 state.loading = false;
//                 state.error = action.payload;
//             });
//     }
// });

// export const { clean } = userSlice.actions;
// export default userSlice.reducer;