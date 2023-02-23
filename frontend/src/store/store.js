import {combineReducers, configureStore} from '@reduxjs/toolkit';
import { setupListeners } from '@reduxjs/toolkit/dist/query';
import buttonsReducer from './reducers/buttons';
import pagesReducer from './reducers/page';
// import authReducer from './reducers/auth';
// import userReducer from './reducers/user';
import { jokeService } from '../services/service';

const mainReducer = combineReducers({
    buttonsReducer,
    pagesReducer,
    [jokeService.reducerPath]: jokeService.reducer,
    // userSlice: userReducer,
    // gitAuth: authReducer
})

export const setupStore = () => configureStore({
    reducer: mainReducer,
    middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(
        jokeService.middleware,
    )
});


setupListeners(setupStore().dispatch);