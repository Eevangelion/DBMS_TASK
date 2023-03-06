import {combineReducers, configureStore} from '@reduxjs/toolkit';
import { setupListeners } from '@reduxjs/toolkit/dist/query';
import buttonsReducer from './reducers/buttons';
import pagesReducer from './reducers/page';
import { authService } from '../services/auth';
import { jokeService } from '../services/service';

const mainReducer = combineReducers({
    buttonsReducer,
    pagesReducer,
    [jokeService.reducerPath]: jokeService.reducer,
    [authService.reducerPath]: authService.reducer,
})

export const setupStore = () => configureStore({
    reducer: mainReducer,
    middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(
        jokeService.middleware,
        authService.middleware,
    )
});


setupListeners(setupStore().dispatch);