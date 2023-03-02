import React, {useEffect, useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import styles from "../styles/Feed.module.css";
import JokePost from "../components/JokePost/JokePost";
import JokeSorter from "../components/Sorter/Sorter";
import TopPanel from "../components/TopPanel/TopPanel";
import PageSelector from "../components/PageSelector/PageSelector";
import { useGetJokesQuery } from "../services/service";
import { useGetTokenMutation } from "../services/auth";
import LoadingModal from "../components/LoadingModal/LoadingModal";

const Feed = (props) => {
    const [pageState, setPage] = useState(1);
    const [pageContent, setContent] = useState(<></>);
    const activeButton = useSelector(state => state.buttonsReducer.sort);
    const isActive = useSelector(state => state.pagesReducer.feedIsActive);
    const expTime = localStorage.getItem("token_exp_time");

    const {
        data: response,
        isLoading: loadingJokes,
        error,
    } = useGetJokesQuery({page: pageState, sortBy: activeButton});
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
    useEffect(() => {
        if (!loadingJokes) {
            const {jokes, amount} = response; 
            if (!jokes) {
                setContent(
                    <>
                        <div className={styles.txt}>Никто пока ничего не публиковал</div>
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
    const amount = response.amount;

    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                <div className={styles.feed}>
                    <JokeSorter />
                    {pageContent}
                    <Pagination className={styles.pagination} count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} shape="rounded"/>
                </div>
                <PageSelector pageState={true} />
            </div>
        </div>
    );
}

export default Feed;