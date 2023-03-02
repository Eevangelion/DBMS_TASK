import React, {useEffect, useState} from "react";
import { useSelector } from "react-redux";
import Pagination from '@mui/material/Pagination';
import {useParams, useSearchParams} from "react-router-dom";
import styles from "../styles/SearchPage.module.css";
import JokePost from "../components/JokePost/JokePost";
import Sorter from "../components/Sorter/Sorter";
import TopPanel from "../components/TopPanel/TopPanel";
import { useGetSearchResultQuery } from "../services/service";
import UserPost from "../components/UserPost/UserPost";
import LoadingModal from "../components/LoadingModal/LoadingModal";
import { useGetTokenMutation } from "../services/auth";

const SearchPage = (props) => {

    const [pageState, setPage] = useState(1);
    const [pageContent, setContent] = useState(<></>);
    const activeButton = useSelector(state => state.buttonsReducer.sort);
    const isActive = useSelector(state => state.pagesReducer.searchPageIsActive);
    const expTime = localStorage.getItem("token_exp_time");

    const [searchParams] = useSearchParams();
    const { type: typeArg } = useParams();
    const queryArg = searchParams.get('query');


    const {
        data: response,
        isLoading: loadingSearch,
        error,
    } = useGetSearchResultQuery({q: queryArg, t: typeArg, page: pageState, sortBy: activeButton});

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
        if (!loadingSearch) {
            if (typeArg === 'keyword' || typeArg === 'tag') {
                const jokes = response ? response.jokes : []; 
        
                if (!jokes) {
                    setContent(
                        <>                                        
                            <Sorter />
                            <div className={styles.txt}>По данному запросу ничего не найдено</div>
                        </>
                    );
                } else {
                    const posts = jokes.map((joke) =>
                    {
                        return <JokePost key={joke.id} joke={joke}/>
                    });
                    setContent(
                        <>
                            <Sorter />
                            <div className={styles.txt}>Результаты поиска по {typeArg === 'keyword' ? `ключевому слову ${queryArg}` : `тэгу ${queryArg}`}</div> <br/>
                            <ul className={styles.jokePostList}>
                                {posts}
                            </ul>
                        </>
                    );
                }
            } else {
                const people = response; 

                if (!people) {
                    setContent(
                        <>
                            <div className={styles.txt}>По данному запросу ничего не найдено</div>
                        </>
                    );
                } else {
                    const posts = people.map((user) =>
                    {
                        return <UserPost key={user.id} user={user}/>
                    });
                    setContent(
                        <>
                            <div className={styles.txt}>Результаты поиска пользователей по имени</div> <br/>
                            <ul className={styles.peoplePostList}>
                                {posts}
                            </ul>
                        </>
                    );
                }
            }
        }
    }, [loadingSearch, response, typeArg, queryArg]);

    if (loadingSearch) {
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
    let amount;
    if (typeArg === 'people') {
        amount = response ? response.length : 0;
    } else {
        amount = response.amount;
    }
    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.info} style={isActive ? {} : {backgroundColor: "#676a6c"}}>
                <div className={styles.feed}>
                    {pageContent}
                    <Pagination className={styles.pagination} count={Math.ceil(amount/5)} onChange={(e, value) => setPage(value)} shape="rounded"/>
                </div>
            </div>
        </div>
    );
}

export default SearchPage;