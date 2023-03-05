import React, { createContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useSelector, useDispatch } from 'react-redux';
import { useGetTokenMutation } from '../services/auth';
import { selectUser } from '../store/reducers/user';

export const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {

    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [refreshTokens] = useGetTokenMutation();
    const expTime = localStorage.getItem("token_exp_time");
    const token = localStorage.getItem("access_token");
    const user = useSelector(state => state.userReducer);


    const signOut = () => {
        localStorage.clear();
        dispatch(selectUser({data: 'userID', state: undefined}));
        dispatch(selectUser({data: 'username', state: undefined}));
        dispatch(selectUser({data: 'userRole', state: undefined}));
        navigate("/login");
        window.scrollTo(0,0);
    }

    if (token && expTime - Date.now()/1000 < 0) {
        refreshTokens().then((response) => {
            if (response.error) {
                signOut();
                return;
            }
            const tokens = response.data;
            const accessToken = tokens.jwt_token;
            const refreshToken = tokens.refresh_token;
            const base64Url = accessToken.split('.')[1];
            const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
            const jsonPayload = decodeURIComponent(window.atob(base64).split('').map((c) => {
                return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
            }).join(''));
            const data = JSON.parse(jsonPayload);
            localStorage.setItem("access_token", accessToken);
            localStorage.setItem("token_exp_time", data.exp);
            localStorage.setItem("refresh_token", refreshToken);
            dispatch(selectUser({data: 'userID', state: data.user_id}));
            dispatch(selectUser({data: 'userName', state: data.username}));
            dispatch(selectUser({data: 'userRole', state: data.role}));
            return;
        })
    }

    const value = { user, signOut };

    return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const RequireAuth = ({children}) => {
    const token = localStorage.getItem("access_token");
    const navigate = useNavigate();
    useEffect(()=> {
        if (!token)
            navigate("/login");
    }, [token, navigate]);
    return children;
}

export const RequireAdmin = ({children}) => {
    const token = localStorage.getItem("access_token");
    const navigate = useNavigate();

    if (!token) {
        navigate("/login");
    }

    const user = useSelector(state => state.userReducer);

    useEffect(() => {
        if (user && user.userRole !== 'admin') {
            navigate('../');
            window.scrollTo(0,0);
        }
    }, [user, navigate])

    return children;
}