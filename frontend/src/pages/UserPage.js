import React, {useEffect, useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import { useParams } from "react-router-dom";
import styles from "../styles/UserPage.module.css";
import JokePost from "../components/JokePost/JokePost";
import Profile from "../components/Profile/Profile";
import TopPanel from "../components/TopPanel/TopPanel";
import JokeSorter from "../components/Sorter/Sorter";
import {useGetJokesByAuthorNameQuery} from "../services/service";
import LoadingModal from "../components/LoadingModal/LoadingModal";
import { useGetTokenMutation } from "../services/auth";

const UserPage = (props) => {

    const [pageState, setPage] = useState(1);
    const [pageContent, setContent] = useState(<></>);
    const activeButton = useSelector(state => state.buttonsReducer.sort);
    const isActive = useSelector(state => state.pagesReducer.userPageIsActive);
    const expTime = localStorage.getItem("token_exp_time");

    const {username} = useParams();

    const {
        data: response,
        isLoading: loadingJokes,
        error,
    } = useGetJokesByAuthorNameQuery({name: username, page: pageState, sortBy: activeButton});

    const [refreshTokens] = useGetTokenMutation();
    useEffect(() => {
        if (expTime - Date.now()/1000 < 0) {
            refreshTokens().then((response) => {
                const tokens = response.data;
                const accessToken = tokens.jwt_token;
                const refreshToken = tokens.refresh_token;
                const base64Url = accessToken.split('.')[1];
                const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
                const jsonPayload = decodeURIComponent(window.atob(base64).split('').map((c) => {
                    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
                }).join(''));
                const user = JSON.parse(jsonPayload);
                localStorage.setItem("userID", user.user_id);
                localStorage.setItem("userName", user.username);
                localStorage.setItem("userRole", user.role);
                localStorage.setItem("access_token", accessToken);
                localStorage.setItem("token_exp_time", user.exp);
                localStorage.setItem("refresh_token", refreshToken);
            })
        }
    }, [expTime, refreshTokens]);

    useEffect(()=>{
        if (!loadingJokes) {
            const {jokes, amount} = response; 
            if (!jokes) {
                setContent(
                    <>
                        <div className={styles.txt}>Пользователь пока ничего не опубликовал</div>
                    </>
                );
            } else {
                const posts = jokes.map((joke) =>
                {
                    return <JokePost key={joke.id} joke={joke}/>
                });
                setContent(
                    <>
                        <div className={styles.txt}>Всего опубликовано: {amount}</div> <br/>
                        <ul className={styles.jokePostList}>
                            {posts}
                        </ul>
                    </>
                );
            }
        }
    }, [loadingJokes, response]);

    if (loadingJokes) {
        return <LoadingModal />;
    }
    if (error) {
        if (error && 'status' in error) {
            const errMsg = 'error' in error ? error.error : JSON.stringify(error.data);

            return (
                <div>
                    <div>An error has occurred:</div>
                    <div>{errMsg}</div>
                </div>
            );
        } else {
            return <div>{error?.message}</div>;
        }
    }
    const amount = response ? response.amount : 0;
    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.userInfo}  style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                <div className={styles.feed}>
                    <JokeSorter />
                    {pageContent}
                    <Pagination className={styles.pagination} count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} shape="rounded"/>
                </div>
                <Profile username={username} />
            </div>
        </div>
    );
}

export default UserPage;