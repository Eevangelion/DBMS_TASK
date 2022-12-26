import {combineReducers, configureStore} from '@reduxjs/toolkit';
import { setupListeners } from '@reduxjs/toolkit/dist/query';
import buttonsReducer from './reducers/buttons';
// import authReducer from './reducers/auth';
// import userReducer from './reducers/user';
import { userService } from '../services/User';
import { jokeService } from '../services/Joke';
import { feedService } from '../services/Feed';
import { reportService } from '../services/Report';
import { searchService } from '../services/Search';
import { settingsService } from '../services/Settings';

const mainReducer = combineReducers({
    buttonsReducer,
    [userService.reducerPath]: userService.reducer,
    [jokeService.reducerPath]: jokeService.reducer,
    [feedService.reducerPath]: feedService.reducer,
    [reportService.reducerPath]: reportService.reducer,
    [searchService.reducerPath]: searchService.reducer,
    [settingsService.reducerPath]: settingsService.reducer,
    // userSlice: userReducer,
    // gitAuth: authReducer
})

export const setupStore = () => configureStore({
    reducer: mainReducer,
    middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(
        userService.middleware,
        jokeService.middleware,
        feedService.middleware,
        reportService.middleware,
        searchService.middleware,
        settingsService.middleware,
    )
});


setupListeners(setupStore().dispatch);