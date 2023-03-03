import React, { createContext } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from '../hooks/redux';
import { getAccessToken, getJWToken, getUserInfo } from '../store/actions/authActions';
import { logout } from '../store/reducers/AuthSlice';
import { clean } from '../store/reducers/UserSlice';

export const AuthContext = createContext(null);

export function AuthProvider({ children }) {

    const dispatch = useAppDispatch(); 
    const navigate = useNavigate();
    const user = useAppSelector(state => state.VkAuth.user);
    const token = localStorage.getItem('access_token');    

    const signin = () => { 
        return dispatch(getAccessToken())
            .then(() => 
                dispatch(getJWToken())
                    .then(() => 
                        dispatch(getUserInfo())
                    )
            );                 
    };

    if (token && !user) {
        dispatch(getUserInfo()).catch(() => signin());
    };

    const signout = () => {
        dispatch(clean());
        navigate('../');
        window.scrollTo(0,0);
        return dispatch(logout());
    };

    const value = { user, signin, signout };

    return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};