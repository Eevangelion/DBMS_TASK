import {combineReducers, configureStore} from '@reduxjs/toolkit';
import { setupListeners } from '@reduxjs/toolkit/dist/query';
import buttonsReducer from './reducers/buttons';
import pagesReducer from './reducers/page';
// import authReducer from './reducers/auth';
// import userReducer from './reducers/user';
import { jokeService } from '../services/Joke';
import { reportService } from '../services/Report';
import { searchService } from '../services/Search';
import { settingsService } from '../services/Settings';

const mainReducer = combineReducers({
    buttonsReducer,
    pagesReducer,
    [jokeService.reducerPath]: jokeService.reducer,
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
        jokeService.middleware,
        reportService.middleware,
        searchService.middleware,
        settingsService.middleware,
    )
});


setupListeners(setupStore().dispatch);