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
import ErrorPage from "./ErrorPage";

const UserPage = (props) => {

    const [pageState, setPage] = useState(1);
    const [pageContent, setContent] = useState(<></>);
    const activeButton = useSelector(state => state.buttonsReducer.sort);
    const isActive = useSelector(state => state.pagesReducer.userPageIsActive);

    const {username} = useParams();

    const {
        data: response,
        isLoading: loadingJokes,
        error,
    } = useGetJokesByAuthorNameQuery({name: username, page: pageState, sortBy: activeButton});

    useEffect(()=>{
        if (response && !loadingJokes) {
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

    if ((error && error.status === 400) && response === undefined) {
        return <ErrorPage />;
    }
    if (error) {
        if (error && 'status' in error) {
            const errMsg = 'error' in error ? error.error : JSON.stringify(error.data);

            return (
                <div className="error-page">
                    <div className="error-text">
                        <div>An error has occurred:</div>
                        <div>{errMsg}</div>
                    </div>
                </div>
            );
        } else {
            return <div className="error-page">
                <div className="error-text">
                    <div>{error?.message}</div>
                </div>
            </div>;
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