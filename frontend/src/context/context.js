import React, { createContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { useGetTokenMutation } from '../services/auth';
import { useGetUnbanDateQuery } from '../services/service';
import LoadingModal from '../components/LoadingModal/LoadingModal';

export const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {

    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [refreshTokens] = useGetTokenMutation();
    const expTime = localStorage.getItem("token_exp_time");
    const token = localStorage.getItem("access_token");


    const signOut = () => {
        localStorage.clear();
        navigate("/login");
        window.scrollTo(0,0);
    }

    useEffect(() => {
        if (token && expTime - Date.now()/1000 < 0) {
            refreshTokens().then((response) => {
                console.log(response);
                if (response.error && response.error.status === 401) {
                    localStorage.clear();
                    navigate("/login");
                    window.scrollTo(0,0);
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
                localStorage.setItem("userID", data.user_id);
                localStorage.setItem("userName", data.username);
                localStorage.setItem("userRole", data.role);
                localStorage.setItem("access_token", accessToken);
                localStorage.setItem("token_exp_time", data.exp);
                localStorage.setItem("refresh_token", refreshToken);
                return;
            })
        }
    }, [expTime, token, dispatch, refreshTokens, navigate]);

    const value = { signOut };

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

    const userRole = localStorage.getItem("userRole");

    useEffect(() => {
        if (userRole !== 'admin') {
            navigate('../');
            window.scrollTo(0,0);
        }
    }, [userRole, navigate])

    return children;
}

export const RequireNotBanned = ({children}) => {
    const navigate = useNavigate();
    const {
        data: unparsedUnbanDate,
        isLoading
    } = useGetUnbanDateQuery();

    useEffect(() => {
        const unbanDate = Math.round((Date.now() - Date.parse(unparsedUnbanDate))/1000);
        if (!isLoading && unbanDate < 0) {
            navigate('/banned/');
            window.scrollTo(0, 0);
        }
    }, [isLoading, unparsedUnbanDate, navigate])

    if (isLoading) {
        return <LoadingModal />
    }

    return children;
}